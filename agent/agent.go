package agent

/*
type Agent interface {
	ReviewCancel(CancelDetail)
}
*/

type AgentID int

const (
	IDFirst AgentID = iota
	IDVia
)

type CancelDetail struct {
	PassengerName string
	JourneyID     int64
}

type CancellationDetail struct {
	Passengers []Passenger
}

type Passenger struct {
	ID            int64
	IsCancellable bool
}

/*
func New(agentID int) Agent {
	switch agentID {
	case via.ID:
		return &via.AgentVia{}
	case first.ID:
		return &first.AgentFirst{}

	}
	return nil
}
*/
