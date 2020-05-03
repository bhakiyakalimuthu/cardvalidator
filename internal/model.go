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

type ServerVerifyResp struct {
	Number struct {
		Length int  `json:"length"`
		Luhn   bool `json:"luhn"`
	} `json:"number"`
	Scheme  string `json:"scheme"`
	Type    string `json:"type"`
	Brand   string `json:"brand"`
	Prepaid bool   `json:"prepaid"`
	Country struct {
		Numeric   string `json:"numeric"`
		Alpha2    string `json:"alpha2"`
		Name      string `json:"name"`
		Emoji     string `json:"emoji"`
		Currency  string `json:"currency"`
		Latitude  int    `json:"latitude"`
		Longitude int    `json:"longitude"`
	} `json:"country"`
	Bank struct {
		Name  string `json:"name"`
		URL   string `json:"url"`
		Phone string `json:"phone"`
		City  string `json:"city"`
	} `json:"bank"`
}