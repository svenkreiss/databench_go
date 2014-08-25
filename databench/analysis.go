package databench

import (
	"log"
)

type Listener struct {
	Signal string
	Callback func(message string)
}

type AnalysisI interface {
	AddListener(*Listener)
	Emit(string, string)
}

type Analysis struct {
	listeners []*Listener
	emit_fn func(signal string, message string)
}

func (analysis *Analysis) AddListener(l *Listener) {
	log.Printf("AddListener()\n")
	analysis.listeners = append(analysis.listeners, l)
	log.Printf("listeners: %v\n", analysis.listeners)
}

func (analysis *Analysis) Emit(signal string, message string) {
	analysis.emit_fn(signal, message)
}
