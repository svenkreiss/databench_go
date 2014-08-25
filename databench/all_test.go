package databench

import (
    "log"
    "testing"
)

func createAnalysis() AnalysisI {
    analysis := new(Analysis)

    analysis.AddListener(&Listener{"test", func(message string) {
        log.Printf("Listener for test: %s\n", message)
    }})

    return analysis
}

func TestMeta(t *testing.T) {
    log.Printf("Start test\n")

    meta := NewMeta("dummypi_go", "Bla bla", createAnalysis)
    meta.EventLoop()
}
