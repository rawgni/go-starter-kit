package agent

type via struct{}

func (a via) GetCancellationDetail(id int64) CancellationDetail {
	// hit via
	// map reason type to agent type

	// agent via , p1 is not cancellable , p2 is

	return CancellationDetail{
		Passengers: []Passenger{
			Passenger{
				ID:            1,
				IsCancellable: false,
			},
			Passenger{
				ID:            2,
				IsCancellable: true,
			},
		},
	}
}
