package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}

func ApplyQueryDecorator[H any, R any](handler QueryHandler[H, R], logger *logrus.Entry, metricsClient MetricsClient) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		base: queryMetricsDecorator[H, R]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}
