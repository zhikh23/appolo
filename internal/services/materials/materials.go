package materials

import (
	"appolo-register/internal/domain"
	"appolo-register/internal/storage"
	"appolo-register/pkg/e"
	"context"
	"errors"
	"log/slog"
)

//go:generate mockgen -destination ./mocks/storage_mock.go "appolo-register/internal/services/materials" Storage
type Storage interface {
	Save(ctx context.Context, m *domain.Material) (id uint64, err error)
	MaterialById(ctx context.Context, id uint64) (m *domain.Material, err error)
	MaterialsByTags(ctx context.Context, tags []string) (ms []*domain.Material, err error)
}

type Service struct {
	logger  *slog.Logger
	Storage Storage
}

func New(logger *slog.Logger, storage Storage) *Service {
	return &Service{
		logger: logger,
		Storage: storage,
	}
}

func (s *Service) Register(ctx context.Context, m *domain.Material) (id uint64, err error) {
	const op = "services.materials.Service.Register"

	id, err = s.Storage.Save(ctx, m)
	if err != nil {
		err = e.WrapError(op, err)
		s.logger.Error(err.Error())
		return 0, err
	}

	return 
}

func (s *Service) MaterialById(ctx context.Context, id uint64) (m *domain.Material, err error) {
	const op = "services.materials.Service.MaterialById"

	m, err = s.Storage.MaterialById(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrMaterialNotFound) {
			return nil, ErrMaterialNotFound
		}
		err = e.WrapError(op, err)
		s.logger.Error(err.Error())
		return nil, err
	}

	return 
}

func (s *Service) MaterialsByTags(ctx context.Context, tags []string) (ms []*domain.Material, err error) {
	const op = "services.materials.Service.MaterialByTags"

	ms, err = s.Storage.MaterialsByTags(ctx, tags)
	if err != nil {
		err = e.WrapError(op, err)
		s.logger.Error(err.Error())
		return nil, err
	}

	return 
}
