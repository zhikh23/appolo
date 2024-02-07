package pgstorage

import (
	"appolo-register/internal/domain"
	"appolo-register/internal/storage"
	"appolo-register/pkg/e"
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		Db: db,
	}
}

func (s *Storage) Save(ctx context.Context, m *domain.Material) (uint64, error) {
	const op = "storage.pgstorage.Storage.Save"

	var id uint64
	err := s.Db.GetContext(ctx, &id, save_material_query, m.Name, m.Description, m.Tags, m.Url)
	return id, e.WrapIfError(op, err)
}

func (s *Storage) MaterialById(ctx context.Context, id uint64) (*domain.Material, error) {
	const op = "storage.pgstorage.Storage.MaterialById"

	var dto MaterialDto
	err := s.Db.GetContext(ctx, &dto, select_by_id, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrMaterialNotFound
		}
		return nil, e.WrapIfError(op, err)
	}
	return convertDtoToModel(&dto), nil
}

func (s *Storage) MaterialsByTags(ctx context.Context, tags []string) ([]*domain.Material, error) {
	const op = "storage.pgstorage.Storage.MaterialsByTags"

	var dtoArr []MaterialDto
	err := s.Db.SelectContext(ctx, &dtoArr, select_by_tags, tags)
	if err != nil {
		return nil, e.WrapIfError(op, err)
	}

	return convertDtosToModels(dtoArr), nil
}

func convertDtoToModel(dto *MaterialDto) *domain.Material {
	return &domain.Material{
		Id: dto.Id,
		Name: dto.Name,
		Description: dto.Description,
		Tags: dto.Tags,
		Url: dto.Url,
	}
}

func convertDtosToModels(dtoArr []MaterialDto) []*domain.Material {
	models := make([]*domain.Material, len(dtoArr))
	for i, dto := range dtoArr {
		models[i] = convertDtoToModel(&dto)
	}
	return models
}
