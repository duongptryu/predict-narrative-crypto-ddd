package query

import (
	"context"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/common/decorator"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
	"github.com/sirupsen/logrus"
)

type CrawlNarrativeVolumnReq struct {
}

type CrawlNarrativeVolumnHandler decorator.QueryHandler[CrawlNarrativeVolumnReq, []domain.NarrativeVolume]

type CrawlerNarrativeVolumeAdapter interface {
	QueryNarrativeVolumeData(ctx context.Context) ([]domain.NarrativeVolume, error)
}

type crawlerNarrativeVolumeHandler struct {
	crawler CrawlerNarrativeVolumeAdapter
}

func NewCrawlerNarrativeVolumeHandler(crawler CrawlerNarrativeVolumeAdapter, logger *logrus.Entry, metricsClient decorator.MetricsClient) CrawlNarrativeVolumnHandler {
	return decorator.ApplyQueryDecorator[CrawlNarrativeVolumnReq, []domain.NarrativeVolume](
		crawlerNarrativeVolumeHandler{
			crawler: crawler,
		}, logger, metricsClient,
	)
}

func (h crawlerNarrativeVolumeHandler) Handle(ctx context.Context, query CrawlNarrativeVolumnReq) ([]domain.NarrativeVolume, error) {
	return h.crawler.QueryNarrativeVolumeData(ctx)
}
