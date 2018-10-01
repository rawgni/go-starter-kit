package cancelAction

func (c CancelAction) Estimate() {
	c.SelectedAgent.ReviewCancel(c.CancelDetail)
}
