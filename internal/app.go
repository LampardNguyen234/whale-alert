package internal

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/config"
	"github.com/LampardNguyen234/whale-alert/db"
	"github.com/LampardNguyen234/whale-alert/internal/api"
	"github.com/LampardNguyen234/whale-alert/internal/clients"
	"github.com/LampardNguyen234/whale-alert/internal/listener"
	"github.com/LampardNguyen234/whale-alert/internal/processor"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/LampardNguyen234/whale-alert/webhook"
	"os"
	"os/signal"
	"syscall"
)

// App is the main application of the project.
type App struct {
	log        logger.Logger
	listener   *listener.Listener
	whm        *webhook.WebHookManager
	httpServer *api.HTTPServer
}

// NewApp creates a new main application.
func NewApp(cfg *config.Config) (*App, error) {
	if _, err := cfg.IsValid(); err != nil {
		return nil, err
	}

	// Setup logger
	var log logger.Logger
	if cfg.Logger.Color {
		log = logger.NewZeroLoggerWithColor(cfg.Logger.LogPath, "APP")
	} else {
		log = logger.NewZeroLogger(cfg.Logger.LogPath, "APP")
	}
	log.SetLogLevel(logger.LogLevel(cfg.Logger.Level))

	levelDb, err := db.NewLvlDB(cfg.LevelDB.Path)
	if err != nil {
		return nil, err
	}
	tmpStore := store.NewStore(levelDb)
	err = tmpStore.Init()
	if err != nil {
		return nil, err
	}

	tmpClients, err := clients.NewClientsFromConfig(cfg.Clients, tmpStore, log)
	if err != nil {
		return nil, err
	}

	whm, err := webhook.NewWebHookManagerFromConfig(cfg.Webhooks, log)
	if err != nil {
		return nil, err
	}

	processors, err := processor.NewProcessors(cfg.Processors, tmpClients, tmpStore, log)
	if err != nil {
		return nil, err
	}

	tmpListener, err := listener.NewListener(cfg.Listener, tmpClients, processors, tmpStore, log, *whm)
	if err != nil {
		return nil, err
	}

	httpServer, err := api.NewHTTPServer(tmpStore, log)
	if err != nil {
		return nil, err
	}

	return &App{
		log:        log.WithPrefix("App"),
		listener:   tmpListener,
		whm:        whm,
		httpServer: httpServer,
	}, nil
}

func (app *App) Start(ctx context.Context) {
	err := app.whm.Start()
	if err != nil {
		app.log.Errorf("failed to start whm: %v", err)
		return
	}
	go app.listener.Start(ctx)
	go app.httpServer.Start(ctx)

	sysErr := make(chan os.Signal, 1)
	signal.Notify(sysErr,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	select {
	case _ = <-ctx.Done():
		app.log.Infof("terminating due to ctx.Done")
		app.whm.Stop()
		return
	case sig := <-sysErr:
		app.log.Infof("terminating got `[%v]` signal", sig)
		return
	}
}
