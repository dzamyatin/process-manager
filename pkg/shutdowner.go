package process

import (
	"context"
	"go.uber.org/zap"
)

type ShutdownerFunc func() error

type Shutdowner struct {
	logger     *zap.Logger
	shutdowner ShutdownerFunc
}

func NewShutdowner(
	logger *zap.Logger,
	shutdowner ShutdownerFunc,
) *Shutdowner {
	return &Shutdowner{
		logger:     logger,
		shutdowner: shutdowner,
	}
}

func (r Shutdowner) Start(ctx context.Context) error {
	<-ctx.Done()

	return nil
}

func (r Shutdowner) Shutdown() error {
	return r.shutdowner()
}
