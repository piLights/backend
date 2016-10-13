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

var (
	//Holds all streams
	streams map[string]LighterGRPC.Lighter_CheckConnectionServer

	onState bool

	colorStream chan *LighterGRPC.ColorMessage

	//Errors
	errNotAuthorized  = errors.New("Not authorized")
	errNotImplemented = errors.New("Not implemented")
)

//SetColor sets the color of the Dioder-strips
func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	if DioderConfiguration.Debug {
		logChan <- fmt.Sprint("SetColor:", colorMessage)
	}

	if DioderConfiguration.Password != "" && DioderConfiguration.Password != colorMessage.Password {
		if DioderConfiguration.Debug {
			logChan <- "Not authorized"
		}
		return nil, errNotAuthorized
	}

	opacity := uint8(colorMessage.Opacity)
	red := uint8(colorMessage.R)
	green := uint8(colorMessage.G)
	blue := uint8(colorMessage.B)

	colorSet := color.RGBA{red, green, blue, opacity}

	dioderInstance.SetAll(colorSet)

	if len(streams) > 0 {
		if DioderConfiguration.Debug {
			logChan <- fmt.Sprintf("Sending the colormessage to all remote devices")
		}
		colorStream <- colorMessage
	}

	return &LighterGRPC.Confirmation{Success: true}, nil
}

func (s *server) CheckConnection(request *LighterGRPC.Request, stream LighterGRPC.Lighter_CheckConnectionServer) error {
	if DioderConfiguration.Debug {
		logChan <- fmt.Sprint("CheckConnection", request)
	}

	if DioderConfiguration.Password != "" && DioderConfiguration.Password != request.Password {
		error := errNotAuthorized
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

	streams[request.DeviceID] = stream

	//@ToDo: Do not use unkeyed fields!
	error := stream.Send(&LighterGRPC.ColorMessage{onState, int32(colorSet.R), int32(colorSet.G), int32(colorSet.B), int32(colorSet.A), "Dioder-Server", ""})
	if error != nil && DioderConfiguration.Debug {
		logChan <- error
	}

	for colorMessage := range colorStream {
		error = stream.Send(colorMessage)
		if error != nil && DioderConfiguration.Debug {
			logChan <- error
		}
	}

	return error
}

func (s *server) ChangeServerParameter(ctx context.Context, changeParameterMessage *LighterGRPC.ChangeParameterMessage) (*LighterGRPC.Confirmation, error) {
	return nil, errNotImplemented
}

func (s *server) LoadServerConfiguration(ctx context.Context, changeParameterMessage *LighterGRPC.Request) (*LighterGRPC.ServerConfiguration, error) {
	return nil, errNotImplemented
}

func (s *server) SetServerConfiguration(ctx context.Context, serverConfiguration *LighterGRPC.ServerConfiguration) (*LighterGRPC.Confirmation, error) {
	return nil, errNotImplemented
}

func (s *server) LoadServerLog(logRequest *LighterGRPC.LogRequest, server LighterGRPC.Lighter_LoadServerLogServer) error {
	for _, logEntry := range getLogEntryList().EntryList {
		error := server.Send(logEntry)
		if error != nil {
			return error
		}
	}

	return nil
}

func (s *server) ScheduleSwitchState(ctx context.Context, changeParameterMessage *LighterGRPC.ScheduledSwitch) (*LighterGRPC.Confirmation, error) {
	return nil, errNotImplemented
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
		return nil, errNotAuthorized
	}

	if stateMessage.Onstate {
		dioderInstance.TurnOn()
	} else {
		dioderInstance.TurnOff()
	}

	onState = stateMessage.Onstate

	return &LighterGRPC.Confirmation{Success: true}, nil
}

//startServer starts the GRPC-server and binds to the defined address
func startServer() {
	if DioderConfiguration.Debug {
		logChan <- fmt.Sprintf("Binding to %s", DioderConfiguration.BindTo)
	}

	//Initialize the streams-map
	streams = make(map[string]LighterGRPC.Lighter_CheckConnectionServer)
	colorStream = make(chan *LighterGRPC.ColorMessage)

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
