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

//Holds all streams
var streams map[string]LighterGRPC.Lighter_CheckConnectionServer

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

	for deviceID, stream := range streams {
		if deviceID != colorMessage.DeviceID {
			if *debug {
				log.Printf("Sending the colormessage to remote device %s\n", deviceID)
			}

			stream.Send(colorMessage)
		}
	}

	return &LighterGRPC.Confirmation{true}, nil
}

func (s *server) CheckConnection(initMessage *LighterGRPC.InitMessage, stream LighterGRPC.Lighter_CheckConnectionServer) error {
	if *debug {
		log.Println("CheckConnection", initMessage)
	}

	if *password != "" && *password != initMessage.Password {
		error := errors.New("Not authorized")
		if *debug {
			log.Println(error)
		}
		return error
	}

	colorSet := dioder.GetCurrentColor()

	onstate := true

	if colorSet.R == 0 && colorSet.G == 0 && colorSet.B == 0 {
		onstate = false
	}

	if *debug {
		log.Println("CheckConnection: Returning the current settings:", colorSet)
		log.Println("Saving the stream-connection")
	}

	streams[initMessage.DeviceID] = stream

	error := stream.Send(&LighterGRPC.ColorMessage{onstate, int32(colorSet.R), int32(colorSet.G), int32(colorSet.B), int32(colorSet.A), "Dioder-Server", ""})
	if error != nil && *debug {
		log.Println(error)
	}

	return error
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

	//Initialize the streams-map
	streams = make(map[string]LighterGRPC.Lighter_CheckConnectionServer)

	listener, error := net.Listen("tcp", *bindTo)
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}

	grpcServer := grpc.NewServer()

	LighterGRPC.RegisterLighterServer(grpcServer, &server{})

	grpcServer.Serve(listener)
}
