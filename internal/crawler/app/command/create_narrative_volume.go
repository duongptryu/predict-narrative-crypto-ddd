package command

import (
	"context"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/common/decorator"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
	"github.com/sirupsen/logrus"
)

type CreateNarrativesVolume struct {
	Data []domain.NarrativeVolume
}

type CreateNarrativesVolumeHandler decorator.CommandHandler[CreateNarrativesVolume]

type createNarrativesVolumeHandler struct {
	repo domain.NarrativeVolumeRepository
}

func NewCreateNarrativesVolume(repo domain.NarrativeVolumeRepository, logger *logrus.Entry, metricsClient decorator.MetricsClient) CreateNarrativesVolumeHandler {
	return decorator.ApplyCommandDecorator[CreateNarrativesVolume](
		createNarrativesVolumeHandler{
			repo: repo,
		},
		logger,
		metricsClient,
	)
}

func (h createNarrativesVolumeHandler) Handle(ctx context.Context, cmd CreateNarrativesVolume) error {
	return h.repo.CreateNarrativesVolume(ctx, cmd.Data)
}
