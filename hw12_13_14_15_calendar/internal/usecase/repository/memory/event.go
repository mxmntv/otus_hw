package repository

import (
	"context"
	"fmt"

	"github.com/mxmntv/otus_hw/hw12_calendar/internal/model"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/logger"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/memstorage"
)

type EventRepository struct {
	*memstorage.MemStorage
	*logger.Logger
}

func NewEventRepoIM(mstor *memstorage.MemStorage, l *logger.Logger) *EventRepository {
	return &EventRepository{mstor, l}
}

func (repo *EventRepository) CreateEvent(ctx context.Context, event *model.Event) error {
	select {
	case <-ctx.Done():
		// todo log timeout context
		return fmt.Errorf("context timeout has occurred")
	default:
		repo.Lock()
		defer repo.Unlock()
		_, ok := repo.Storage[event.ID]
		if !ok {
			repo.Storage[event.ID] = event
			return nil
		}
		return fmt.Errorf("event ID is busy")
	}
}

func (repo *EventRepository) UpdateEvent(ctx context.Context, event *model.Event) error {
	select {
	case <-ctx.Done():
		// todo log timeout context
		return fmt.Errorf("context timeout has occurred")
	default:
		repo.Lock()
		defer repo.Unlock()
		_, ok := repo.Storage[event.ID]
		if ok {
			repo.Storage[event.ID] = event
			return nil
		}
		return fmt.Errorf("event %s not found", event.ID)
	}
}

func (repo *EventRepository) EventList(ctx context.Context) ([]model.Event, error) {
	events := make([]model.Event, 0)
	select {
	case <-ctx.Done():
		// todo log timeout context
		return nil, fmt.Errorf("context timeout has occurred")
	default:
		repo.Lock()
		defer repo.Unlock()
		for _, v := range repo.Storage {
			vt := v.(*model.Event)
			events = append(events, *vt)
		}
		return events, nil
	}
}

func (repo *EventRepository) DeleteEvent(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		// todo log timeout context
		return fmt.Errorf("context timeout has occurred")
	default:
		_, ok := repo.Storage[id]
		repo.Lock()
		defer repo.Unlock()
		if ok {
			delete(repo.Storage, id)
			return nil
		}
		return fmt.Errorf("event not found")
	}
}
