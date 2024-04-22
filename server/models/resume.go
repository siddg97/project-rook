package models

type CreateResumeRequest struct {
	UserID string `json:"userId"`
}

type UpdateResumeRequest struct {
	UserID     string `json:"userId"`
	Experience string `json:"experience"`
}

type GetResumeRequest struct {
	UserID string `json:"userId"`
}

type GetResumeResponse struct {
	Resume        ResumeDocument          `json:"resume"`
	PromptHistory []PromptHistoryDocument `json:"promptHistory"`
}
