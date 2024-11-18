package adapters

import (
	"context"
	"time"

	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/common/pkg/dune"
	"github.com/duongptryu/predict-narrative-crypto-ddd/internal/crawler/domain"
)

type DuneHttp struct {
	client *dune.DuneClient
}

func NewDuneHttp(client *dune.DuneClient) *DuneHttp {
	return &DuneHttp{
		client: client,
	}
}

type NarrativeLast7DaysPerfModel struct {
	Result struct {
		Rows []struct {
			AvgLast7DPercentAcrossAllNarratives float64 `json:"avg_last_7d_percent_across_all_narratives"`
			Last7DPercent                       float64 `json:"last_7d_percent"`
			Narrative                           string  `json:"narrative"`
		}
	} `json:"result"`
}

type NarrativeMTDsPerfModel struct {
	Result struct {
		Rows []struct {
			AvgMTDPerfAcrossAllNarratives float64 `json:"avg_mtd_perf_across_all_narratives"`
			MTDPerf                       float64 `json:"mtd_perf"`
			Narrative                     string  `json:"narrative"`
		}
	} `json:"result"`
}

type NarrativeLast3MonthsPerfModel struct {
	Result struct {
		Rows []struct {
			AvgQuarterPerfAcrossAllNarratives float64 `json:"avg_quarter_perf_across_all_narratives"`
			Narrative                         string  `json:"narrative"`
			QuarterPerf                       float64 `json:"quarter_perf"`
		}
	} `json:"result"`
}

type NarrativeYTDsPerfModel struct {
	Result struct {
		Rows []struct {
			AvgYTDPerfAcrossAllNarratives float64 `json:"avg_ytd_perf_across_all_narratives"`
			Narrative                     string  `json:"narrative"`
			YTDPerf                       float64 `json:"ytd_perf"`
		}
	} `json:"result"`
}

type Narrative6MonthsPointModel struct {
	Result struct {
		Rows []struct {
			Narrative string `json:"narrative"`
			Points    int    `json:"points"`
		}
	} `json:"result"`
}

// 3276483
type NarrativeVolumeModel struct {
	Result struct {
		Rows []struct {
			Narrative         string `json:"narrative"`
			AvgVolumePerAsset int    `json:"avg_volume_per_asset"`
		}
	} `json:"result"`
}

func (d *DuneHttp) QueryLast7DaysNarrativePerfData(ctx context.Context) ([]domain.NarrativePerf, error) {
	var resp NarrativeLast7DaysPerfModel
	if err := d.client.QueryData(ctx, dune.Filter{
		APINumber: "3619448",
		Limit:     1000,
	}, &resp); err != nil {
		return nil, err
	}

	timeNow := time.Now()
	var result = make([]domain.NarrativePerf, len(resp.Result.Rows))
	for i, row := range resp.Result.Rows {
		result[i] = *domain.NewNarrativePerf(row.Narrative, row.Last7DPercent, domain.Narrative7Days, row.AvgLast7DPercentAcrossAllNarratives, timeNow)
	}

	return result, nil
}

func (d *DuneHttp) QueryMTDNarrativePerfData(ctx context.Context) ([]domain.NarrativePerf, error) {
	var resp NarrativeMTDsPerfModel
	if err := d.client.QueryData(ctx, dune.Filter{
		APINumber: "3242787",
		Limit:     1000,
	}, &resp); err != nil {
		return nil, err
	}

	timeNow := time.Now()
	var result = make([]domain.NarrativePerf, len(resp.Result.Rows))
	for i, row := range resp.Result.Rows {
		result[i] = *domain.NewNarrativePerf(row.Narrative, row.MTDPerf, domain.Narrative7Days, row.AvgMTDPerfAcrossAllNarratives, timeNow)
	}

	return result, nil
}

func (d *DuneHttp) Query3MonthsNarrativePerfData(ctx context.Context) ([]domain.NarrativePerf, error) {
	var resp NarrativeLast3MonthsPerfModel
	if err := d.client.QueryData(ctx, dune.Filter{
		APINumber: "3318044",
		Limit:     1000,
	}, &resp); err != nil {
		return nil, err
	}

	timeNow := time.Now()
	var result = make([]domain.NarrativePerf, len(resp.Result.Rows))
	for i, row := range resp.Result.Rows {
		result[i] = *domain.NewNarrativePerf(row.Narrative, row.QuarterPerf, domain.Narrative7Days, row.AvgQuarterPerfAcrossAllNarratives, timeNow)
	}

	return result, nil
}

func (d *DuneHttp) QueryYTDNarrativePerfData(ctx context.Context) ([]domain.NarrativePerf, error) {
	var resp NarrativeYTDsPerfModel
	if err := d.client.QueryData(ctx, dune.Filter{
		APINumber: "3242851",
		Limit:     1000,
	}, &resp); err != nil {
		return nil, err
	}

	timeNow := time.Now()
	var result = make([]domain.NarrativePerf, len(resp.Result.Rows))
	for i, row := range resp.Result.Rows {
		result[i] = *domain.NewNarrativePerf(row.Narrative, row.YTDPerf, domain.Narrative7Days, row.AvgYTDPerfAcrossAllNarratives, timeNow)
	}

	return result, nil
}

func (d *DuneHttp) QueryNarrative6MonthsPointData(ctx context.Context) ([]domain.NarrativePoint, error) {
	var resp Narrative6MonthsPointModel
	if err := d.client.QueryData(ctx, dune.Filter{
		APINumber: "3486556",
		Limit:     1000,
	}, &resp); err != nil {
		return nil, err
	}

	timeNow := time.Now()
	var result = make([]domain.NarrativePoint, len(resp.Result.Rows))
	for i, row := range resp.Result.Rows {
		result[i] = *domain.NewNarrativePoint(row.Narrative, row.Points, domain.Narrative6Months, timeNow)
	}

	return result, nil
}

func (d *DuneHttp) QueryNarrativeVolumeData(ctx context.Context) ([]domain.NarrativeVolume, error) {
	var resp Narrative6MonthsPointModel
	if err := d.client.QueryData(ctx, dune.Filter{
		APINumber: "3276483",
		Limit:     1000,
	}, &resp); err != nil {
		return nil, err
	}

	timeNow := time.Now()
	var result = make([]domain.NarrativeVolume, len(resp.Result.Rows))
	for i, row := range resp.Result.Rows {
		result[i] = *domain.NewNarrativeVolume(row.Narrative, row.Points, domain.Narrative6Months, timeNow)
	}

	return result, nil
}
