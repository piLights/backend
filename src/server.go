package main

import (
	"errors"
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

//calculateOpacity calculates the value of colorValue after applying some opacity
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

//SetColor sets the color of the Dioder-strips
func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	if *debug {
		log.Println("SetColor:", colorMessage)
	}

	if *password != "" && *password != colorMessage.Password {
		if *debug {
			log.Println("Not authorized")
		}
		return nil, errors.New("Not authorized")
	}

	opacity := uint8(colorMessage.Opacity)
	red := calculateOpacity(uint8(colorMessage.R), opacity)
	green := calculateOpacity(uint8(colorMessage.G), opacity)
	blue := calculateOpacity(uint8(colorMessage.B), opacity)

	colorSet := color.RGBA{red, green, blue, opacity}

	dioder.SetAll(colorSet)

	return &LighterGRPC.Confirmation{true}, nil
}

//CheckConnection checks the connection and returns the current settings
func (s *server) CheckConnection(ctx context.Context, initMessage *LighterGRPC.InitMessage) (*LighterGRPC.ColorMessage, error) {
	if *debug {
		log.Println("CheckConnection", initMessage)
	}

	if *password != "" && *password != initMessage.Password {
		if *debug {
			log.Println("Not authorized")
		}
		return nil, errors.New("Not authorized")
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

//SwitchState switches the state (on/off) of the Didoer-Strips
func (s *server) SwitchState(ctx context.Context, stateMessage *LighterGRPC.StateMessage) (*LighterGRPC.Confirmation, error) {
	if *debug {
		log.Println("SwitchState", stateMessage)
	}

	if *password != "" && *password != stateMessage.Password {
		if *debug {
			log.Println("Not authorized")
		}
		return nil, errors.New("Not authorized")
	}

	if stateMessage.Onstate {
		dioder.TurnOn()
	} else {
		dioder.TurnOff()
	}

	return &LighterGRPC.Confirmation{true}, nil
}

//startServer starts the GRPC-server and binds to the defined address
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
