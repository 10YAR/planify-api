package utils

import (
	"api/types"
)

func E503(message string) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 503}
}

func E400(message string) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 400}
}

func E401(message string) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 401}
}

func E403(message string) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 403}
}

func E404(message string) types.HttpResponse {
	return types.HttpResponse{Status: 0, Message: message, HttpCode: 404}
}
