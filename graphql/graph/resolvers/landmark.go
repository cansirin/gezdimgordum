package resolvers

//import (
//	"context"
//)
//
//func (r *landmarkResolver) CreateLandmark(ctx context.Context, input model.NewLandmark) (*model.Landmark, error) {
//	//if err := validateCreateLandmarkRequest(req); err != nil {
//	//	return nil, err
//	//}
//
//	result, err := r.Backend.CreateLandmark(ctx, input.Name, *input.Description, input.Address, input.State, input.UserID)
//	if err != nil {
//		return nil, err
//	}
//
//	landmark := model.Landmark{
//		ID:          result.ID.String(),
//		Name:        result.Name,
//		Slug:        result.Slug,
//		Description: result.Description,
//		Address:     result.Address,
//	}
//
//	return &landmark, nil
//}
