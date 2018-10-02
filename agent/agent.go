package agent

type Agent interface {
}

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

type CancellationEstimation struct {
	PassengerID     int64
	JourneyID       int64
	EstimatedRefund float64
}

func New(agentID AgentID) Agent {
	switch agentID {
	case IDFirst:
		return &via{}
	case IDVia:
		return &first{}

	}
	return nil
}
