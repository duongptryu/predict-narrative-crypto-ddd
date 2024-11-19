package command

import (
	"context"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/common/decorator"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
	"github.com/sirupsen/logrus"
)

type CreateNarrativesPerf struct {
	Data []domain.NarrativePerf
}

type CreateNarrativesPerfHandler decorator.CommandHandler[CreateNarrativesPerf]

type createNarrativesPerfHandler struct {
	repo domain.NarrativePerfRepository
}

func NewCreateNarrativesPerf(repo domain.NarrativePerfRepository, logger *logrus.Entry, metricsClient decorator.MetricsClient) CreateNarrativesPerfHandler {
	return decorator.ApplyCommandDecorator[CreateNarrativesPerf](
		createNarrativesPerfHandler{
			repo: repo,
		},
		logger,
		metricsClient,
	)
}

func (h createNarrativesPerfHandler) Handle(ctx context.Context, cmd CreateNarrativesPerf) error {
	return h.repo.CreateNarrativesPerf(ctx, cmd.Data)
}
