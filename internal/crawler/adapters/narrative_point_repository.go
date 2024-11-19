package adapters

import (
	"context"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
)

type sqlNarrativePointRepository struct {
}

func NewSqlNarrativePointRepository() *sqlNarrativePointRepository {
	return &sqlNarrativePointRepository{}
}

func (s sqlNarrativePointRepository) CreateNarrativesPoint(ctx context.Context, data []domain.NarrativePoint) error {
	return nil
}
