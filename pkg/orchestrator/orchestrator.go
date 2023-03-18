package orchestrator

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Orchestrator struct {
	url          string
	method       string
	bearerString string
	sets         int
	calls        int
	invoker      func(func() int, int) map[int]int
	callback     func() int
}

func (orchestrator Orchestrator) GetUrl() string {
	return orchestrator.url
}

func (orchestrator Orchestrator) GetMethod() string {
	return orchestrator.method
}

func (orchestrator Orchestrator) GetBearerString() string {
	return orchestrator.bearerString
}

func (orchestrator Orchestrator) GetSets() int {
	return orchestrator.sets
}

func (orchestrator Orchestrator) GetCalls() int {
	return orchestrator.calls
}

func (orchestrator *Orchestrator) Run() {
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t\n", "1xx", "2xx", "3xx", "4xx", "5xx", "exception")

	for set := 0; set < orchestrator.sets; set++ {
		results := orchestrator.invoker(orchestrator.callback, orchestrator.calls)
		printResults(w, results)
	}
	w.Flush()
}

func buildOrchestrator() *Orchestrator {
	return &Orchestrator{url: "", method: "GET", bearerString: "", sets: 1, calls: 0}
}
