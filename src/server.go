package main

import (
	"image/color"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/piLights/dioder"
	"github.com/piLights/dioder-rpc/src/proto"
	"golang.org/x/net/context"
)

//server implements the server-interface required by GRPC
type server struct{}

func calculateOpacity(colorValue uint8, opacity uint8) uint8 {
	var calculatedValue float32

	if opacity != 100 {
		calculatedValue = float32(colorValue) / 100 * float32(opacity)
	} else {
		calculatedValue = float32(colorValue)
	}

	if *debug {
		log.Printf("Applying opacity (%d) to color with value %d - Result: %d", opacity, colorValue, uint8(calculatedValue))
	}

	return uint8(calculatedValue)
}

func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	if *debug {
		log.Println("SetColor:", colorMessage)
	}

	opacity := uint8(colorMessage.Opacity)
	red := calculateOpacity(uint8(colorMessage.R), opacity)
	green := calculateOpacity(uint8(colorMessage.G), opacity)
	blue := calculateOpacity(uint8(colorMessage.B), opacity)

	colorSet := color.RGBA{red, green, blue, opacity}

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
	return &LighterGRPC.ColorMessage{onstate, int32(colorSet.R), int32(colorSet.G), int32(colorSet.B), int32(colorSet.A), "Dioder-Server", ""}, nil
}

func (s *server) SwitchState(ctx context.Context, stateMessage *LighterGRPC.StateMessage) (*LighterGRPC.Confirmation, error) {
	if *debug {
		log.Println("SwitchState", stateMessage)
	}

	if stateMessage.Onstate {
		dioder.SetAll(color.RGBA{255, 255, 255, 100})
	} else {
		dioder.SetAll(color.RGBA{})
	}

	return &LighterGRPC.Confirmation{true}, nil
}

func startServer() {
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
