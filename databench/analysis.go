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
	setEmitFn(func(id int, signal string, message interface{}))
	setID(id int)
}

type Analysis struct {
	id int
	listeners []*Listener
	emitFn func(analysisID int, signal string, message interface{})
}

func (analysis *Analysis) AddListener(l *Listener) {
	log.Printf("AddListener()\n")
	analysis.listeners = append(analysis.listeners, l)
	log.Printf("listeners: %v\n", analysis.listeners)
}

func (analysis *Analysis) Emit(signal string, message interface{}) {
	log.Printf("emit called %s -- %v\n", signal, message)
	analysis.emitFn(analysis.id, signal, message)
	log.Printf("---------------------\n")
}

func (analysis *Analysis) setEmitFn(f func(int, string, interface{})) {
	analysis.emitFn = f
}

func (analysis *Analysis) setID(id int) {
	analysis.id = id
}

func (analysis *Analysis) callListener(signal string, message interface{}) {
	for _, l := range analysis.listeners {
		if l.Signal == signal {
			l.Callback(message)
		}
	}
}
