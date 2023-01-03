package usecase

import (
	"context"

	"github.com/mxmntv/otus_hw/hw12_calendar/internal/model"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event *model.Event) error
	UpdateEvent(ctx context.Context, event *model.Event) error
	EventList(ctx context.Context) ([]model.Event, error)
	DeleteEvent(ctx context.Context, id string) error
}

type EventUsecase struct {
	repository EventRepository
}

func NewEventUsecase(r EventRepository) *EventUsecase {
	return &EventUsecase{r}
}

func (u EventUsecase) CreateEvent(ctx context.Context, event *model.Event) error {
	if err := u.repository.CreateEvent(ctx, event); err != nil {
		return err
	}
	return nil
}

func (u EventUsecase) UpdateEvent(ctx context.Context, event *model.Event) error {
	if err := u.repository.UpdateEvent(ctx, event); err != nil {
		return err
	}
	return nil
}

func (u EventUsecase) GetList(ctx context.Context) ([]model.Event, error) {
	events, err := u.repository.EventList(ctx)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (u EventUsecase) DeleteEvent(ctx context.Context, id string) error {
	if err := u.repository.DeleteEvent(ctx, id); err != nil {
		return err
	}
	return nil
}
