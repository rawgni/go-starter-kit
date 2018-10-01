package via

import (
	"fmt"

	"github.com/srdhoni/go-starter-kit/cancel/cancelAction"
)

type AgentVia struct {
	ID int
}

func (a AgentVia) ReviewCancel(param cancelAction.CancelDetail) {
	fmt.Println(" ReviewCancel for Passenger ", param, " Started by agentID: ", ID)
}
