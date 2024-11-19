package adapters

import (
	"context"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
)

type sqlNarrativePerfRepository struct {
}

func NewSqlNarrativePerfRepository() *sqlNarrativePerfRepository {
	return &sqlNarrativePerfRepository{}
}

func (s sqlNarrativePerfRepository) CreateNarrativesPerf(ctx context.Context, data []domain.NarrativePerf) error {
	return nil
}
