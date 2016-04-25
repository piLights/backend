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
	bindTo   = flag.String("bindTo", ":13337", "Address and port to listen on, defaults to 0.0.0.0:13337")
	redPin   = flag.Int("red", "18", "Number of the red-pin")
	greenPin = flag.Int("green", "4", "Number of the green-pin")
	bluePin  = flag.Int("red", "17", "Number of the blue-pin")
)

//server implements the server-interface required by GRPC
type server struct{}

func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	dioder.SetAll(uint8(colorMessage.R), uint8(colorMessage.G), uint8(colorMessage.B))

	return &LighterGRPC.Confirmation{true}, nil
}

func (s *server) CheckConnection(ctx context.Context, initMessage *LighterGRPC.InitMessage) (*LighterGRPC.ColorMessage, error) {
	colorMap := dioder.GetCurrentColor()

	onstate := true

	if colorMap[0] == 0 && colorMap[1] == 0 && colorMap[2] == 0 {
		onstate = false
	}

	return &LighterGRPC.ColorMessage{onstate, int32(colorMap[0]), int32(colorMap[1]), int32(colorMap[2]), 0, "Dioder-Server"}, nil
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

	//Set the pins
	dioder.SetPins(*redPin, *greenPin, *bluePin)

	listener, error := net.Listen("tcp", *bindTo)
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}

	grpcServer := grpc.NewServer()

	LighterGRPC.RegisterLighterServer(grpcServer, &server{})

	grpcServer.Serve(listener)
}
