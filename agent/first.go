package agent

type first struct {
}

func (a first) GetCancellationDetail(id int64) CancellationDetail {
	// hit first
	// map to agent type

	// first travel , passenger 1 is cancellable , passenger 2 is not

	return CancellationDetail{
		Passengers: []Passenger{
			Passenger{
				ID:            1,
				IsCancellable: true,
			},
			Passenger{
				ID:            2,
				IsCancellable: false,
			},
		},
	}
}
