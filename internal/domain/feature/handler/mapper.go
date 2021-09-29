package handler

import (
	"features/internal/domain/feature/model"
	"time"
)

func mapFeaturesToCustomerFeatures(features []*model.Feature) []*CustomerFeature {
	res := make([]*CustomerFeature, 0, len(features))
	for _, feature := range features {
		res = append(res, mapFeatureToCustomerFeature(feature))
	}
	return res
}

func mapFeatureToCustomerFeature(feature *model.Feature) *CustomerFeature {
	currentTime := time.Now()
	return &CustomerFeature{
		Name:     feature.TechnicalName,
		Active:   feature.DeletedAt < currentTime.Unix(),
		Inverted: feature.Inverted,
		Expired:  feature.ExpiresOn < currentTime.Unix(),
	}
}

func mapFeatureToModel(request *featureRequestResponse) *model.Feature {
	return &model.Feature{
		ID:            request.ID,
		DisplayName:   request.DisplayName,
		TechnicalName: request.TechnicalName,
		ExpiresOn:     request.ExpiresOn,
		Description:   request.Description,
		Inverted:      request.Inverted,
		CustomerIDs:   request.CustomerIDs,
		CreatedAt:     request.CreatedAt,
		UpdatedAt:     request.UpdatedAt,
		DeletedAt:     request.DeletedAt,
	}
}

func mapModelsToResponse(features []*model.Feature) *featureResponse {
	res := make([]*featureRequestResponse, 0, len(features))
	for _, feature := range features {
		res = append(res, mapModelToFeature(feature))
	}
	return &featureResponse{Features: res}
}

func mapModelToFeature(feature *model.Feature) *featureRequestResponse {
	return &featureRequestResponse{
		ID:            feature.ID,
		DisplayName:   feature.DisplayName,
		TechnicalName: feature.TechnicalName,
		ExpiresOn:     feature.ExpiresOn,
		Description:   feature.Description,
		Inverted:      feature.Inverted,
		CustomerIDs:   feature.CustomerIDs,
		CreatedAt:     feature.CreatedAt,
		UpdatedAt:     feature.UpdatedAt,
		DeletedAt:     feature.DeletedAt,
	}
}
