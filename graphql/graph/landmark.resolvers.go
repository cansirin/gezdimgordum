package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cansirin/gezdimgordum/graphql/graph/generated"
	"github.com/cansirin/gezdimgordum/graphql/graph/helper"
	"github.com/cansirin/gezdimgordum/graphql/graph/model"
)

func (r *landmarkResolver) User(ctx context.Context, obj *model.Landmark) (*model.User, error) {
	user, err := r.Backend.GetUser(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}

	return helper.NewUserGQLModel(user), nil
}

// Landmark returns generated.LandmarkResolver implementation.
func (r *Resolver) Landmark() generated.LandmarkResolver { return &landmarkResolver{r} }

type landmarkResolver struct{ *Resolver }
