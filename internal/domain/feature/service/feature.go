package service

import (
	"context"
	"features/internal/domain/feature/model"
	"features/internal/domain/feature/repository"
	"github.com/sirupsen/logrus"
	"time"
)

type Feature interface {
	FindByCustomerID(ctx context.Context, customerID string) ([]*model.Feature, error)
	GetAll(ctx context.Context) ([]*model.Feature, error)
	CreateFeature(ctx context.Context, feature *model.Feature) error
	UpdateFeature(ctx context.Context, feature *model.Feature) error
	ArchiveFeature(ctx context.Context, id string) error
}

type service struct {
	repo repository.Feature
	log  *logrus.Entry
}

func NewFeature(repo repository.Feature, log *logrus.Entry) Feature {
	return &service{repo: repo, log: log}
}

func (s *service) FindByCustomerID(ctx context.Context, customerID string) ([]*model.Feature, error) {
	res, err := s.repo.FindByCustomerID(ctx, customerID)
	if err != nil {
		s.log.WithError(err).Error("Failed to find features by customer id")
		return nil, err
	}
	return mapEntitiesToModels(res), nil
}

func (s *service) GetAll(ctx context.Context) ([]*model.Feature, error) {
	features, err := s.repo.GetAll(ctx)
	if err != nil {
		s.log.WithError(err).Error("Failed to get all features")
		return nil, err
	}
	return mapEntitiesToModels(features), nil
}

func (s *service) CreateFeature(ctx context.Context, feature *model.Feature) error {
	err := s.repo.CreateFeature(ctx, mapFeatureModelToEntity(feature))
	if err != nil {
		s.log.WithField("feature", feature).WithError(err).Error("Failed to create feature")
		return err
	}
	if len(feature.CustomerIDs) > 0 {
		err := s.repo.InsertUserIDs(ctx, feature.ID, feature.CustomerIDs)
		if err != nil {
			s.log.WithField("feature", feature).WithError(err).Error("Failed to create feature")
			return err
		}
	}
	return nil
}

func (s *service) UpdateFeature(ctx context.Context, feature *model.Feature) error {
	feature.UpdatedAt = time.Now().Unix()
	err := s.repo.UpdateFeature(ctx, mapFeatureModelToEntity(feature))
	if err != nil {
		s.log.WithField("feature", feature).WithError(err).Error("Failed to update feature")
		return err
	}
	if len(feature.CustomerIDs) > 0 {
		err := s.repo.InsertUserIDs(ctx, feature.ID, feature.CustomerIDs)
		if err != nil {
			s.log.WithField("feature", feature).WithError(err).Error("Failed to create feature")
			return err
		}
	}
	return nil
}

func (s *service) ArchiveFeature(ctx context.Context, id string) error {
	err := s.repo.ArchiveFeature(ctx, id)
	if err != nil {
		s.log.WithField("feature_id", id).WithError(err).Error("Failed to archive feature")
		return err
	}
	return nil
}
