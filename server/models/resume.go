package models

import "mime/multipart"

type CreateResumeRequest struct {
	UserId string
	Resume *multipart.FileHeader
}

type CreateResumeResponse struct {
	// TODO: Change to use ResumeDocument
	Resume        string             `json:"resume"`
	ResumeDetails ModelResumeDetails `json:"resumeDetails"`
}

type UpdateResumeRequest struct {
	Experience string `json:"experience"`
	UserId     string
}

type UpdateResumeResponse struct {
	// TODO: Change to use ResumeDocument
	Resume        string             `json:"resume"`
	ResumeDetails ModelResumeDetails `json:"resumeDetails"`
}

type GetResumeRequest struct {
	UserID string `json:"userId"`
}

type GetResumeResponse struct {
	Resume        ResumeDocument          `json:"resume"`
	PromptHistory []PromptHistoryDocument `json:"promptHistory"`
}
