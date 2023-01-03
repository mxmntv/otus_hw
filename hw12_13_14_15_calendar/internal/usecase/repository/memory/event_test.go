package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/mxmntv/otus_hw/hw12_calendar/internal/model"
	repository "github.com/mxmntv/otus_hw/hw12_calendar/internal/usecase/repository/memory"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/memstorage"
	"github.com/stretchr/testify/require"
)

func TestMemStorage(t *testing.T) {
	var storage *memstorage.MemStorage
	var repo repository.EventRepository
	data := []model.Event{
		{
			ID:               "1",
			Title:            "Example 1",
			DateTime:         time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC),
			Duration:         30,
			Description:      "Example description 1",
			OwnerID:          "1",
			NotificationTime: time.Date(2023, 1, 2, 11, 45, 0, 0, time.UTC),
		},
		{
			ID:               "2",
			Title:            "Example 2",
			DateTime:         time.Date(2023, 1, 2, 11, 0, 0, 0, time.UTC),
			Duration:         30,
			Description:      "Example description 2",
			OwnerID:          "1",
			NotificationTime: time.Date(2023, 1, 2, 10, 45, 0, 0, time.UTC),
		},
		{
			ID:               "3",
			Title:            "Example 3",
			DateTime:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
			Duration:         30,
			Description:      "Example description 3",
			OwnerID:          "1",
			NotificationTime: time.Date(2023, 1, 2, 9, 45, 0, 0, time.UTC),
		},
	}

	t.Run("create_storage", func(t *testing.T) {
		storage = memstorage.NewMemStorage()
		storage.Storage["1"] = data[0]
		res := storage.Storage["1"]
		require.Equal(t, "Example 1", res.(model.Event).Title)
		delete(storage.Storage, "1")
		require.Nil(t, storage.Storage["1"])
	})

	t.Run("add events", func(t *testing.T) {
		repo = *repository.NewEventRepoIM(storage, nil)
		for _, v := range data {
			repo.CreateEvent(context.Background(), &v)
		}
		busy := model.Event{
			ID:               "3",
			Title:            "Example 3",
			DateTime:         time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC),
			Duration:         30,
			Description:      "Example description",
			OwnerID:          "1",
			NotificationTime: time.Date(2023, 1, 2, 11, 45, 0, 0, time.UTC),
		}
		err := repo.CreateEvent(context.Background(), &busy)
		require.Error(t, err, "event ID is busy")
	})

	t.Run("update event", func(t *testing.T) {
		candidate1 := model.Event{
			ID:               "3",
			Title:            "Updated event",
			DateTime:         time.Date(2023, 2, 1, 10, 0, 0, 0, time.UTC),
			Duration:         45,
			Description:      "Updated description",
			OwnerID:          "1",
			NotificationTime: time.Date(2023, 2, 1, 9, 45, 0, 0, time.UTC),
		}
		err := repo.UpdateEvent(context.Background(), &candidate1)
		require.NoError(t, err, "event not found")

		candidate2 := model.Event{
			ID:               "35",
			Title:            "Updated event",
			DateTime:         time.Date(2023, 2, 1, 10, 0, 0, 0, time.UTC),
			Duration:         45,
			Description:      "Updated description",
			OwnerID:          "1",
			NotificationTime: time.Date(2023, 2, 1, 9, 45, 0, 0, time.UTC),
		}
		err = repo.UpdateEvent(context.Background(), &candidate2)
		require.Error(t, err, "event not found")

		event := storage.Storage["3"]
		updatedTite := event.(*model.Event).Title
		require.Equal(t, "Updated event", updatedTite)
	})

	t.Run("event list", func(t *testing.T) {
		_, err := repo.EventList(context.Background())
		require.NoError(t, err)
	})

	t.Run("delete event", func(t *testing.T) {
		err := repo.DeleteEvent(context.Background(), "10")
		require.Error(t, err, "event not found")
		err = repo.DeleteEvent(context.Background(), "3")
		require.NoError(t, err, "event not found")
	})
}
