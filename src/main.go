package main

import (
	"image/color"
	"log"
	"net"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/piLights/dioder"
	"golang.org/x/net/context"

	LighterGRPC "github.com/piLights/dioder-rpc/src/proto"
	"google.golang.org/grpc"
)

//server implements the server-interface required by GRPC
type server struct{}

func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	if *debug {
		log.Println("SetColor:", colorMessage)
	}

	colorSet := color.RGBA{uint8(colorMessage.R), uint8(colorMessage.G), uint8(colorMessage.B), uint8(colorMessage.Opacity)}

	dioder.SetAll(colorSet)

	return &LighterGRPC.Confirmation{true}, nil
}

func (s *server) CheckConnection(ctx context.Context, initMessage *LighterGRPC.InitMessage) (*LighterGRPC.ColorMessage, error) {
	if *debug {
		log.Println("CheckConnection", initMessage)
	}

	colorSet := dioder.GetCurrentColor()

	onstate := true

	if colorSet.R == 0 && colorSet.G == 0 && colorSet.B == 0 {
		onstate = false
	}

	if *debug {
		log.Println("CheckConnection: Returning the current settings:", colorSet)
	}
	return &LighterGRPC.ColorMessage{onstate, int32(colorSet.R), int32(colorSet.G), int32(colorSet.B), int32(colorSet.A), "Dioder-Server"}, nil
}

func (s *server) SwitchState(ctx context.Context, stateMessage *LighterGRPC.StateMessage) (*LighterGRPC.Confirmation, error) {
	if *debug {
		log.Println("SwitchState", stateMessage)
	}

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
	if *debug {
		log.Printf("Configuring the Pins to: Red: %d, Green: %d, Blue: %d\n", *redPin, *greenPin, *bluePin)
	}
	dioder.SetPins(*redPin, *greenPin, *bluePin)

	if *debug {
		log.Printf("Binding to %s", *bindTo)
	}
	listener, error := net.Listen("tcp", *bindTo)
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}

	grpcServer := grpc.NewServer()

	LighterGRPC.RegisterLighterServer(grpcServer, &server{})

	grpcServer.Serve(listener)
}
