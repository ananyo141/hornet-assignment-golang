package utils

type ResponseFormat struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Page    int         `json:"page"`
	Data    interface{} `json:"data"`
}

func HttpResponse(success bool, message string, data any) ResponseFormat {
	return ResponseFormat{
		Success: success,
		Message: message,
		Data:    data,
	}
}
