package cancelAction

func New(agent Agent, detail CancelDetail) *CancelAction {
	return &CancelAction{
		SelectedAgent: agent,
		CancelDetail:  detail,
	}
}
