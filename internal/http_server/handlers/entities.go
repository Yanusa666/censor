package handlers

type ErrorResp struct {
	Error string `json:"error"`
}

type CheckReq struct {
	Text string `json:"text"`
}

type CheckResp struct {
	Status bool `json:"status"`
}
