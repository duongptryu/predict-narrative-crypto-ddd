package domain

import "time"

type NarrativePoint struct {
	Id            int
	Name          string
	Point         int
	RangeTime     RangeTimeNarrative
	LastUpdatedAt time.Time
}

func NewNarrativePoint(name string, point int, RangeTime RangeTimeNarrative, lastUpdatedAt time.Time) *NarrativePoint {
	return &NarrativePoint{
		Name:          name,
		Point:         point,
		RangeTime:     RangeTime,
		LastUpdatedAt: lastUpdatedAt,
	}
}
