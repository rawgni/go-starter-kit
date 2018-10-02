package apicancel

/*
func HandlerEstimate() {

		flight := order.Flight{
			Journeys: []order.Journey{{
				ID:      100,
				AgentID: 1,
			}},
		}
		cancelReq := agent.CancelDetail{
			PassengerName: "Jono",
			JourneyID:     100,
		}

		newCancel := cancel.New(flight, cancelReq)
		newCancel.Estimate()
}

func HandlerEstimate(m map[agent.AgentID]interface{}) {
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
			order.Passenger{
				ID: 3,
			},
		},
	}

	list := []order.CancellationEntry{
		order.CancellationEntry{
			PassengerID: 1,
			JourneyID:   1111,
		},
		order.CancellationEntry{
			PassengerID: 2,
			JourneyID:   1111,
		},
		order.CancellationEntry{
			PassengerID: 3,
			JourneyID:   2222,
		},
	}

	log.Printf("%#v\n", flight.GetEstimate(m, list))
}
*/
