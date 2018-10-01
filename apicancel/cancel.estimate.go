package apicancel

import (
	"github.com/srdhoni/go-starter-kit/cancel"
	"github.com/srdhoni/go-starter-kit/cancel/cancelAction"
	"github.com/srdhoni/go-starter-kit/order"
)

func HandlerEstimate() {

	flight := order.Flight{
		Journeys: []order.Journey{{
			ID:      100,
			AgentID: 1,
		}},
	}
	cancelReq := cancelAction.CancelDetail{
		PassengerName: "Jono",
		JourneyID:     100,
	}

	newCancel := cancel.New(flight, cancelReq)
	newCancel.Estimate()
}
