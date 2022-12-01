package reflows

type (
	ReflowRequest interface {
		GetWorkflows() []string
	}

	ReflowWorkflows struct {
		Workflows []string `json:"workflows,omitempty"`
	}

	ReflowByEventsRequest struct {
		Events []string `json:"events,omitempty"`
		ReflowWorkflows
	}

	ReflowBySubjectsRequest struct {
		Subjects []string `json:"subjects,omitempty"`
		ReflowWorkflows
	}
)

func (r *ReflowByEventsRequest) GetWorkflows() []string {
	return r.Workflows
}
func (r *ReflowBySubjectsRequest) GetWorkflows() []string {
	return r.Workflows
}
