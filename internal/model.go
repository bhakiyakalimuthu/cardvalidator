package internal

type VerifyResponse struct {
	Success bool `json:"success"`
	Payload VerifyResponsePayload `json:"payload"`
}
type  VerifyResponsePayload struct {
	Schema string `json:"schema"`
	Type string `json:"type"`
	Bank string `json:"bank"`
}

type StatsResponse struct {
	Success bool `json:"success"`
	Start string `json:"start"`
	Limit string `json:"limit"`
	Size string `json:"size"`
	Payload map[string]string `json:"payload"`
}
