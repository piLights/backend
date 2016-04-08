package main

import (
	"errors"
	"flag"
	"log"
	"net"

	"github.com/piLights/dioder"
	"golang.org/x/net/context"

	LighterGRPC "./proto"
	"google.golang.org/grpc"
)

var (
	bindTo = flag.String("bindTo", ":13337", "Address and port to listen on, defaults to 0.0.0.0:13337")
)

/*
type LighterServer interface {
	SetColor(context.Context, *ColorMessage) (*Confirmation, error)
	CheckConnection(context.Context, *InitMessage) (*ColorMessage, error)
	SwitchState(context.Context, *StateMessage) (*Confirmation, error)
}
*/
type server struct{}

func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	dioder.SetAll(uint8(colorMessage.R), uint8(colorMessage.G), uint8(colorMessage.B))

	return &LighterGRPC.Confirmation{true}, nil
}

func (s *server) CheckConnection(ctx context.Context, initMessage *LighterGRPC.InitMessage) (*LighterGRPC.ColorMessage, error) {
	return nil, errors.New("CheckConnection is currently not implemented")
}

func (s *server) SwitchState(ctx context.Context, stateMessage *LighterGRPC.StateMessage) (*LighterGRPC.Confirmation, error) {
	if stateMessage.Onstate {
		dioder.SetAll(255, 255, 255)
	} else {
		dioder.SetAll(0, 0, 0)
	}

	return &LighterGRPC.Confirmation{true}, nil
}

func main() {
	flag.Parse()

	listener, error := net.Listen("tcp", *bindTo)
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}

	grpcServer := grpc.NewServer()

	LighterGRPC.RegisterLighterServer(grpcServer, &server{})

	grpcServer.Serve(listener)
}
