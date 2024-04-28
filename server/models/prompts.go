package models

import (
	"fmt"
	"strings"
)

type ResumeDetails struct {
	UserID           string            `json:"userId"`
	Skills           []Skill           `json:"skills"`
	WorkSummaries    []WorkSummary     `json:"workSummaries"`
	PersonalProjects []PersonalProject `json:"personalProjects"`
	Educations       []Education       `json:"educations"`
}

type Skill struct {
	Name              string  `json:"name"`
	YearsOfExperience float64 `json:"yearsOfExperience"`
}

type WorkSummary struct {
	Title     string   `json:"title"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Company   string   `json:"company"`
	Work      []string `json:"work"`
}

type PersonalProject struct {
	Name                string   `json:"name"`
	Role                string   `json:"role"`
	StartDate           string   `json:"startDate"`
	EndDate             string   `json:"endDate"`
	Contributions       []string `json:"contributions"`
	NotableAchievements []string `json:"notableAchievements,omitempty"`
}

type Education struct {
	Institution         string   `json:"institution"`
	Certification       string   `json:"certification"`
	CertifiedDate       string   `json:"certifiedDate"`
	NotableAchievements []string `json:"notableAchievements,omitempty"`
}

func GetInitialResumeCreationPrompt(userId, extractedResumeText string) string {
	return fmt.Sprintf(`
		You are a resume maintainance and enhancement bot that helps user track their work summaries across time and update it to the best of your ability such that it increases the user's chance of receiving a call back from companies they applied for.

		This is the initial resume text extracted from a resume file uploaded by the user denoted by their id %s.
		The resume extracted text is bounded within ~~~.
		~~~
		%s
		~~~
		
		Please keep note of this initial resume state going forward and expect new work summaries to be provided from the user in the future. 
		Please always incorporate the new work summaries into the resume only if it is significant enough. 
		Given the following valid example JSON bounded within ~~~.
		~~~
{
	"userId": "a-user-id",
	"skills": [
		{
			"name": "AWS Lambda",
			"yearsOfExperience": 4
		}
	],
	"workSummaries": [
		{
			"title": "Software engineer",
			"startDate": "December 2020",
			"endDate": "Present",
			"company": "Amazon",
			"work": [
				"Developed an API for tracking return-to-office attendance"
			]
		}
	],
	"personalProjects": [
		{
			"name": "Google AI Hackathon",
			"role": "Software Engineer",
			"startDate": "December 2018",
			"endDate": "Jan 2019",
			"contributions": [
				"Integrated backend server with Google Gemini client",
				"Implemented dark mode on frontend ui"
			],
			"notableAchievements": "Won best original idea"
		}
	],
	"educations": [
		{
			"institution": "University of British Columbia",
			"certification": "Bachelor of Science, Major in Computer Science",
			"certifiedDate": "December 2020",
			"notableAchievements": [
				"First place in internal hackathon in 2018",
				"First place in internal hackathon in 2019"
			]
		}
	]
}
		~~~
		Please generate a similar JSON object for the work that the user has done based on the given resume extracted text, remembering to remove ~~~. 
	`, userId, extractedResumeText)
}

func GetAddExperiencePrompt(userId string, promptHistory []PromptHistoryDocument, newExperience string) string {
	var promptBuilder strings.Builder

	// Setup all of the previous user prompts. This includes:
	// - the initial context prompt with system prompt that the model is to act as a resume maintenance and enhancement bot
	// - any previous experience that were added
	// - every step's model output with text prompting that this was the model's response before
	for _, doc := range promptHistory {
		if doc.Role == "user" {
			promptBuilder.WriteString(fmt.Sprintf("On %v, the user has reported the following", doc.CreatedAt))
			promptBuilder.WriteString(doc.Text)
			promptBuilder.WriteString("\n\n")
		}

		if doc.Role == "model" {
			promptBuilder.WriteString(fmt.Sprintf("At this point, you have concluded that the below are the details of the resume in valid JSON escaped string format: %s", doc.Text))
			promptBuilder.WriteString("\n\n")
		}
	}

	// Add the new experience prompt into the end of the prompt
	promptBuilder.WriteString(
		fmt.Sprintf(`
			The user has reported another skill, work summary, personal project, or education update. The update is bound within three tilde
			~~~
			%s
			~~~
			Please take note of this report and generate the newest, most appropriate JSON object for the work that the user has done based on all of the context above. 
		`, newExperience))

	return promptBuilder.String()
}

func GetGenerateResumePrompt(userId string, resumeDetails string) string {
	return fmt.Sprintf(`
	Given the following JSON object depicting the user %s resume details bounded within three tildes
	~~~
	%s
	~~~
	please generate a resume for applying to Software Engineer related roles with the objective to get past applicant tracking systems (ASTs) used by many companies out there. 
	We want the best chances for the user to succeed in getting a call back from their job applications.

	Please limit the length of the resume to one US letter size page, or 1500 to 1800 characters. Please generate this in escaped text format and include any unicode characters as you see fit.
	`, userId, resumeDetails)
}
