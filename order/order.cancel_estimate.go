package order

/*
type Estimator interface {
	Estimate(int64, []int64) []CancellationEstimationResponse
}

type CancellationEntry struct {
	PassengerID int64
	JourneyID   int64
}

type CancellationEstimationResponse struct {
	PassengerID     int64
	JourneyID       int64
	EstimatedRefund float64
}

func (f Flight) GetEstimate(m map[agent.AgentID]interface{}, list []CancellationEntry) []CancellationEstimationResponse {

	estimations := make([]CancellationEstimationResponse, 0)

	for _, journey := range f.Journeys {
		est, ok := m[journey.AgentID].(Estimator)
		if !ok {
			continue
		}
		pList := make([]int64, 0)
		for _, e := range list {
			if e.JourneyID == journey.ID {
				pList = append(pList, e.PassengerID)
			}
		}
		res := est.Estimate(journey.ID, pList)
		for _, r := range res {
			estimations = append(estimations, CancellationEstimationResponse(r))
		}
	}

	return estimations
}
*/
