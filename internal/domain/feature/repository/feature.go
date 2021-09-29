package repository

import (
	"context"
	"features/internal/domain/feature/entity"
	"github.com/jmoiron/sqlx"
)

type Feature interface {
	FindByCustomerID(ctx context.Context, customerID string) ([]*entity.Feature, error)
	GetAll(ctx context.Context) ([]*entity.Feature, error)
	CreateFeature(ctx context.Context, feature *entity.Feature) error
	UpdateFeature(ctx context.Context, feature *entity.Feature) error
	ArchiveFeature(ctx context.Context, id string) error
	InsertUserIDs(ctx context.Context, featureID string, userIDs []string) error
}

type repository struct {
	db *sqlx.DB
}

func NewFeature(db *sqlx.DB) Feature {
	return &repository{db: db}
}
func (r *repository) FindByCustomerID(ctx context.Context, customerID string) ([]*entity.Feature, error) {
	var res []*entity.Feature
	err := r.db.Select(res, queryFindByCustomerID, customerID)
	return res, err
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.Feature, error) {
	var res []*entity.Feature
	err := r.db.Select(res, querySelectAll)
	return res, err
}

func (r *repository) CreateFeature(ctx context.Context, feature *entity.Feature) error {
	_, err := r.db.NamedExec(queryInsertFeature, mapFeatureToSQLMap(feature))
	return err
}

func (r *repository) UpdateFeature(ctx context.Context, feature *entity.Feature) error {
	_, err := r.db.NamedExec(queryUpdateFeature, mapFeatureToSQLMap(feature))
	return err
}

func (r *repository) ArchiveFeature(ctx context.Context, id string) error {
	_, err := r.db.Exec(queryArchiveFeature, id)
	return err
}

func (r *repository) InsertUserIDs(ctx context.Context, featureID string, ids []string) error {
	_, err := r.db.Exec(queryInsertUserIDs, featureID, ids)
	return err
}

const (
	queryFindByCustomerID = `select id,
       display_name,
       technical_name,
       expires_on,
       description,
       created_at,
       updated_at,
       deleted_at from features
join feature_users fu on features.id = fu.feature_id
where fu.user_id = ?`

	querySelectAll = `select id,
       display_name,
       technical_name,
       expires_on,
       description,
       created_at,
       updated_at,
       deleted_at,
       array_to_string(array_agg(fu.user_id), ',') customer_ids
from features
         left join feature_users fu on features.id = fu.feature_id
group by id, display_name, technical_name, expires_on, description, created_at, updated_at, deleted_at`

	queryInsertFeature = `
insert into features(id, display_name, technical_name, expires_on, description, created_at, updated_at, deleted_at) 
values (:id, :display_name, :technical_name, :expires_on, description:description, :created_at, :updated_at, :deleted_at)`

	queryUpdateFeature = `
update features
    set display_name = :display_name,
        technical_name = :technical_name,
		expires_on = :expires_on,
        description = :description,
        updated_at = :updated_at
where id = :id
`
	queryArchiveFeature = `
update features
	set deleted_at = :deleted_at
where id = :id`

	queryInsertUserIDs = `
insert into feature_users (feature_id, user_id) values (:feature_id,:user_id)
on conflict do nothing`
)
