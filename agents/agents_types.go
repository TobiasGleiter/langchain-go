package agents

type AgentAction struct {
	Tool      string
	ToolInput string
}

type AgentStep struct {
	Action      AgentAction
	Observation string
}

type AgentFinish struct {
	ReturnValues map[string]any
}