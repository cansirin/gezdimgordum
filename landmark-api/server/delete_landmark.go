package server

import (
	"context"
	api "github.com/cansirin/gezdimgordum/landmark-api/rpc/landmark-api"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *LandmarkAPIServer) DeleteLandmark(ctx context.Context, req *api.DeleteLandmarkRequest) (*emptypb.Empty, error) {
	if err := validateDeleteLandmarkRequest(req); err != nil {
		return nil, err
	}

	err := s.backend.DeleteLandmark(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteLandmarkRequest(req *api.DeleteLandmarkRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("ID")
	}
	return nil
}
