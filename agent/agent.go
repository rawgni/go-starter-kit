package agent

import (
	"github.com/srdhoni/go-starter-kit/agent/first"
	"github.com/srdhoni/go-starter-kit/agent/via"
	"github.com/srdhoni/go-starter-kit/cancel/cancelAction"
)

type Agent interface {
	ReviewCancel(cancelAction.CancelDetail)
}

func New(agentID int) Agent {
	switch agentID {
	case via.ID:
		return &via.AgentVia{}
	case first.ID:
		return &first.AgentFirst{}

	}
	return nil
}
