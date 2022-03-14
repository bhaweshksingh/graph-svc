package app

import (
	"go.uber.org/zap"
	"graph-svc/pkg/config"
	"graph-svc/pkg/graphinfo"
	"graph-svc/pkg/http/router"
	"graph-svc/pkg/http/server"
	"graph-svc/pkg/reporters"
	"graph-svc/pkg/repository"
	"io"
	"net/http"
	"os"
)

func initHTTPServer(configFile string) {
	config := config.NewConfig(configFile)
	logger := initLogger(config)
	rt := initRouter(config, logger)

	server.NewServer(config, logger, rt).Start()
}

func initRouter(cfg config.Config, logger *zap.Logger) http.Handler {
	graphRepo := initRepository(cfg)
	graphService := initService(graphRepo)

	return router.NewRouter(logger, graphService)
}

func initService(graphRepository repository.GraphRepository) graphinfo.Service {
	eventService := graphinfo.NewGraphService(graphRepository)

	return eventService
}

func initRepository(cfg config.Config) repository.GraphRepository {
	return repository.NewGraphRepository(cfg.GetCSVPath())
}

func initLogger(cfg config.Config) *zap.Logger {
	return reporters.NewLogger(
		cfg.GetLogConfig().GetLevel(),
		getWriters(cfg.GetLogFileConfig())...,
	)
}

func getWriters(cfg config.LogFileConfig) []io.Writer {
	return []io.Writer{
		os.Stdout,
		reporters.NewExternalLogFile(cfg),
	}
}
