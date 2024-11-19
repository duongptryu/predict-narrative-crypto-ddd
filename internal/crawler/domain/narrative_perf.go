package domain

import "time"

type NarrativePerf struct {
	Id                  int
	Name                string
	Performance         float64
	AvgPerfAllNarrative float64
	RangeTime           RangeTimeNarrative
	LastUpdatedAt       time.Time
}

func NewNarrativePerf(name string, performance float64, RangeTime RangeTimeNarrative, avgPerfAllNarrative float64, lastUpdatedAt time.Time) *NarrativePerf {
	return &NarrativePerf{
		Name:                name,
		Performance:         performance,
		AvgPerfAllNarrative: avgPerfAllNarrative,
		RangeTime:           RangeTime,
		LastUpdatedAt:       lastUpdatedAt,
	}
}
