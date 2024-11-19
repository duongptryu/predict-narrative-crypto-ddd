package query

import (
	"context"
	"errors"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/common/decorator"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
	"github.com/sirupsen/logrus"
)

type CrawlNarrativePointReq struct {
	TimeRange domain.RangeTimeNarrative
}

type CrawlNarrativePointHandler decorator.QueryHandler[CrawlNarrativePointReq, []domain.NarrativePoint]

type CrawlerNarrativePointAdapter interface {
	QueryNarrative6MonthsPointData(ctx context.Context) ([]domain.NarrativePoint, error)
}

type crawlerNarrativePointHandler struct {
	crawler CrawlerNarrativePointAdapter
}

func NewCrawlerNarrativePointHandler(crawler CrawlerNarrativePointAdapter, logger *logrus.Entry, metricsClientt decorator.MetricsClient) CrawlNarrativePointHandler {
	return decorator.ApplyQueryDecorator[CrawlNarrativePointReq, []domain.NarrativePoint](
		crawlerNarrativePointHandler{
			crawler: crawler,
		}, logger, metricsClientt,
	)
}

func (h crawlerNarrativePointHandler) Handle(ctx context.Context, query CrawlNarrativePointReq) ([]domain.NarrativePoint, error) {
	switch query.TimeRange {
	case domain.Narrative6Months:
		return h.crawler.QueryNarrative6MonthsPointData(ctx)
	default:
		return nil, errors.New("invalid time range")
	}
}
