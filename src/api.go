package main

import (
	"context"
	"fmt"
	"net/http"

	"gitlab.com/piLights/dioder-rpc/src/configuration"
	"gitlab.com/piLights/dioder-rpc/src/logging"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "gitlab.com/piLights/proto"
	"google.golang.org/grpc"
)

func startAPI() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	grpcOptions := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial(configuration.DioderConfiguration.BindTo, grpcOptions...)
	if err != nil {
		logging.FatalChan <- err
	}

	defer func() {
		if err != nil {
			cerr := conn.Close()
			if cerr != nil {
				logging.FatalChan <- fmt.Sprintf("Failed to close conn to %s: %v", configuration.DioderConfiguration.BindTo, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			cerr := conn.Close()
			if cerr != nil {
				logging.FatalChan <- fmt.Sprintf("Failed to close conn to %s: %v", configuration.DioderConfiguration.BindTo, cerr)
			}
		}()
	}()

	err = gw.RegisterSystemHandler(ctx, mux, conn)
	if err != nil {
		logging.FatalChan <- err
	}

	err = gw.RegisterLighterHandler(ctx, mux, conn)
	if err != nil {
		logging.FatalChan <- err
	}

	http.ListenAndServe(":8080", mux)
}
