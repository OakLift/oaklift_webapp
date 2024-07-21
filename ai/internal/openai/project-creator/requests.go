package project

type CreateRequest struct {
	Language Language `json:"language"`
	Level ProjectLevel `json:"level"`
}
