package http

import (
	"context"
	"net/http"

	"github.com/mxmntv/otus_hw/hw12_calendar/internal/model"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/logger"
)

const version = "/v1"

type EventUsecase interface {
	CreateEvent(ctx context.Context, event *model.Event) error
	UpdateEvent(ctx context.Context, event *model.Event) error
	GetList(ctx context.Context) ([]model.Event, error)
	DeleteEvent(ctx context.Context, id string) error
}

type EventHandler struct {
	usecase EventUsecase
	logger  logger.Interface
}

func NewEventHandler(u EventUsecase, l logger.Interface) EventHandler {
	return EventHandler{
		usecase: u,
		logger:  l,
	}
}

func (h EventHandler) Register(handler *http.ServeMux) { // как-то не очень выглядит
	handler.Handle(version+"/", loggingMiddleware(h.logger, http.HandlerFunc(h.heartbeat)))
}

func (h EventHandler) heartbeat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// func (h EventHandler) create(w http.ResponseWriter, r *http.Request) {
// 	// TODO implement
// }

// func (h EventHandler) update(w http.ResponseWriter, r *http.Request) {
// 	// TODO implement
// }

// func (h EventHandler) list(w http.ResponseWriter, r *http.Request) {
// 	// TODO implement
// }

// func (h EventHandler) delete(w http.ResponseWriter, r *http.Request) {
// 	// TODO implement
// }
