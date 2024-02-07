package pgstorage_test

import (
	"appolo-register/internal/domain"
	s "appolo-register/internal/storage"
	"appolo-register/internal/storage/pgstorage"
	"context"
	"testing"

	"github.com/QuickDrone-Backend/pgconn" // TODO: заменить ссылку на свой форк
	"github.com/stretchr/testify/require"
)

var (
	cfg = pgconn.ConnConfig{
		Port: "32260",
		Host: "localhost",
		User: "postgres",
		Password: "postgres",
		DbName: "testdb",
		SslMode: "disable",
	}
)

func TestPostgresStorage_Save(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	t.Parallel()

	m := &domain.Material{
		Name: "sample",
		Description: "lorem ispum",
		Tags: []string{ "A", "B" },
		Url: "https://test.com/",
	}

	t.Run("Successfully saved material", func(t *testing.T) {
		t.Parallel()

		// Arrange
		db, err := pgconn.Connect(context.Background(), cfg.Url(), nil)
		require.NoError(t, err)
		storage := pgstorage.New(db)

		// Act
		id, err := storage.Save(context.Background(), m)

		// Assert
		require.NoError(t, err)

		got, err := storage.MaterialById(context.Background(), id)
		require.NoError(t, err)

		m.Id = id
		require.Equal(t, *m, *got)
	})	
}

func TestPostgresStorage_GetOrderById(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	t.Parallel()

	t.Run("Get an error if the material does not exist", func(t *testing.T) {
		t.Parallel()

		// Arrange
		db, err := pgconn.Connect(context.Background(), cfg.Url(), nil)
		require.NoError(t, err)
		storage := pgstorage.New(db)

		// Act
		_, err = storage.MaterialById(context.Background(), 0)

		// Assert
		require.ErrorIs(t, err, s.ErrMaterialNotFound)
	})
}

func TestPostgresStorage_GetOrdersFiltered(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	t.Parallel()

	t.Run("Return materails with the specified tags", func(t *testing.T) {
		t.Parallel()

		// Arrange
		db, err := pgconn.Connect(context.Background(), cfg.Url(), nil)
		require.NoError(t, err)
		storage := pgstorage.New(db)

		m1 := &domain.Material{
			Name: "sample_1",
			Description: "lorem ispum",
			Tags: []string{ "Программирование", "РК" },
			Url: "https://test.com/",
		}
		m2 := &domain.Material{
			Name: "sample_2",
			Description: "lorem ispum",
			Tags: []string{ "Программирование", "ЭКЗ" },
			Url: "https://test.com/",
		}
		m3 := &domain.Material{
			Name: "sample_3",
			Description: "lorem ispum",
			Tags: []string{ "Математический анализ", "ДЗ" },
			Url: "https://test.com/",
		}

		_, err = storage.Save(context.Background(), m1)
		require.NoError(t, err)
		_, err = storage.Save(context.Background(), m2)
		require.NoError(t, err)
		_, err = storage.Save(context.Background(), m3)
		require.NoError(t, err)

		// Act
		ms, err := storage.MaterialsByTags(context.Background(), []string{ "Программирование" })

		// Assert
		require.NoError(t, err)
		require.NotEmpty(t, ms)
		for _, m := range ms {
			require.Contains(t, m.Tags, "Программирование")
		}
	})
}
