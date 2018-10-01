package order

import (
	"github.com/srdhoni/go-starter-kit/agent"
)

type CancelDetailGetter interface {
	GetCancellationDetail(int64) agent.CancellationDetail
}

type CancelDetail struct {
	Cancellables []Cancellable
}

type Cancellable struct {
	PassengerID int64
	JourneyID   int64
}

// Get order cancellation detail
func (f Flight) GetCancellationDetail(m map[agent.AgentID]interface{}) CancelDetail {

	var cancellables []Cancellable

	for _, journey := range f.Journeys {
		cdg, ok := m[journey.AgentID].(CancelDetailGetter)
		// no cancel detail getter implementation
		if !ok {
			continue
		}
		res := journey.processGet(cdg)
		cancellables = append(cancellables, res...)
	}

	return CancelDetail{
		Cancellables: cancellables,
	}
}

func (j Journey) processGet(dg CancelDetailGetter) []Cancellable {
	var c []Cancellable

	res := dg.GetCancellationDetail(j.ID)

	for _, p := range res.Passengers {
		if p.IsCancellable {
			c = append(c, Cancellable{
				PassengerID: p.ID,
				JourneyID:   j.ID,
			})
		}
	}
	return c
}
