package process

import "sync"

type NamedShutdowner struct {
	name       string
	shutdowner ShutdownerFunc
}

type ShutdownerRegistry struct {
	shutdowners []NamedShutdowner
	mx          sync.Mutex
}

func NewShutdownerRegistry() *ShutdownerRegistry {
	return &ShutdownerRegistry{}
}

func (r *ShutdownerRegistry) Add(name string, f ShutdownerFunc) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.shutdowners = append(r.shutdowners, NamedShutdowner{
		name:       name,
		shutdowner: f,
	})
}
