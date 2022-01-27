package server

import (
	"context"
	api "github.com/gezdimgordum/landmark-api/rpc/landmark-api"
	"github.com/twitchtv/twirp"
)

func (s LandmarkAPIServer) CreateState(ctx context.Context, req *api.CreateStateRequest) (*api.CreateStateResponse, error) {
	if err := validateCreateStateRequest(req); err != nil {
		return nil, err
	}

	state, err := s.backend.CreateState(ctx, req.Name)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &api.CreateStateResponse{State: &api.State{ID: state.ID.String(), Name: state.Name, Slug: state.Slug}}, nil
}

func validateCreateStateRequest(req *api.CreateStateRequest) error {
	if req.Name == "" {
		return twirp.RequiredArgumentError("name")
	}
	return nil
}
