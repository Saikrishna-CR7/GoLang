package dto

type JsonResponse struct {
	Type    string    `json:"type"`
	Data    []Product `json:"data"`
	Message string    `json:"message"`
}