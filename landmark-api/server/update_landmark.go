package server

import (
	"context"
	api "github.com/cansirin/gezdimgordum/landmark-api/rpc/landmark-api"
	"github.com/cansirin/gezdimgordum/landmark-api/server/helper"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *LandmarkAPIServer) UpdateLandmark(ctx context.Context, req *api.UpdateLandmarkRequest) (*emptypb.Empty, error) {
	if err := validateUpdateLandmarkRequest(req); err != nil {
		return nil, err
	}
	err := s.backend.UpdateLandmark(ctx, req.ID, helper.ConvertToStringPtr(req.Name), helper.ConvertToStringPtr(req.Description),
		helper.ConvertToStringPtr(req.Address), helper.ConvertToStringPtr(req.StateID))
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	return &emptypb.Empty{}, nil
}

func validateUpdateLandmarkRequest(req *api.UpdateLandmarkRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("ID")
	}
	return nil
}
