package server

import (
	"context"
	api "github.com/cansirin/gezdimgordum/landmark-api/rpc/landmark-api"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *LandmarkAPIServer) GetAllLandmarks(ctx context.Context, _ *emptypb.Empty) (*api.GetAllLandmarksResponse, error) {
	batch, err := s.backend.GetAllLandmarks(ctx)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	var landmarks []*api.Landmark
	for _, model := range batch {
		landmark := &api.Landmark{
			ID:          model.ID.String(),
			Name:        model.Name,
			Slug:        model.Slug,
			Address:     model.Address,
			StateId:     model.StateID,
			Description: model.Description,
			UserId:      model.UserID,
		}
		landmarks = append(landmarks, landmark)
	}

	return &api.GetAllLandmarksResponse{Landmarks: landmarks}, nil
}
