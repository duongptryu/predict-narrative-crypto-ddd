package domain

import "context"

type NarrativePerfRepository interface {
	CreateNarrativesPerf(ctx context.Context, data []NarrativePerf) error
}

type NarrativePointRepository interface {
	CreateNarrativesPoint(ctx context.Context, data []NarrativePoint) error
}

type NarrativeVolumeRepository interface {
	CreateNarrativesVolume(ctx context.Context, data []NarrativeVolume) error
}
