package service

import (
	"database/sql"
	"features/internal/domain/feature/entity"
	"features/internal/domain/feature/model"
	"strings"
	"time"
)

func mapFeatureModelToEntity(model *model.Feature) *entity.Feature {
	return &entity.Feature{
		ID:            model.ID,
		DisplayName:   model.DisplayName,
		TechnicalName: model.TechnicalName,
		ExpiresOn:     sql.NullTime{Valid: model.ExpiresOn != 0, Time: time.Unix(model.ExpiresOn, 0)},
		Description:   model.Description,
		Inverted:      model.Inverted,
		CreatedAt:     time.Unix(model.CreatedAt, 0),
		UpdatedAt:     sql.NullTime{Valid: model.UpdatedAt != 0, Time: time.Unix(model.UpdatedAt, 0)},
		DeletedAt:     sql.NullTime{Valid: model.DeletedAt != 0, Time: time.Unix(model.DeletedAt, 0)},
	}
}

func mapEntitiesToModels(entities []*entity.Feature) []*model.Feature {
	res := make([]*model.Feature, 0, len(entities))
	for _, feature := range entities {
		res = append(res, mapEntityToModel(feature))
	}
	return res
}

func mapEntityToModel(entity *entity.Feature) *model.Feature {
	return &model.Feature{
		ID:            entity.ID,
		DisplayName:   entity.DisplayName,
		TechnicalName: entity.TechnicalName,
		ExpiresOn:     entity.ExpiresOn.Time.Unix(),
		Description:   entity.Description,
		Inverted:      entity.Inverted,
		CustomerIDs:   mapUserIDs(entity.CustomerIDs),
		CreatedAt:     entity.CreatedAt.Unix(),
		UpdatedAt:     entity.UpdatedAt.Time.Unix(),
		DeletedAt:     entity.UpdatedAt.Time.Unix(),
	}
}

func mapUserIDs(customerIDs string) []string {
	if customerIDs == "" {
		return nil
	}
	res := strings.Split(customerIDs, ",")
	return res
}
