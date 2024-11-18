package dune

import (
	"errors"
	"fmt"
)

type Filter struct {
	APINumber string
	Limit     int
}

func (f *Filter) BuildRequestPath() (string, error) {
	if len(f.APINumber) == 0 {
		return "", errors.New("APINumber is required")
	}
	if f.Limit == 0 {
		f.Limit = 1000
	}
	return fmt.Sprintf("api/v1/query/%s/results?limit=%d", f.APINumber, f.Limit), nil
}
