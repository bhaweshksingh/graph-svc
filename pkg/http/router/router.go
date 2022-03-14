package router

import (
	"graph-svc/pkg/graphinfo"
	"graph-svc/pkg/http/internal/handler"
	"graph-svc/pkg/http/internal/middleware"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewRouter(lgr *zap.Logger, eventsService graphinfo.Service) http.Handler {
	router := mux.NewRouter()
	router.Use(handlers.RecoveryHandler())

	eventsHandler := handler.NewGraphHandler(lgr, eventsService)

	router.
		HandleFunc("/chart-data", withMiddlewares(lgr, middleware.WithErrorHandler(lgr, eventsHandler.GetChartData))).
		Methods(http.MethodGet)

	return router
}

func withMiddlewares(lgr *zap.Logger, hnd http.HandlerFunc) http.HandlerFunc {
	return middleware.WithSecurityHeaders(middleware.WithReqResLog(lgr, hnd))
}
