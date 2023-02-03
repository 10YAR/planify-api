package utils

import (
	"api/types"
)

func E503(message string, err error) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 503, Error: err}
}

func E400(message string, err error) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 400, Error: err}
}

func E401(message string, err error) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 401, Error: err}
}

func E403(message string, err error) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 403, Error: err}
}

func E404(message string, err error) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 404, Error: err}
}
