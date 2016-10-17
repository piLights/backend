package rpc

import (
	"errors"
	"fmt"
	"image/color"

	"gitlab.com/piLights/dioder-rpc/src/configuration"
	"gitlab.com/piLights/dioder-rpc/src/logging"
	"gitlab.com/piLights/proto"
	"golang.org/x/net/context"
)

//server implements the server-interface required by GRPC
type lighterServer struct{}

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
func (s *lighterServer) SetColor(ctx context.Context, colorMessage *LighterGRPC.ColorMessage) (*LighterGRPC.Confirmation, error) {
	if configuration.DioderConfiguration.Debug {
		logging.LogChan <- fmt.Sprint("SetColor:", colorMessage)
	}

	if !checkAccess(colorMessage) {
		return nil, errNotAuthorized
	}

	opacity := uint8(colorMessage.Opacity)
	red := uint8(colorMessage.R)
	green := uint8(colorMessage.G)
	blue := uint8(colorMessage.B)

	colorSet := color.RGBA{red, green, blue, opacity}

	configuration.DioderConfiguration.DioderInstance.SetAll(colorSet)

	if len(streams) > 0 {
		if configuration.DioderConfiguration.Debug {
			logging.LogChan <- fmt.Sprintf("Sending the colormessage to all remote devices")
		}
		colorStream <- colorMessage
	}

	return &LighterGRPC.Confirmation{Success: true}, nil
}

func (s *lighterServer) CheckConnection(request *LighterGRPC.Request, stream LighterGRPC.Lighter_CheckConnectionServer) error {
	if configuration.DioderConfiguration.Debug {
		logging.LogChan <- fmt.Sprint("CheckConnection", request)
	}

	if !checkAccess(request) {
		return errNotAuthorized
	}

	colorSet := configuration.DioderConfiguration.DioderInstance.GetCurrentColor()

	if configuration.DioderConfiguration.Debug {
		logging.LogChan <- fmt.Sprint("CheckConnection: Returning the current settings:", colorSet)
		logging.LogChan <- "Saving the stream-connection"
	}

	streams[request.DeviceID] = stream

	//@ToDo: Do not use unkeyed fields!
	error := stream.Send(&LighterGRPC.ColorMessage{onState, int32(colorSet.R), int32(colorSet.G), int32(colorSet.B), int32(colorSet.A), "Dioder-Server", ""})
	if error != nil && configuration.DioderConfiguration.Debug {
		logging.LogChan <- error
	}

	for colorMessage := range colorStream {
		error = stream.Send(colorMessage)
		if error != nil && configuration.DioderConfiguration.Debug {
			logging.LogChan <- error
		}
	}

	return error
}

func (s *lighterServer) LoadServerLog(logRequest *LighterGRPC.LogRequest, server LighterGRPC.Lighter_LoadServerLogServer) error {
	for _, logEntry := range logging.GetLogEntryList(logRequest.Amount) {
		error := server.Send(logEntry)
		if error != nil {
			return error
		}
	}

	return nil
}

func (s *lighterServer) ScheduleSwitchState(ctx context.Context, changeParameterMessage *LighterGRPC.ScheduledSwitch) (*LighterGRPC.Confirmation, error) {
	return nil, errNotImplemented
}

//SwitchState switches the state (on/off) of the Didoer-Strips
func (s *lighterServer) SwitchState(ctx context.Context, stateMessage *LighterGRPC.StateMessage) (*LighterGRPC.Confirmation, error) {
	if configuration.DioderConfiguration.Debug {
		logging.LogChan <- fmt.Sprintln("SwitchState", stateMessage)
	}

	if configuration.DioderConfiguration.Password != "" && configuration.DioderConfiguration.Password != stateMessage.Password {
		if configuration.DioderConfiguration.Debug {
			logging.LogChan <- "Not authorized"
		}
		return nil, errNotAuthorized
	}

	if stateMessage.Onstate {
		configuration.DioderConfiguration.DioderInstance.TurnOn()
	} else {
		configuration.DioderConfiguration.DioderInstance.TurnOff()
	}

	onState = stateMessage.Onstate

	return &LighterGRPC.Confirmation{Success: true}, nil
}

func checkAccess(request interface{}) bool {
	if configuration.DioderConfiguration.Password != "" && configuration.DioderConfiguration.Password != request.(struct{ Password string }).Password {
		if configuration.DioderConfiguration.Debug {
			logging.LogChan <- "Not authorized"
		}
		return false
	}

	return true
}
