package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/samarec1812/segmentation-service/internal/app/repository/postgres/slug"
	"github.com/samarec1812/segmentation-service/internal/app/repository/postgres/user"
	"github.com/samarec1812/segmentation-service/internal/app/service"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	ht "github.com/samarec1812/segmentation-service/internal/app/ports/http"
	"github.com/samarec1812/segmentation-service/internal/config"
	"github.com/samarec1812/segmentation-service/internal/pkg/logger"
	"github.com/samarec1812/segmentation-service/internal/pkg/postgres"
)

func Run(cfg *config.Config) {

	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")
	log.Debug("debug message")

	db, err := postgres.Connect(cfg.DB_URL)
	if err != nil {
		log.Error("error connect database", err)
		os.Exit(1)
	}

	log.Info("database connect successful")
	slugRepo := slug.NewSlugRepository(db)
	userRepo := user.NewUserRepository(db)

	app := service.NewApp(slugRepo, userRepo)
	srv := ht.NewHTTPServer(cfg.Address, log, app)

	eg, ctx := errgroup.WithContext(context.Background())
	sigQuit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			log.Info("captured signal:", slog.String("signal", s.String()))
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	eg.Go(func() error {
		log.Info("starting http server, listening on:", slog.String("address", srv.Addr))
		defer log.Info("close http server listening on:", slog.String("address", srv.Addr))

		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err = srv.Shutdown(shCtx); err != nil {
				log.Error("can't close http server listening on %s: %s", slog.String("address", srv.Addr), slog.String("error", err.Error()))
				os.Exit(1)
			}

			close(errCh)
		}()

		go func() {
			if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err = <-errCh:
			return fmt.Errorf("http server can't listen and serve requests: %w", err)
		}
	})

	if err = eg.Wait(); err != nil {
		log.Info("gracefully shutting down the servers:", slog.String("error", err.Error()))
	}

}