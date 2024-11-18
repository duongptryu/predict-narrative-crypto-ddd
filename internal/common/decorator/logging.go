package decorator

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger *logrus.Entry
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	handleType := generateActionName(cmd)

	logger := d.logger.WithFields(logrus.Fields{
		"command":      handleType,
		"command_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("Executing command")
	defer func() {
		if err == nil {
			logger.Info("Command executed successfully")
		} else {
			logger.WithError(err).Error("Failede to execute command")
		}
	}()

	return d.base.Handle(ctx, cmd)
}
