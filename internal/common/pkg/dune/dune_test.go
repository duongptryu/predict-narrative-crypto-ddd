package dune

import (
	"context"
	"fmt"
	"testing"
)

func mustNewDuneClient() *DuneClient {
	return NewDuneClient("7pfUttO4tgrgxgJZUqQIDhwIXPAywY0A", "https://api.dune.com")
}

func Test_dune_QueryData(t *testing.T) {
	client := mustNewDuneClient()

	filter := Filter{
		APINumber: "3242787",
		Limit:     1000,
	}

	var result any
	if err := client.QueryData(context.Background(), filter, &result); err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}
