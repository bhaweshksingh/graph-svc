package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"graph-svc/pkg/graphinfo/model"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const (
	validCSVFileData = `id,timestamp,residentId,location
1,2022-01-21T05:53:19Z,resident0001,First_Floor_Bathroom
1,2022-01-21T05:56:19Z,resident0001,First_Floor_Bathroom
2,2021-06-13T02:21:48Z,resident0001,First_Floor_Bathroom
`
	invalidCSVFileData = ""
)

func setUp(t *testing.T, csvFileData string) (string, context.Context) {
	uniqueTempDir, err := os.MkdirTemp(os.TempDir(), "*-testData")
	if err != nil {
		assert.NoError(t, err, "expected temp directory creation to succeed.")
	}

	filePath := filepath.Join(uniqueTempDir, "visit_data.csv")
	csvFile, err := os.Create(filePath)
	if err != nil {
		assert.NoError(t, err, "expected csvFile creation succeed.")
	}
	defer csvFile.Close()

	_, err = csvFile.Write([]byte(csvFileData))
	if err != nil {
		assert.NoError(t, err, "expected csvFile writing to succeed.")
	}

	return filePath, context.Background()
}

func TestGraphRepository_GetDataPoints(t *testing.T) {
	filePath, ctx := setUp(t, validCSVFileData)
	expectedPoints := []model.DataPoint{
		{Day: time.Date(2021, time.June, 13, 0, 0, 0, 0, time.UTC), VisitCount: 1},
		{Day: time.Date(2022, time.January, 21, 0, 0, 0, 0, time.UTC), VisitCount: 2},
	}

	graphRepo := NewGraphRepository(filePath)

	actualPoints, err := graphRepo.GetDataPoints(ctx)

	assert.Nil(t, err)
	assert.Equal(t, expectedPoints, actualPoints)

}

func TestGraphRepository_GetDataPoints_fails(t *testing.T) {
	filePath, ctx := setUp(t, invalidCSVFileData)
	expectedErrorMessage := "failed to parse visit data from CSV"

	graphRepo := NewGraphRepository(filePath)

	actualPoints, err := graphRepo.GetDataPoints(ctx)

	assert.Equal(t, 0, len(actualPoints))
	assert.Contains(t, err.Error(), expectedErrorMessage)

}

func TestGraphRepository_GetDataPoints_fails_for_missing_csv(t *testing.T) {
	ctx := context.Background()
	expectedErrorMessage := "failed to read CSV file from path"

	graphRepo := NewGraphRepository("invalid-csv-filepath")

	actualPoints, err := graphRepo.GetDataPoints(ctx)

	assert.Equal(t, 0, len(actualPoints))
	assert.Contains(t, err.Error(), expectedErrorMessage)

}
