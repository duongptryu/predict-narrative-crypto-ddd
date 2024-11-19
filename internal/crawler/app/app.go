package app

import (
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/app/command"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/app/query"
)

type Application struct {
	Commands Commands
	Querys   Querys
}

type Commands struct {
	CreateNarrativesPerf   command.CreateNarrativesPerfHandler
	CreateNarrativesPoint  command.CreateNarrativePointHandler
	CreateNarrativesVolume command.CreateNarrativesVolumeHandler
}

type Querys struct {
	CrawlNarrativesPerf   query.CrawlNarrativePerfHandler
	CrawlNarrativesPoint  query.CrawlNarrativePointHandler
	CrawlNarrativesVolume query.CrawlNarrativeVolumnHandler
}
