package apicancel

import (
	"log"

	"github.com/srdhoni/go-starter-kit/agent"
	"github.com/srdhoni/go-starter-kit/cancel"
)

// func HandlerGet(m map[agent.AgentID]interface{}) {
func HandlerGet() {
	c := cancel.Cancel{
		Journeys: []cancel.Journey{
			cancel.Journey{
				ID:      1111,
				AgentID: agent.IDFirst,
			},
			cancel.Journey{
				ID:      2222,
				AgentID: agent.IDVia,
			},
		},
	}
	// log.Printf("%#v\n", c.GetCancellationDetail(m))
	log.Printf("%#v\n", c.GetCancellationDetail())
}
