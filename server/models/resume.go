package models

import "mime/multipart"

type CreateResumeRequest struct {
	Resume *multipart.FileHeader
}

type UpdateResumeRequest struct {
	Experience string `json:"experience"`
}

type GetResumeRequest struct {
	UserID string `json:"userId"`
}

type GetResumeResponse struct {
	Resume        ResumeDocument          `json:"resume"`
	PromptHistory []PromptHistoryDocument `json:"promptHistory"`
}
