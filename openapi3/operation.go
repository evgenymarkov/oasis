package openapi3

type Operation struct {
	OperationID string `json:"operationId"`
	Summary     string `json:"summary"`
}

func NewOperation() *Operation {
	return &Operation{
		OperationID: "",
		Summary:     "",
	}
}

func (o *Operation) SetOperationID(id string) *Operation {
	o.OperationID = id

	return o
}

func (o *Operation) SetSummary(summary string) *Operation {
	o.Summary = summary

	return o
}
