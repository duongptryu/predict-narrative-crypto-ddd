package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func ApplyCommandDecorator[H any](handler CommandHandler[H], logger *logrus.Entry, mtricsClient MetricsClient) CommandHandler[H] {
	return commandLoggingDecorator[H]{
		base: commandMetricsDecorator[H]{
			base:   handler,
			client: mtricsClient,
		},
		logger: logger,
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
