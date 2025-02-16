package requests

type LogPOST struct {
	UserID string `json:"user_id"`
	Action string `json:"type"`
	Info   string `json:"info"`
}
