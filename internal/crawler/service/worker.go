package service

import (
	"context"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/common/metrics"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/common/pkg/dune"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/adapters"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/app"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/app/command"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/app/query"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
	"github.com/sirupsen/logrus"
)

type worker struct {
	app app.Application
}

func NewWorker() worker {
	app := newApplication(context.Background())

	return worker{
		app: app,
	}
}

func (w worker) Run(ctx context.Context) {
	w.crawlNarrativePerfAndStore(ctx)
	w.crawlNarrativePointAndStore(ctx)
	w.crawlNarrativeVolumeAndStore(ctx)
}

func (w worker) crawlNarrativePerfAndStore(ctx context.Context) {
	var resultNarrativePerf = make([]domain.NarrativePerf, 0)
	// crawl narrative perf last 7 days
	result, err := w.app.Querys.CrawlNarrativesPerf.Handle(ctx, query.CrawlNarrativePerfReq{
		TimeRange: domain.Narrative7Days,
	})
	if err != nil {
		panic(err)
	}
	resultNarrativePerf = append(resultNarrativePerf, result...)

	// crawl narrative perf MTD
	result, err = w.app.Querys.CrawlNarrativesPerf.Handle(ctx, query.CrawlNarrativePerfReq{
		TimeRange: domain.NarrativeMTD,
	})
	if err != nil {
		panic(err)
	}
	resultNarrativePerf = append(resultNarrativePerf, result...)

	// crawl narrative perf last 3 months
	result, err = w.app.Querys.CrawlNarrativesPerf.Handle(ctx, query.CrawlNarrativePerfReq{
		TimeRange: domain.Narrative3Months,
	})
	if err != nil {
		panic(err)
	}
	resultNarrativePerf = append(resultNarrativePerf, result...)

	// crawl narrative perf YTD
	result, err = w.app.Querys.CrawlNarrativesPerf.Handle(ctx, query.CrawlNarrativePerfReq{
		TimeRange: domain.NarrativeMTD,
	})
	if err != nil {
		panic(err)
	}
	resultNarrativePerf = append(resultNarrativePerf, result...)

	// Insert intto database
	err = w.app.Commands.CreateNarrativesPerf.Handle(ctx, command.CreateNarrativesPerf{
		Data: resultNarrativePerf,
	})
	if err != nil {
		panic(err)
	}
}

func (w worker) crawlNarrativePointAndStore(ctx context.Context) {
	result, err := w.app.Querys.CrawlNarrativesPoint.Handle(ctx, query.CrawlNarrativePointReq{
		TimeRange: domain.Narrative6Months,
	})
	if err != nil {
		panic(err)
	}

	// Insert intto database
	err = w.app.Commands.CreateNarrativesPoint.Handle(ctx, command.CreateNarrativesPoint{
		Data: result,
	})
	if err != nil {
		panic(err)
	}
}

func (w worker) crawlNarrativeVolumeAndStore(ctx context.Context) {
	result, err := w.app.Querys.CrawlNarrativesVolume.Handle(ctx, query.CrawlNarrativeVolumnReq{})
	if err != nil {
		panic(err)
	}

	// Insert intto database
	err = w.app.Commands.CreateNarrativesVolume.Handle(ctx, command.CreateNarrativesVolume{
		Data: result,
	})
	if err != nil {
		panic(err)
	}
}

func newApplication(ctx context.Context) app.Application {

	narrativePeftRepo := adapters.NewSqlNarrativePerfRepository()
	narrativePointRepo := adapters.NewSqlNarrativePointRepository()
	narrativeVolumeRepo := adapters.NewNarrativeVolumeRepository()

	duneClient := dune.NewDuneClient("7pfUttO4tgrgxgJZUqQIDhwIXPAywY0A", "https://api.dune.com")
	crawlerNarrativeClient := adapters.NewDuneHttp(duneClient)

	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{
			CreateNarrativesPerf:   command.NewCreateNarrativesPerf(narrativePeftRepo, logger, metricsClient),
			CreateNarrativesPoint:  command.NewCreateNarrativePointHandler(narrativePointRepo, logger, metricsClient),
			CreateNarrativesVolume: command.NewCreateNarrativesVolume(narrativeVolumeRepo, logger, metricsClient),
		},
		Querys: app.Querys{
			CrawlNarrativesPerf:   query.NewCrawlNarrativePerfHandler(crawlerNarrativeClient, logger, metricsClient),
			CrawlNarrativesPoint:  query.NewCrawlerNarrativePointHandler(crawlerNarrativeClient, logger, metricsClient),
			CrawlNarrativesVolume: query.NewCrawlerNarrativeVolumeHandler(crawlerNarrativeClient, logger, metricsClient),
		},
	}
}
