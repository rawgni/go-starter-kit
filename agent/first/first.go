package first

import (
	"fmt"

	"github.com/srdhoni/go-starter-kit/agent"
)

type AgentFirst struct {
}

func (a AgentFirst) ReviewCancel(param agent.CancelDetail) {
	fmt.Println(" ReviewCancel for Passenger ", param, " Started by dummy agent")
}

func (a AgentFirst) GetCancellationDetail(id int64) agent.CancellationDetail {
	// hit first
	// map to agent type

	// first travel , passenger 1 is cancellable , passenger 2 is not

	return agent.CancellationDetail{
		Passengers: []agent.Passenger{
			agent.Passenger{
				ID:            1,
				IsCancellable: true,
			},
			agent.Passenger{
				ID:            2,
				IsCancellable: false,
			},
		},
	}
}
