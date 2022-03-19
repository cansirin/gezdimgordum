package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/cansirin/gezdimgordum/graphql/graph/generated"
	"github.com/cansirin/gezdimgordum/graphql/graph/helper"
	"github.com/cansirin/gezdimgordum/graphql/graph/model"
)

func (r *userResolver) Landmarks(ctx context.Context, obj *model.User) ([]*model.Landmark, error) {
	landmarks, err := r.Backend.GetUserLandmarks(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	var allLandmarks []*model.Landmark
	for _, l := range landmarks {
		landmark := helper.NewLandmarkGQLModel(l)
		allLandmarks = append(allLandmarks, landmark)
	}
	return allLandmarks, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
