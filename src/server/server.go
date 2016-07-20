package server

import (
	"errors"
	"fmt"
	"image/color"
	"net"

	"google.golang.org/grpc"

	"github.com/piLights/dioder"
	"github.com/piLights/dioder-rpc/src/configuration"
	"github.com/piLights/dioder-rpc/src/logging"
	"github.com/piLights/dioder-rpc/src/proto"
	"golang.org/x/net/context"
)

//server implements the server-interface required by GRPC
type server struct{}

//Holds all streams
var (
	dioderInstance dioder.Dioder
	streams        map[string]LighterGRPC.Lighter_CheckConnectionServer
	onState        bool
)

//SetColor sets the color of the Dioder-strips
func (s *server) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	if configuration.DioderConfiguration.Debug {
		logging.Log.LogChan <- fmt.Sprint("SetColor:", colorMessage)
	}

	if configuration.DioderConfiguration.Password != "" && configuration.DioderConfiguration.Password != colorMessage.Password {
		if configuration.DioderConfiguration.Debug {
			logging.Log.LogChan <- "Not authorized"
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
			if configuration.DioderConfiguration.Debug {
				logging.Log.LogChan <- fmt.Sprintf("Sending the colormessage to remote device %s\n", deviceID)
			}

			stream.Send(colorMessage)
		}
	}

	return &LighterGRPC.Confirmation{true}, nil
}

func (s *server) CheckConnection(initMessage *LighterGRPC.InitMessage, stream LighterGRPC.Lighter_CheckConnectionServer) error {
	if configuration.DioderConfiguration.Debug {
		logging.Log.LogChan <- fmt.Sprint("CheckConnection", initMessage)
	}

	if configuration.DioderConfiguration.Password != "" && configuration.DioderConfiguration.Password != initMessage.Password {
		error := errors.New("Not authorized")
		if configuration.DioderConfiguration.Debug {
			logging.Log.LogChan <- error
		}
		return error
	}

	colorSet := dioderInstance.GetCurrentColor()

	if configuration.DioderConfiguration.Debug {
		logging.Log.LogChan <- fmt.Sprint("CheckConnection: Returning the current settings:", colorSet)
		logging.Log.LogChan <- "Saving the stream-connection"
	}

	streams[initMessage.DeviceID] = stream

	error := stream.Send(&LighterGRPC.ColorMessage{onState, int32(colorSet.R), int32(colorSet.G), int32(colorSet.B), int32(colorSet.A), "Dioder-Server", ""})
	if error != nil && configuration.DioderConfiguration.Debug {
		logging.Log.LogChan <- error
	}

	return error
}

func (s *server) ChangeServerParameter(ctx context.Context, changeParameterMessage *LighterGRPC.ChangeParameterMessage) (*LighterGRPC.Confirmation, error) {
	return nil, errors.New("Not implemented")
}

func (s *server) LoadServerConfig(ctx context.Context, changeParameterMessage *LighterGRPC.LoadConfigRequest) (*LighterGRPC.ServerConfig, error) {
	return nil, errors.New("Not implemented")
}

func (s *server) ScheduleSwitchState(ctx context.Context, changeParameterMessage *LighterGRPC.ScheduledSwitch) (*LighterGRPC.Confirmation, error) {
	return nil, errors.New("Not implemented")
}

//SwitchState switches the state (on/off) of the Didoer-Strips
func (s *server) SwitchState(ctx context.Context, stateMessage *LighterGRPC.StateMessage) (*LighterGRPC.Confirmation, error) {
	if configuration.DioderConfiguration.Debug {
		logging.Log.LogChan <- fmt.Sprintln("SwitchState", stateMessage)
	}

	if configuration.DioderConfiguration.Password != "" && configuration.DioderConfiguration.Password != stateMessage.Password {
		if configuration.DioderConfiguration.Debug {
			logging.Log.LogChan <- "Not authorized"
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

//StartServer starts the GRPC-server and binds to the defined address
func StartServer(dioderInst dioder.Dioder) {
	if configuration.DioderConfiguration.Debug {
		logging.Log.LogChan <- fmt.Sprintf("Binding to %s", configuration.DioderConfiguration.BindTo)
	}

	dioderInstance = dioderInst

	//Initialize the streams-map
	streams = make(map[string]LighterGRPC.Lighter_CheckConnectionServer)

	protocol := "tcp"

	if configuration.DioderConfiguration.IPv4Only {
		if configuration.DioderConfiguration.Debug {
			logging.Log.LogChan <- "Forced to listen on IPv4 only."
		}

		protocol = "tcp4"
	}

	if configuration.DioderConfiguration.IPv6Only {
		if configuration.DioderConfiguration.Debug {
			logging.Log.LogChan <- "Forced to listen on IPv6 only."
		}

		protocol = "tcp6"
	}

	listener, error := net.Listen(protocol, configuration.DioderConfiguration.BindTo)
	if error != nil {
		logging.Log.LogChan <- fmt.Sprintf("failed to listen: %v", error)
	}

	grpcServer := grpc.NewServer()

	LighterGRPC.RegisterLighterServer(grpcServer, &server{})

	fmt.Printf("Listening on %s...\n", configuration.DioderConfiguration.BindTo)

	grpcServer.Serve(listener)
}
