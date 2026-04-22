package process

import "context"

type ProcessStarter interface {
	Start(ctx context.Context) error
}

type ProcessShutdowner interface {
	Shutdown() error
}

type StartFunc func(ctx context.Context) error
type ShutdownFunc func() error

type Processor struct {
	starter  StartFunc
	shutdown ShutdownFunc
}

func NewProcessor(starter StartFunc, shutdown ShutdownFunc) *Processor {
	return &Processor{starter: starter, shutdown: shutdown}
}

func (p *Processor) Start(ctx context.Context) error {
	return p.starter(ctx)
}

func (s *Processor) Shutdown() error {
	return s.shutdown()
}
