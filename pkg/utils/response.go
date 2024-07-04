package utils

import "chatapp/model"

func Response(status int, message string, data interface{}) model.Response {
	return model.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
