package orchestrator

import (
	httpResultslogging "github.com/devMarcus21/HttpCooker/pkg/logging/httpResultsLogging"
)

type Orchestrator struct {
	url           string
	method        string
	bearerString  string
	groups        int
	sets          int
	calls         int
	invoker       func(func() int, int) map[int]int
	callback      func() int
	resultsLogger httpResultslogging.HttpResultsLogger
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

func (orchestrator Orchestrator) GetGroups() int {
	return orchestrator.groups
}

func (orchestrator Orchestrator) GetSets() int {
	return orchestrator.sets
}

func (orchestrator Orchestrator) GetCalls() int {
	return orchestrator.calls
}

func (orchestrator *Orchestrator) Run() {
	orchestrator.resultsLogger.LogHeaderMessage()

	for group := 0; group < orchestrator.groups; group++ {
		for set := 0; set < orchestrator.sets; set++ {
			results := orchestrator.invoker(orchestrator.callback, orchestrator.calls)
			orchestrator.resultsLogger.LogHttpResults(results)
		}
		orchestrator.resultsLogger.LogPostGroup()
	}
	orchestrator.resultsLogger.LoggerCleanUp()
}

func buildOrchestrator() *Orchestrator {
	return &Orchestrator{url: "", method: "GET", bearerString: "", groups: 1, sets: 1, calls: 0}
}
