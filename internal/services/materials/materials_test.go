package materials_test

import (
	"appolo-register/internal/domain"
	"appolo-register/internal/services/materials"
	mock_materials "appolo-register/internal/services/materials/mocks"
	"appolo-register/internal/storage"
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestMaterialsService_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_materials.NewMockStorage(ctrl)

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	service := materials.New(logger, store)

	ctx := context.Background()

	m := &domain.Material{
		Name: "sample",
		Description: "lorem ispum",
		Url: "https://something.test/",
		Tags: []string{
			"Программирование",
			"РК",
		},
	}
	id := uint64(42)

	store.EXPECT().Save(ctx, gomock.Eq(m)).Return(id, nil)

	id, err := service.Register(ctx, m)
	require.NoError(t, err)
	require.Equal(t, id, id)
}

func TestMaterialsService_GetById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_materials.NewMockStorage(ctrl)

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	service := materials.New(logger, store)

	ctx := context.Background()

	id := uint64(42)
	m := &domain.Material{
		Id: id,
		Name: "sample",
		Description: "lorem ispum",
		Url: "https://something.test/",
		Tags: []string{
			"Программирование",
			"РК",
		},
	}

	store.EXPECT().MaterialById(ctx, gomock.Eq(id)).Return(m, nil)
	store.EXPECT().MaterialById(ctx, gomock.Not(gomock.Eq(id))).Return(nil, storage.ErrMaterialNotFound)

	t.Run("Material found", func(t *testing.T) {
		got, err := service.MaterialById(ctx, id)
		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, got, m)
	})

	t.Run("Material not found", func(t *testing.T) {
		_, err := service.MaterialById(ctx, id+1)
		require.ErrorIs(t, err, materials.ErrMaterialNotFound)
	})
}

func TestMaterialsService_GetByTags(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_materials.NewMockStorage(ctrl)

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	service := materials.New(logger, store)

	ctx := context.Background()

	m1 := &domain.Material{
		Id: uint64(10),
		Name: "sample",
		Description: "lorem ispum",
		Url: "https://something.test/",
		Tags: []string{
			"Программирование",
			"РК",
		},
	}
	m2 := &domain.Material{
		Id: uint64(11),
		Name: "sample",
		Description: "lorem ispum",
		Url: "https://something.test/",
		Tags: []string{
			"Программирование",
			"ЭКЗ",
		},
	}

	store.EXPECT().MaterialsByTags(ctx, gomock.InAnyOrder([]string{
		"Программирование",
	})).Return([]*domain.Material{
		m1, m2,
	}, nil)
	store.EXPECT().MaterialsByTags(ctx, gomock.InAnyOrder([]string{
		"История",
	})).Return([]*domain.Material{}, nil)

	t.Run("Materials found", func(t *testing.T) {
		got, err := service.MaterialsByTags(ctx, []string{ "Программирование" })
		require.NoError(t, err)
		require.NotNil(t, got)
		require.Len(t, got, 2)
		require.Contains(t, got, m1)
		require.Contains(t, got, m2)
	})

	t.Run("Materials not found", func(t *testing.T) {
		got, err := service.MaterialsByTags(ctx, []string{ "История" })
		require.NoError(t, err)
		require.NotNil(t, got)
		require.Len(t, got, 0)
	})
}
