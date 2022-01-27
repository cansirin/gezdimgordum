package server

import (
	"context"
	api "github.com/gezdimgordum/landmark-api/rpc/landmark-api"
	"github.com/twitchtv/twirp"
)

func (s *LandmarkAPIServer) GetLandmark(ctx context.Context, req *api.GetLandmarkRequest) (*api.Landmark, error) {
	if err := validateGetLandmarkRequest(req); err != nil {
		return nil, err
	}

	landmark, err := s.backend.GetLandmark(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &api.Landmark{
		ID:          landmark.ID.String(),
		Name:        landmark.Name,
		Slug:        landmark.Slug,
		Address:     landmark.Address,
		StateId:     landmark.StateID,
		Description: landmark.Description,
	}, nil
}

func validateGetLandmarkRequest(req *api.GetLandmarkRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("ID")
	}
	return nil
}
