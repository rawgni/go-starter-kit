package via

import (
	"fmt"

	"github.com/srdhoni/go-starter-kit/agent"
)

type AgentVia struct {
	ID int
}

func (a AgentVia) ReviewCancel(param agent.CancelDetail) {
	fmt.Println(" ReviewCancel for Passenger ", param, " Started by agentID: ", ID)
}

func (a AgentVia) GetCancellationDetail(id int64) agent.CancellationDetail {
	// hit via
	// map reason type to agent type

	// agent via , p1 is not cancellable , p2 is

	return agent.CancellationDetail{
		Passengers: []agent.Passenger{
			agent.Passenger{
				ID:            1,
				IsCancellable: false,
			},
			agent.Passenger{
				ID:            2,
				IsCancellable: true,
			},
		},
	}
}
