package cancelAction

import "github.com/srdhoni/go-starter-kit/agent"

func New(agent Agent, detail agent.CancelDetail) *CancelAction {
	return &CancelAction{
		SelectedAgent: agent,
		CancelDetail:  detail,
	}
}
