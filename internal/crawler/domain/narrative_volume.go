package domain

import "time"

type NarrativeVolume struct {
	Id            int
	Name          string
	Volume        int
	RangeTime     RangeTimeNarrative
	LastUpdatedAt time.Time
}

func NewNarrativeVolume(name string, volume int, RangeTime RangeTimeNarrative, lastUpdatedAt time.Time) *NarrativeVolume {
	return &NarrativeVolume{
		Name:          name,
		Volume:        volume,
		RangeTime:     RangeTime,
		LastUpdatedAt: lastUpdatedAt,
	}
}
