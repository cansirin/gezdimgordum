package helper

import (
	"github.com/cansirin/gezdimgordum/graphql/graph/model"
	"github.com/cansirin/gezdimgordum/graphql/internal/models"
)

func NewLandmarkGQLModel(landmark *models.Landmark) *model.Landmark {
	return &model.Landmark{
		ID:          landmark.ID.String(),
		Name:        landmark.Name,
		Slug:        landmark.Slug,
		Description: landmark.Description,
		Address:     landmark.Address,
		State:       landmark.State,
	}
}

func NewUserGQLModel(user *models.User) *model.User {
	return &model.User{
		ID:       user.ID.String(),
		Name:     user.Username,
		Password: user.Password,
	}
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}
