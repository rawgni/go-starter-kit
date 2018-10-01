package first

import (
	"fmt"

	"github.com/srdhoni/go-starter-kit/cancel/cancelAction"
)

type AgentFirst struct {
}

func (a AgentFirst) ReviewCancel(param cancelAction.CancelDetail) {
	fmt.Println(" ReviewCancel for Passenger ", param, " Started by dummy agent")
}
