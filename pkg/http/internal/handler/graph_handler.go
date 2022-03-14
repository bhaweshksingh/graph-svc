package handler

import (
	"context"
	"fmt"
	"graph-svc/pkg/graphinfo"
	"graph-svc/pkg/http/internal/utils"
	"net/http"

	"go.uber.org/zap"
)

type GraphHandler struct {
	lgr *zap.Logger
	svc graphinfo.Service
}

func NewGraphHandler(lgr *zap.Logger, svc graphinfo.Service) *GraphHandler {
	return &GraphHandler{
		lgr: lgr,
		svc: svc,
	}
}

func (gh *GraphHandler) GetChartData(resp http.ResponseWriter, req *http.Request) error {
	ctx := context.Background()

	chartResponse, err := gh.svc.GetChartData(ctx)
	if err != nil {
		errorMessage := fmt.Sprintf("error occurred while fetching chart data: %v", err)
		gh.lgr.Error(errorMessage)
		return fmt.Errorf(errorMessage)
	}
	utils.WriteSuccessResponse(resp, http.StatusOK, chartResponse)
	return nil
}
