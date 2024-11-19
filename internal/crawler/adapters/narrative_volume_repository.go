package adapters

import (
	"context"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
)

type NarrativeVolumeRepository struct{}

func NewNarrativeVolumeRepository() *NarrativeVolumeRepository {
	return &NarrativeVolumeRepository{}
}

func (s NarrativeVolumeRepository) CreateNarrativesVolume(ctx context.Context, data []domain.NarrativeVolume) error {
	return nil
}
