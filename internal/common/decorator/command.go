package decorator

import (
	"context"
	"github.com/sirupsen/logrus"
)

type CommandHandler[C any, R any] interface {
	Handle(ctx context.Context, command C) (R, error)
}

func ApplyCommandDecorators[C any, R any](handler CommandHandler[C, R], logger *logrus.Entry, client MetricsClient) CommandHandler[C, R] {
	return queryLoggingDecorator[C, R]{
		logger: logger,
		base: queryMetricsDecorator[C, R]{
			base:   handler,
			client: client,
		},
	}
}
