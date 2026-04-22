package process

import (
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type SignalListener struct {
	logger *zap.Logger
}

func NewSignalListener(logger *zap.Logger) *SignalListener {
	return &SignalListener{
		logger: logger,
	}
}

func (s SignalListener) Start(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()

	s.logger.Info("signal listener is shutting down...")

	return nil
}
