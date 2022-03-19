package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cansirin/gezdimgordum/graphql/graph/generated"
	"github.com/cansirin/gezdimgordum/graphql/graph/helper"
	"github.com/cansirin/gezdimgordum/graphql/graph/model"
	"github.com/cansirin/gezdimgordum/graphql/internal/auth"
	"github.com/cansirin/gezdimgordum/graphql/pkg/jwt"
)

func (r *mutationResolver) CreateLandmark(ctx context.Context, input model.NewLandmark) (*model.Landmark, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	landmark, err := r.Backend.CreateLandmark(ctx, input.Name, *input.Description, input.Address, input.State, user.ID.String())
	if err != nil {
		return nil, err
	}

	return helper.NewLandmarkGQLModel(landmark), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := model.User{
		Name:     input.Username,
		Password: input.Password,
	}

	_, err := r.Backend.CreateUser(ctx, user.Name, user.Password)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := model.User{Name: input.Username, Password: input.Password}

	correct := r.Backend.AuthenticateUser(ctx, input.Username, input.Password)
	if !correct {
		return "", &helper.WrongUsernameOrPasswordError{}
	}

	token, err := jwt.GenerateToken(user.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input *model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Landmark(ctx context.Context, id string) (*model.Landmark, error) {
	landmark, err := r.Backend.GetLandmark(ctx, id)
	if err != nil {
		return nil, err
	}
	return helper.NewLandmarkGQLModel(landmark), nil
}

func (r *queryResolver) Landmarks(ctx context.Context) ([]*model.Landmark, error) {
	landmarks, err := r.Backend.GetAllLandmarks(ctx)

	if err != nil {
		return nil, err
	}

	var allLandmarks []*model.Landmark
	for _, landmark := range landmarks {
		l := helper.NewLandmarkGQLModel(landmark)
		allLandmarks = append(allLandmarks, l)
	}

	return allLandmarks, nil
}

func (r *queryResolver) LandmarksByState(ctx context.Context, state string) ([]*model.Landmark, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
