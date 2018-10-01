package cancelAction

import "github.com/srdhoni/go-starter-kit/agent"

type Agent interface {
	ReviewCancel(agent.CancelDetail)
}

type CancelAction struct {
	SelectedAgent Agent
	CancelDetail  agent.CancelDetail
}

type FlightCancel struct {
	CancelDetails []CancelDetail
}

type CancelDetail struct {
	JourneyID     int64
	PassengerName string
}
