package main

import (
	"log"
	"github.com/svenkreiss/databench_go/databench"
)

type statusMessage struct {
	Message string
}

func createAnalysis() databench.AnalysisI {
	analysis := new(databench.Analysis)

	analysis.AddListener(&databench.Listener{"connect", func(message interface{}) {
		analysis.Emit("status", statusMessage{"HelloWorld"})
	}})

	return analysis
}

func main() {
	log.Printf("Starting HelloWorld Go analysis ...\n")

	meta := databench.NewMeta("helloworld_go", "Bla bla", createAnalysis)
	meta.EventLoop()
}
