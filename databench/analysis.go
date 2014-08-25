package databench

import (
	"log"
)

type Listener struct {
	Signal string
	Callback func(message interface{})
}

type AnalysisI interface {
	AddListener(*Listener)
	callListener(signal string, message interface{})

	Emit(string, interface{})
	setEmitFn(func(signal string, message interface{}))
}

type Analysis struct {
	listeners []*Listener
	emitFn func(signal string, message interface{})
}

func (analysis *Analysis) AddListener(l *Listener) {
	log.Printf("AddListener()\n")
	analysis.listeners = append(analysis.listeners, l)
	log.Printf("listeners: %v\n", analysis.listeners)
}

func (analysis *Analysis) Emit(signal string, message interface{}) {
	analysis.emitFn(signal, message)
}

func (analysis *Analysis) setEmitFn(f func(signal string, message interface{})) {
	analysis.emitFn = f
}

func (analysis *Analysis) callListener(signal string, message interface{}) {
	for _, l := range analysis.listeners {
		if l.Signal == signal {
			l.Callback(message)
		}
	}
}
