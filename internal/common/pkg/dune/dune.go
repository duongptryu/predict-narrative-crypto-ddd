package dune

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type DuneClient struct {
	Client *resty.Client
	ApiKey string
}

func NewDuneClient(apiKey string, host string) *DuneClient {
	return &DuneClient{
		Client: resty.New().SetBaseURL(host),
		ApiKey: apiKey,
	}
}

func (d *DuneClient) QueryData(ctx context.Context, filter Filter, result any) error {
	path, err := filter.BuildRequestPath()
	if err != nil {
		return err
	}

	resp, err := d.Client.R().SetContext(ctx).SetHeader("X-Dune-API-Key", d.ApiKey).SetResult(result).Get(path)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("status code: %d \n%v", resp.StatusCode(), resp.Body())
	}

	return nil
}
