package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/piLights/dioder-rpc/src/proto"
	"golang.org/x/net/context"
)

//server implements the server-interface required by GRPC
type server struct{}

//Holds all streams
var streams map[string]LighterGRPC.Lighter_CheckConnectionServer

var onState bool

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
	red := uint8(colorMessage.R)
	green := uint8(colorMessage.G)
	blue := uint8(colorMessage.B)

	colorSet := color.RGBA{red, green, blue, opacity}

	dioderInstance.SetAll(colorSet)

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

	colorSet := dioderInstance.GetCurrentColor()

	if *debug {
		log.Println("CheckConnection: Returning the current settings:", colorSet)
		log.Println("Saving the stream-connection")
	}

	streams[initMessage.DeviceID] = stream

	error := stream.Send(&LighterGRPC.ColorMessage{onState, int32(colorSet.R), int32(colorSet.G), int32(colorSet.B), int32(colorSet.A), "Dioder-Server", ""})
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
		dioderInstance.TurnOn()
	} else {
		dioderInstance.TurnOff()
	}
	
	onState = stateMessage.Onstate
	
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

	fmt.Printf("Listening on %s...\n", *bindTo)

	grpcServer.Serve(listener)
}
