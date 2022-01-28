package server

import (
	"context"
	api "github.com/cansirin/gezdimgordum/landmark-api/rpc/landmark-api"
	"github.com/twitchtv/twirp"
)

func (s *LandmarkAPIServer) GetLandmarksByStateID(ctx context.Context, req *api.GetLandmarksByStateIDRequest) (*api.GetLandmarksByStateIDResponse, error) {
	if err := validateGetLandmarksByStateIDRequest(req); err != nil {
		return nil, err
	}

	batch, err := s.backend.GetLandmarksByStateID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	var landmarks []*api.Landmark
	for _, l := range batch {
		landmark := &api.Landmark{
			ID:          l.ID.String(),
			Name:        l.Name,
			Slug:        l.Slug,
			Address:     l.Address,
			StateId:     l.StateID,
			Description: l.Description,
		}
		landmarks = append(landmarks, landmark)
	}

	return &api.GetLandmarksByStateIDResponse{Landmarks: landmarks}, nil
}

func validateGetLandmarksByStateIDRequest(req *api.GetLandmarksByStateIDRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("state_id")
	}
	return nil
}
