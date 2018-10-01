package order_test

import (
	"reflect"
	"testing"

	"github.com/srdhoni/go-starter-kit/agent"
	"github.com/srdhoni/go-starter-kit/order"
)

// fake detail getter
type fcdg struct {
	data map[int64]agent.CancellationDetail
}

// return whatever is in data
func (f fcdg) GetCancellationDetail(i int64) agent.CancellationDetail {
	return f.data[i]
}

func TestGetCancellationDetail(t *testing.T) {

	// agent 31337
	// can cancel passenger 1 and 3 for journey 1
	// can only cancel passenger 2 for journey 2
	m := map[agent.AgentID]interface{}{
		31337: fcdg{
			data: map[int64]agent.CancellationDetail{
				1: agent.CancellationDetail{
					Passengers: []agent.Passenger{
						agent.Passenger{
							ID:            1,
							IsCancellable: true,
						},
						agent.Passenger{
							ID:            3,
							IsCancellable: true,
						},
					},
				},
				2: agent.CancellationDetail{
					Passengers: []agent.Passenger{
						agent.Passenger{
							ID:            2,
							IsCancellable: true,
						},
					},
				},
			},
		},
	}

	// setup order
	o := order.Flight{
		Journeys: []order.Journey{
			order.Journey{
				ID:      1,
				AgentID: 31337,
			},
			order.Journey{
				ID:      2,
				AgentID: 31337,
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
			order.Passenger{
				ID: 4,
			},
		},
	}

	// p1 p3 cancellable for j1
	// p2 cancellable for j2
	expected := order.CancelDetail{
		Cancellables: []order.Cancellable{
			order.Cancellable{
				PassengerID: 1,
				JourneyID:   1,
			},
			order.Cancellable{
				PassengerID: 3,
				JourneyID:   1,
			},
			order.Cancellable{
				PassengerID: 2,
				JourneyID:   2,
			},
		},
	}

	actual := o.GetCancellationDetail(m)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("fail :'(")
	}
}
