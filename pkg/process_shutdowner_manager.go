package process

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

type ProcessShutdownerManager struct {
	logger             *zap.Logger
	shutdownerRegistry *ShutdownerRegistry
}

func NewProcessShutdownerManager(
	logger *zap.Logger,
	shutdownerRegistry *ShutdownerRegistry,
) *ProcessShutdownerManager {
	return &ProcessShutdownerManager{
		logger:             logger,
		shutdownerRegistry: shutdownerRegistry,
	}
}

func (r *ProcessShutdownerManager) Run(
	ctx context.Context,
	processes ...Process,
) error {
	processes = append(
		processes,
		NewProcess("signal", NewSignalListener(r.logger)),
	)

	return NewProcessManager(
		r.logger,
		processes...,
	).WithGlobalShutdowner(
		func() error {
			var errs error

			for i := len(r.shutdownerRegistry.shutdowners) - 1; i >= 0; i-- {
				r.logger.Info("Shutting down global process start", zap.String("process", r.shutdownerRegistry.shutdowners[i].name))

				if err := r.shutdownerRegistry.shutdowners[i].shutdowner(); err != nil {
					r.logger.Error(
						"Failed to shutdown global process shutdowner",
						zap.String("shutdowner", r.shutdownerRegistry.shutdowners[i].name),
						zap.Error(err),
					)
					errs = errors.Join(errs, err)
				}

				r.logger.Info("Shutting down success", zap.String("process", r.shutdownerRegistry.shutdowners[i].name))
			}

			return errs
		},
	).Start(ctx)
}
