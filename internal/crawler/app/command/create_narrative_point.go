package command

import (
	"context"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/common/decorator"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
	"github.com/sirupsen/logrus"
)

type CreateNarrativesPoint struct {
	Data []domain.NarrativePoint
}

type CreateNarrativePointHandler decorator.CommandHandler[CreateNarrativesPoint]

type createNarrativePointHandler struct {
	repo domain.NarrativePointRepository
}

func NewCreateNarrativePointHandler(repo domain.NarrativePointRepository, logger *logrus.Entry, metricsClient decorator.MetricsClient) CreateNarrativePointHandler {
	return decorator.ApplyCommandDecorator[CreateNarrativesPoint](
		createNarrativePointHandler{
			repo: repo,
		},
		logger,
		metricsClient,
	)
}

func (h createNarrativePointHandler) Handle(ctx context.Context, cmd CreateNarrativesPoint) error {
	return h.repo.CreateNarrativesPoint(ctx, cmd.Data)
}
