package models

import "time"

type ResumeDocument struct {
	UserID     string `firebase:"userId" json:"userId"`
	ResumeID   string `firebase:"resumeId" json:"resumeId"`
	ResumeText string `firebase:"resumeText" json:"resumeText"`
}

type PromptHistoryDocument struct {
	Id        string    `firebase:"id" json:"id"`
	CreatedAt time.Time `firebase:"createdAt" json:"createdAt"`
	Role      string    `firebase:"role" json:"role"`
	Text      string    `firebase:"text" json:"text"`
}
