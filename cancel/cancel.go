package cancel

import "github.com/srdhoni/go-starter-kit/agent"

/*
type Cancel interface {
	Estimate()
}

func New(flight order.Flight, detail action.CancelDetail) Cancel {
	var agentID int
	for i := range flight.Journeys {
		if flight.Journeys[i].ID == detail.JourneyID {
			agentID = flight.Journeys[i].AgentID
			break
		}
	}

	selectedAgent := agent.New(agentID)
	return action.New(selectedAgent, detail)
}
*/

type CancelDetailGetter interface {
	GetCancellationDetail(int64) agent.CancellationDetail
}

type Cancel struct {
	Journeys []Journey
}

type Journey struct {
	AgentID agent.AgentID
	ID      int64
}

type CancelDetail struct {
	Cancellables []Cancellable
}

type Cancellable struct {
	PassengerID int64
	JourneyID   int64
}

// Get order cancellation detail
// func (c Cancel) GetCancellationDetail(m map[agent.AgentID]interface{}) CancelDetail {
func (c Cancel) GetCancellationDetail() CancelDetail {

	var cancellables []Cancellable

	for _, journey := range c.Journeys {
		cdg, ok := agent.New(journey.AgentID).(CancelDetailGetter)
		//	cdg, ok := m[journey.AgentID].(CancelDetailGetter)
		// no cancel detail getter implementation
		if !ok {
			continue
		}
		res := processGet(cdg, journey.ID)
		cancellables = append(cancellables, res...)
	}

	return CancelDetail{
		Cancellables: cancellables,
	}
}

func processGet(dg CancelDetailGetter, id int64) []Cancellable {
	var c []Cancellable

	res := dg.GetCancellationDetail(id)

	for _, p := range res.Passengers {
		if p.IsCancellable {
			c = append(c, Cancellable{
				PassengerID: p.ID,
				JourneyID:   id,
			})
		}
	}
	return c
}
