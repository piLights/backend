package main

import (
	"image/color"
	"log"
	"net"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/piLights/dioder"
	"golang.org/x/net/context"

	LighterGRPC "./proto"
	"google.golang.org/grpc"
)

//server implements the server-interface required by GRPC
type server struct{}

func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	colorSet := color.RGBA{uint8(colorMessage.R), uint8(colorMessage.G), uint8(colorMessage.B), uint8(colorMessage.Opacity)}

	dioder.SetAll(colorSet)

	return &LighterGRPC.Confirmation{true}, nil
}

func (s *server) CheckConnection(ctx context.Context, initMessage *LighterGRPC.InitMessage) (*LighterGRPC.ColorMessage, error) {
	colorSet := dioder.GetCurrentColor()

	onstate := true

	if colorSet.R == 0 && colorSet.G == 0 && colorSet.B == 0 {
		onstate = false
	}

	return &LighterGRPC.ColorMessage{onstate, int32(colorSet.R), int32(colorSet.G), int32(colorSet.B), int32(colorSet.A), "Dioder-Server"}, nil
}

func (s *server) SwitchState(ctx context.Context, stateMessage *LighterGRPC.StateMessage) (*LighterGRPC.Confirmation, error) {
	if stateMessage.Onstate {
		dioder.SetAll(color.RGBA{255, 255, 255, 255})
	} else {
		dioder.SetAll(color.RGBA{})
	}

	return &LighterGRPC.Confirmation{true}, nil
}

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

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
