package rpc

import (
	"errors"
	"fmt"
	"image/color"

	"gitlab.com/piLights/dioder-rpc/configuration"
	"gitlab.com/piLights/dioder-rpc/logging"
	"gitlab.com/piLights/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

//server implements the server-interface required by GRPC
type lighterServer struct{}

var (
	//Holds all streams
	streams map[string]LighterGRPC.RgbService_OpenStreamServer

	onState bool

	colorStream chan *LighterGRPC.ColorMessage

	//Errors
	errNotAuthorized = errors.New("Not authorized")
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

	colorSet := color.RGBA{
		R: red,
		G: green,
		B: blue,
		A: opacity,
	}

	configuration.DioderConfiguration.DioderInstance.SetAll(colorSet)

	if len(streams) > 0 {
		if configuration.DioderConfiguration.Debug {
			logging.LogChan <- fmt.Sprintf("Sending the colormessage to all remote devices")
		}
		colorStream <- colorMessage
	}

	return &LighterGRPC.Confirmation{Success: true}, nil
}

func (s *lighterServer) OpenStream(request *LighterGRPC.Request, stream LighterGRPC.RgbService_OpenStreamServer) error {
	if configuration.DioderConfiguration.Debug {
		logging.LogChan <- fmt.Sprint("OpenStream", request)
	}

	if !checkAccess(request) {
		return errNotAuthorized
	}

	colorSet := configuration.DioderConfiguration.DioderInstance.GetCurrentColor()

	if configuration.DioderConfiguration.Debug {
		logging.LogChan <- fmt.Sprint("OpenStream: Returning the current settings:", colorSet)
		logging.LogChan <- "Saving the stream-connection"
	}

	md, ok := metadata.FromContext(stream.Context())
	if !ok {
		logging.LogChan <- "OpenStream: Unable to get request metadata"
	} else {
		deviceIDSlice := md["DeviceID"]

		logging.LogChan <- fmt.Sprintf("DeviceID slice: %v", deviceIDSlice)
		if len(deviceIDSlice) == 0 {
			logging.LogChan <- "OpenStream: Client did not send his deviceID. Unable to save the stream"
		} else {
			deviceID := deviceIDSlice[0]
			streams[deviceID] = stream
		}
	}

	error := stream.Send(&LighterGRPC.ColorMessage{
		Onstate: onState,
		R:       int32(colorSet.R),
		G:       int32(colorSet.G),
		B:       int32(colorSet.B),
		Opacity: int32(colorSet.A),
	})

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

func (s *lighterServer) GetStatus(ctx context.Context, request *LighterGRPC.Request) (*LighterGRPC.ColorMessage, error) {
	if configuration.DioderConfiguration.Debug {
		logging.LogChan <- fmt.Sprint("GetStatus", request)
	}

	if !checkAccess(request) {
		return nil, errNotAuthorized
	}

	colorSet := configuration.DioderConfiguration.DioderInstance.GetCurrentColor()
	return &LighterGRPC.ColorMessage{
		Onstate: onState,
		R:       int32(colorSet.R),
		G:       int32(colorSet.G),
		B:       int32(colorSet.B),
		Opacity: int32(colorSet.A),
	}, nil
}

func (s *lighterServer) LoadServerLog(logRequest *LighterGRPC.LogRequest, server LighterGRPC.RgbService_LoadServerLogServer) error {
	if !checkAccess(server) {
		return errNotAuthorized
	}

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

	if !checkAccess(stateMessage) {
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
