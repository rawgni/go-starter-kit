package order

import "github.com/srdhoni/go-starter-kit/agent"

type (
	Flight struct {
		ID         int64
		Journeys   []Journey
		Passengers []Passenger
	}

	Journey struct {
		ID      int64
		AgentID agent.AgentID
	}

	Passenger struct {
		ID int64
	}
)
