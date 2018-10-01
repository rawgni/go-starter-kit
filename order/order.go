package order

type (
	Flight struct {
		ID       int64
		Journeys []Journey
	}

	Journey struct {
		ID      int64
		AgentID int
	}
)
