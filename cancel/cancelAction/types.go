package cancelAction

type Agent interface {
	ReviewCancel(CancelDetail)
}

type CancelAction struct {
	SelectedAgent Agent
	CancelDetail  CancelDetail
}

type FlightCancel struct {
	CancelDetails []CancelDetail
}

type CancelDetail struct {
	JourneyID     int64
	PassengerName string
}
