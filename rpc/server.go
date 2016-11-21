package rpc

import (
	"fmt"
	"net"

	"gitlab.com/piLights/dioder-rpc/configuration"
	"gitlab.com/piLights/dioder-rpc/logging"
	"gitlab.com/piLights/proto"

	"google.golang.org/grpc"
)

// StartServer starts the GRPC-server and binds to the defined address
func StartServer() {
	if configuration.DioderConfiguration.Debug {
		logging.LogChan <- fmt.Sprintf("Binding to %s", configuration.DioderConfiguration.BindTo)
	}

	//Initialize the streams-map
	streams = make(map[string]LighterGRPC.Lighter_OpenStreamServer)
	colorStream = make(chan *LighterGRPC.ColorMessage)

	protocol := "tcp"

	if configuration.DioderConfiguration.IPv4Only {
		if configuration.DioderConfiguration.Debug {
			logging.LogChan <- "Forced to listen on IPv4 only."
		}

		protocol = "tcp4"
	}

	if configuration.DioderConfiguration.IPv6Only {
		if configuration.DioderConfiguration.Debug {
			logging.LogChan <- "Forced to listen on IPv6 only."
		}

		protocol = "tcp6"
	}

	listener, error := net.Listen(protocol, configuration.DioderConfiguration.BindTo)
	if error != nil {
		logging.LogChan <- fmt.Sprintf("failed to listen: %v", error)
	}

	grpcServer := grpc.NewServer()

	LighterGRPC.RegisterLighterServer(grpcServer, &lighterServer{})
	LighterGRPC.RegisterSystemServer(grpcServer, &systemServer{})

	fmt.Printf("Listening on %s...\n", configuration.DioderConfiguration.BindTo)

	grpcServer.Serve(listener)
}
