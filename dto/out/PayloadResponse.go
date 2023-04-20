package out

type PayloadResponseSuccess struct {
	Success   bool        `json:"success"`
	RequestID string      `json:"request_id"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type PayloadResponseFailed struct {
	Success      bool   `json:"success"`
	RequestID    string `json:"request_id"`
	ErrorCode    int64  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
