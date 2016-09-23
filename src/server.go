package main

import (
	"errors"
	"fmt"
	"image/color"
	"net"

	"google.golang.org/grpc"

	"gitlab.com/piLights/proto"
	"golang.org/x/net/context"
)

//server implements the server-interface required by GRPC
type server struct{}

//Holds all streams
var streams map[string]LighterGRPC.Lighter_CheckConnectionServer

var onState bool

//SetColor sets the color of the Dioder-strips
func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	if DioderConfiguration.Debug {
		logChan <- fmt.Sprint("SetColor:", colorMessage)
	}

	if DioderConfiguration.Password != "" && DioderConfiguration.Password != colorMessage.Password {
		if DioderConfiguration.Debug {
			logChan <- "Not authorized"
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
			if DioderConfiguration.Debug {
				logChan <- fmt.Sprintf("Sending the colormessage to remote device %s\n", deviceID)
			}

			stream.Send(colorMessage)
		}
	}

	return &LighterGRPC.Confirmation{true}, nil
}

func (s *server) CheckConnection(initMessage *LighterGRPC.InitMessage, stream LighterGRPC.Lighter_CheckConnectionServer) error {
	if DioderConfiguration.Debug {
		logChan <- fmt.Sprint("CheckConnection", initMessage)
	}

	if DioderConfiguration.Password != "" && DioderConfiguration.Password != initMessage.Password {
		error := errors.New("Not authorized")
		if DioderConfiguration.Debug {
			logChan <- error
		}
		return error
	}

	colorSet := dioderInstance.GetCurrentColor()

	if DioderConfiguration.Debug {
		logChan <- fmt.Sprint("CheckConnection: Returning the current settings:", colorSet)
		logChan <- "Saving the stream-connection"
	}

	streams[initMessage.DeviceID] = stream

	error := stream.Send(&LighterGRPC.ColorMessage{onState, int32(colorSet.R), int32(colorSet.G), int32(colorSet.B), int32(colorSet.A), "Dioder-Server", ""})
	if error != nil && DioderConfiguration.Debug {
		logChan <- error
	}

	return error
}

func (s *server) ChangeServerParameter(ctx context.Context, changeParameterMessage *LighterGRPC.ChangeParameterMessage) (*LighterGRPC.Confirmation, error) {
	return nil, errors.New("Not implemented")
}

func (s *server) LoadServerConfig(ctx context.Context, changeParameterMessage *LighterGRPC.LoadConfigRequest) (*LighterGRPC.ServerConfig, error) {
	return nil, errors.New("Not implemented")
}

func (s *server) LoadServerLog(logRequest *LighterGRPC.LogRequest, server LighterGRPC.Lighter_LoadServerLogServer) error {
	return errors.New("Not implemented")
}

func (s *server) ScheduleSwitchState(ctx context.Context, changeParameterMessage *LighterGRPC.ScheduledSwitch) (*LighterGRPC.Confirmation, error) {
	return nil, errors.New("Not implemented")
}

//SwitchState switches the state (on/off) of the Didoer-Strips
func (s *server) SwitchState(ctx context.Context, stateMessage *LighterGRPC.StateMessage) (*LighterGRPC.Confirmation, error) {
	if DioderConfiguration.Debug {
		logChan <- fmt.Sprintln("SwitchState", stateMessage)
	}

	if DioderConfiguration.Password != "" && DioderConfiguration.Password != stateMessage.Password {
		if DioderConfiguration.Debug {
			logChan <- "Not authorized"
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
	if DioderConfiguration.Debug {
		logChan <- fmt.Sprintf("Binding to %s", DioderConfiguration.BindTo)
	}

	//Initialize the streams-map
	streams = make(map[string]LighterGRPC.Lighter_CheckConnectionServer)

	protocol := "tcp"

	if DioderConfiguration.IPv4Only {
		if DioderConfiguration.Debug {
			logChan <- "Forced to listen on IPv4 only."
		}

		protocol = "tcp4"
	}

	if DioderConfiguration.IPv6Only {
		if DioderConfiguration.Debug {
			logChan <- "Forced to listen on IPv6 only."
		}

		protocol = "tcp6"
	}

	listener, error := net.Listen(protocol, DioderConfiguration.BindTo)
	if error != nil {
		logChan <- fmt.Sprintf("failed to listen: %v", error)
	}

	grpcServer := grpc.NewServer()

	LighterGRPC.RegisterLighterServer(grpcServer, &server{})

	fmt.Printf("Listening on %s...\n", DioderConfiguration.BindTo)

	grpcServer.Serve(listener)
}
