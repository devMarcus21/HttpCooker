package orchestrator

type OrchestratorBuilder struct {
	orchestrator *Orchestrator
}

func (builder *OrchestratorBuilder) SetUrl(url string) *OrchestratorBuilder {
	builder.orchestrator.url = url
	return builder
}

func (builder *OrchestratorBuilder) SetMethod(method string) *OrchestratorBuilder {
	builder.orchestrator.method = method
	return builder
}

func (builder *OrchestratorBuilder) SetBearerString(bearer string) *OrchestratorBuilder {
	builder.orchestrator.bearerString = bearer
	return builder
}

func (builder *OrchestratorBuilder) SetSets(sets int) *OrchestratorBuilder {
	builder.orchestrator.sets = sets
	return builder
}

func (builder *OrchestratorBuilder) SetCalls(calls int) *OrchestratorBuilder {
	builder.orchestrator.calls = calls
	return builder
}

func (builder *OrchestratorBuilder) SetInvoker(invoker func(func() int, int) map[int]int) *OrchestratorBuilder {
	builder.orchestrator.invoker = invoker
	return builder
}

func (builder *OrchestratorBuilder) AddHttpClientFunctionCallback(factory func(string, string, string) func() int) *OrchestratorBuilder {
	httpClientFunction := factory(
		builder.orchestrator.url,
		builder.orchestrator.method,
		builder.orchestrator.bearerString)
	builder.orchestrator.callback = httpClientFunction

	return builder
}

func (builder *OrchestratorBuilder) Build() *Orchestrator {
	return builder.orchestrator
}

func GetOrchestratorBuilder() *OrchestratorBuilder {
	return &OrchestratorBuilder{orchestrator: buildOrchestrator()}
}
