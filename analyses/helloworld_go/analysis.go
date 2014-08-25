package main

import (
	"log"
	"github.com/svenkreiss/databench_go/databench"
)


func createAnalysis() databench.AnalysisI {
    analysis := new(databench.Analysis)

    analysis.AddListener(&databench.Listener{"connect", func(message interface{}) {
        log.Printf("Listener for connect: %v\n", message)
    }})

    return analysis
}

func main() {
	log.Printf("Starting HelloWorld Go analysis ...\n")

    meta := databench.NewMeta("dummypi_go", "Bla bla", createAnalysis)
    meta.EventLoop()
}
