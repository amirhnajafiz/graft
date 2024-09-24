package v1

type logResponse struct {
	Term    int    `json:"term"`
	Index   int    `json:"index"`
	Command string `json:"command"`
}
