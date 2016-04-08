package main

import (
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

type server struct{}

func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	dioder.SetAll(uint8(colorMessage.R), uint8(colorMessage.G), uint8(colorMessage.B))

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
