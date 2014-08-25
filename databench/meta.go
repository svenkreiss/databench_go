package databench

import (
	"log"
	zmq "github.com/pebbe/zmq4"
)

type MetaI interface {
	EventLoop()
	emitZmq()
}

type Meta struct {
	name string
	description string
	analysisCreator func() AnalysisI

	analyses []AnalysisI
	zmq_publisher *zmq.Socket
}

func NewMeta(name string, description string, analysisCreator func() AnalysisI) *Meta {
	return &Meta{name: name, description: description, analysisCreator: analysisCreator}
}

func (meta *Meta) instantiateAnalysis(id int) {
	log.Printf("instantiateAnalysis(%d)\n", id)
	meta.analyses = append(meta.analyses, meta.analysisCreator())
}

func (meta *Meta) emitZmq(signal string, message string) {
	if meta.zmq_publisher == nil {
		log.Printf("Cannot send %s -- %s because not ready to publish to zmq.\n", signal, message)
		return
	}

	log.Printf("Sending via zmq: %s -- %s\n", signal, message)
	meta.zmq_publisher.SendMessage(signal, message)
}

func (meta *Meta) EventLoop() {
	log.Printf("EventLoop()\n")

	zmq_subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer zmq_subscriber.Close()
	zmq_subscriber.Connect("tcp://127.0.0.1:8041")
	zmq_subscriber.SetSubscribe("")

	meta.instantiateAnalysis(1)

	for {
		msg, err := zmq_subscriber.Recv(0)
		if err != nil { break }

		log.Printf("Received: %s\n", msg)
	}

	log.Printf("go kernel: end\n")
}
