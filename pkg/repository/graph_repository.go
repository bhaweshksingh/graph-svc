package repository

import (
	"context"
	"graph-svc/pkg/graphinfo/model"
	"time"
)

//go:generate moq -out mock/GraphRepository.go -pkg mock . GraphRepository
type GraphRepository interface {
	GetDataPoints(ctx context.Context) ([]model.DataPoint, error)
}

type gormGraphRepository struct {
	csvPath string
}

func (gbr *gormGraphRepository) GetDataPoints(ctx context.Context) ([]model.DataPoint, error) {
	var res []model.DataPoint
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()


	return res, nil
}

func NewGraphRepository(csvPath string) GraphRepository {
	return &gormGraphRepository{
		csvPath: csvPath,
	}
}
