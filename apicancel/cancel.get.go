package apicancel

import (
	"log"

	"github.com/srdhoni/go-starter-kit/agent"
	"github.com/srdhoni/go-starter-kit/order"
)

func HandlerGet(m map[agent.AgentID]interface{}) {

	flight := order.Flight{
		Journeys: []order.Journey{
			order.Journey{
				ID:      1111,
				AgentID: agent.IDFirst,
			},
			order.Journey{
				ID:      2222,
				AgentID: agent.IDVia,
			},
		},
		Passengers: []order.Passenger{
			order.Passenger{
				ID: 1,
			},
			order.Passenger{
				ID: 2,
			},
		},
	}

	log.Printf("%#v\n", flight.GetCancellationDetail(m))
}
