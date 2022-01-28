package server

import (
	"context"
	api "github.com/cansirin/gezdimgordum/landmark-api/rpc/landmark-api"
	"github.com/twitchtv/twirp"
)

func (s LandmarkAPIServer) CreateLandmark(ctx context.Context, req *api.CreateLandmarkRequest) (*api.Landmark, error) {
	if err := validateCreateLandmarkRequest(req); err != nil {
		return nil, err
	}

	landmark, err := s.backend.CreateLandmark(ctx, req.Name, req.Description, req.Address, req.StateId, req.UserId)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &api.Landmark{
		ID:          landmark.ID.String(),
		Name:        landmark.Name,
		Slug:        landmark.Slug,
		Address:     landmark.Address,
		Description: landmark.Description,
		StateId:     landmark.StateID,
		UserId:      landmark.UserID,
	}, nil
}

func validateCreateLandmarkRequest(req *api.CreateLandmarkRequest) error {
	if req.Name == "" {
		return twirp.RequiredArgumentError("name")
	}
	if req.StateId == "" {
		return twirp.RequiredArgumentError("state_id")
	}
	if req.Description == "" {
		return twirp.RequiredArgumentError("description")
	}
	if req.Address == "" {
		return twirp.RequiredArgumentError("address")
	}

	return nil
}
