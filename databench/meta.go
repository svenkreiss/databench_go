package databench

import (
	"log"
	"time"
	"strconv"
	"encoding/json"
	zmq "github.com/pebbe/zmq4"
)


type APIPublishOnPort struct {
	Namespace string `json:"__databench_namespace"`
	Port int `json:"publish_on_port"`
}

type APIMessage struct {
	Signal string `json:"signal"`
	Message interface{} `json:"message"`
}

type APISignal struct {
	Namespace string `json:"__databench_namespace"`
	AnalysisID int `json:"__analysis_id"`
	Message APIMessage `json:"message"`
}


type Meta struct {
	name string
	description string
	analysisCreator func() AnalysisI

	analyses map[int]AnalysisI
	zmqPublisher *zmq.Socket
}

func NewMeta(name string, description string, analysisCreator func() AnalysisI) *Meta {
	return &Meta{
		name: name,
		description: description,
		analysisCreator: analysisCreator,
		analyses: map[int]AnalysisI{},
	}
}

func (meta *Meta) instantiateAnalysis(id int) AnalysisI {
	log.Printf("instantiateAnalysis(%d)\n", id)
	i := meta.analysisCreator()
	i.setID(id)
	meta.analyses[id] = i
	return i
}

func (meta *Meta) emitZmq(analysisID int, signal string, message interface{}) {
	if meta.zmqPublisher == nil {
		log.Printf("Cannot send %s -- %v because not ready to publish to zmq.\n", signal, message)
		return
	}

	log.Printf("Sending via zmq: %d, %s -- %v\n", analysisID, signal, message)
	m := APISignal{meta.name, analysisID, APIMessage{signal, message}}
	msg, _ := json.Marshal(m)
	log.Printf("Json encoded: %s\n", msg)
	meta.zmqPublisher.SendBytes(msg, 0)
}

func (meta *Meta) EventLoop() {
	log.Printf("EventLoop() for %s.\n", meta.name)

	zmqSubscriber, _ := zmq.NewSocket(zmq.SUB)
	defer zmqSubscriber.Close()
	zmqSubscriber.Connect("tcp://127.0.0.1:8041")
	zmqSubscriber.SetSubscribe("")

	for {
		msg, err := zmqSubscriber.RecvBytes(0)
		if err != nil { break }

		log.Printf("Received: %s\n", msg)

		// try whether this is a signal message
		signal := new(APISignal)
		errU2 := json.Unmarshal(msg, signal)
		if errU2 == nil && signal.Message.Signal != "" {
			if signal.Namespace != meta.name {
				continue
			}

			// check whether we already have an analysis with this id
			if _, ok := meta.analyses[signal.AnalysisID]; !ok {
				i := meta.instantiateAnalysis(signal.AnalysisID)
				i.setEmitFn(meta.emitZmq)
			}
			log.Printf("signal: %v\n", signal)
			meta.analyses[signal.AnalysisID].callListener(signal.Message.Signal, signal.Message.Message)

			continue
		}

		if meta.zmqPublisher == nil {
			// try whether it is a PublishOnPort message
			pop := new(APIPublishOnPort)
			errU := json.Unmarshal(msg, pop)
			if errU == nil && pop.Port != 0 {
				if pop.Namespace != meta.name {
					continue
				}
				log.Printf("pop: %v\n", pop)

				log.Printf("Go kernel: Initialize zmq publisher\n")
				meta.zmqPublisher, _ = zmq.NewSocket(zmq.PUB)
				meta.zmqPublisher.Bind("tcp://127.0.0.1:"+strconv.Itoa(pop.Port))

				// wait for slow tcp bind
				time.Sleep(500 * time.Millisecond)

				// send hello
				meta.zmqPublisher.Send("{\"__databench_namespace\":\""+meta.name+"\"}", 0)
				continue
			}
		}
	}

	log.Printf("go kernel: end\n")
}
