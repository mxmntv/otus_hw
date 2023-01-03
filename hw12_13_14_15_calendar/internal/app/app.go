package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mxmntv/otus_hw/hw12_calendar/config"
	handler "github.com/mxmntv/otus_hw/hw12_calendar/internal/controller/http/v1"
	"github.com/mxmntv/otus_hw/hw12_calendar/internal/usecase"
	repository_mem "github.com/mxmntv/otus_hw/hw12_calendar/internal/usecase/repository/memory"
	repository "github.com/mxmntv/otus_hw/hw12_calendar/internal/usecase/repository/pg"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/httpserver"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/logger"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/memstorage"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/postgres"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.Log.Level)

	var repo usecase.EventRepository

	switch s := cfg.App.Storage; s {
	case "memory":
		storage := memstorage.NewMemStorage()
		repo = repository_mem.NewEventRepoIM(storage, logger)
	case "pg":
		pg, err := postgres.NewPgConnection(cfg.PG.URL)
		if err != nil {
			logger.Fatal("PG error: ", fmt.Errorf("app - run - initial postgress connection err: %w", err))
		}
		repo = *repository.NewEventRepoPg(pg, logger)
	}

	eventUseCase := usecase.NewEventUsecase(repo)

	mux := http.NewServeMux()
	handle := handler.NewEventHandler(eventUseCase, logger)
	handle.Register(mux)

	server := httpserver.NewServer(cfg.HTTP.Host, cfg.HTTP.Port, mux)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logger.Error("failed to stop http server: " + err.Error())
		}
	}()

	logger.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logger.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
