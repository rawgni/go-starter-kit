package cancel

import (
	"github.com/srdhoni/go-starter-kit/agent"
	action "github.com/srdhoni/go-starter-kit/cancel/cancelAction"
	"github.com/srdhoni/go-starter-kit/order"
)

type Agent interface {
	ReviewCancel(action.CancelDetail)
}
type Cancel interface {
	Estimate()
}

func New(flight order.Flight, detail action.CancelDetail) Cancel {
	var agentID int
	for i := range flight.Journeys {
		if flight.Journeys[i].ID == detail.JourneyID {
			agentID = flight.Journeys[i].AgentID
			break
		}
	}

	selectedAgent := agent.New(agentID)
	return action.New(selectedAgent, detail)
}

// func Estimate(flight order.Flight, cancelDetail cancelAction.CancelDetail) Cancel {
// 	var agentID int
// 	for i := range flight.Journeys {
// 		if flight.Journeys[i].ID == cancelDetail.JourneyID {
// 			agentID = flight.Journeys[i].AgentID
// 			break
// 		}
// 	}

// 	selectedAgent := agent.New(agentID)
// 	return cancelAction.New(selectedAgent)
// }
