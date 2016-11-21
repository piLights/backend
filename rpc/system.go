package rpc

import (
	"golang.org/x/net/context"

	"gitlab.com/piLights/dioder-rpc/piLightsVersion"
	"gitlab.com/piLights/proto"
)

type systemServer struct{}

func (s *systemServer) ChangeServerParameter(ctx context.Context, changeParameterMessage *LighterGRPC.ChangeParameterMessage) (*LighterGRPC.Confirmation, error) {
	return nil, errNotImplemented
}

func (s *systemServer) LoadServerConfiguration(ctx context.Context, request *LighterGRPC.Request) (*LighterGRPC.ServerConfiguration, error) {
	return nil, errNotImplemented
}

func (s *systemServer) SetServerConfiguration(ctx context.Context, serverConfiguration *LighterGRPC.ServerConfiguration) (*LighterGRPC.Confirmation, error) {
	return nil, errNotImplemented
}

func (s *systemServer) Version(ctx context.Context, request *LighterGRPC.Empty) (*LighterGRPC.BackendVersion, error) {
	if !checkAccess(request) {
		return nil, errNotAuthorized
	}

	return &LighterGRPC.BackendVersion{
		VersionCode:     piLightsVersion.Version,
		UpdateAvailable: false, // @ToDo!
	}, nil
}
