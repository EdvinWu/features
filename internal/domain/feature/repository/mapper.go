package repository

import "features/internal/domain/feature/entity"

func mapFeatureToSQLMap(feature *entity.Feature) interface{} {
	return map[string]interface{}{
		"id":             feature.ID,
		"display_name":   feature.DisplayName,
		"technical_name": feature.TechnicalName,
		"expires_on":     feature.ExpiresOn,
		"description":    feature.Description,
		"inverted":       feature.Inverted,
		"created_at":     feature.CreatedAt,
		"updated_at":     feature.UpdatedAt,
		"deleted_at":     feature.DeletedAt,
	}
}
