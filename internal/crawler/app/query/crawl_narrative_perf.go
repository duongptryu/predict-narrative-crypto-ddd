package query

import (
	"context"
	"errors"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/common/decorator"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
	"github.com/sirupsen/logrus"
)

type CrawlNarrativePerfReq struct {
	TimeRange domain.RangeTimeNarrative
}

type CrawlerNarrativePerfAdapter interface {
	QueryLast7DaysNarrativePerfData(ctx context.Context) ([]domain.NarrativePerf, error)
	QueryMTDNarrativePerfData(ctx context.Context) ([]domain.NarrativePerf, error)
	Query3MonthsNarrativePerfData(ctx context.Context) ([]domain.NarrativePerf, error)
	QueryYTDNarrativePerfData(ctx context.Context) ([]domain.NarrativePerf, error)
}

type CrawlNarrativePerfHandler decorator.QueryHandler[CrawlNarrativePerfReq, []domain.NarrativePerf]

type crawlNarrativePerfHandler struct {
	crawler CrawlerNarrativePerfAdapter
}

func NewCrawlNarrativePerfHandler(crawler CrawlerNarrativePerfAdapter, logger *logrus.Entry, metricsClient decorator.MetricsClient) CrawlNarrativePerfHandler {
	return decorator.ApplyQueryDecorator[CrawlNarrativePerfReq, []domain.NarrativePerf](
		crawlNarrativePerfHandler{
			crawler: crawler,
		}, logger, metricsClient)
}

func (h crawlNarrativePerfHandler) Handle(ctx context.Context, query CrawlNarrativePerfReq) ([]domain.NarrativePerf, error) {
	switch query.TimeRange {
	case domain.Narrative7Days:
		return h.crawler.QueryLast7DaysNarrativePerfData(ctx)
	case domain.NarrativeMTD:
		return h.crawler.Query3MonthsNarrativePerfData(ctx)
	case domain.Narrative3Months:
		return h.crawler.Query3MonthsNarrativePerfData(ctx)
	case domain.NarrativeYTD:
		return h.crawler.QueryYTDNarrativePerfData(ctx)
	default:
		return nil, errors.New("invalid time range")
	}
}
