package server

import (
	"context"
	"github.com/ixbaseANT/gord/cmd/gorwallet/daemon/pb"
	"github.com/kaspanet/kaspad/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
