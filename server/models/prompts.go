package models

import (
	"fmt"
)

type ModelResumeDetails struct {
	ID      string `json:"id"`
	Profile struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Location string `json:"location"`
		Summary  string `json:"summary"`
		Website  string `json:"website"`
	} `json:"profile"`
	Summary   string              `json:"summary"`
	Skills    map[string][]string `json:"skills"`
	Education []struct {
		Institution string   `json:"institution"`
		Location    string   `json:"location"`
		Degree      string   `json:"degree"`
		Major       string   `json:"major"`
		Graduation  string   `json:"graduation"`
		Gpa         string   `json:"gpa"`
		Awards      []string `json:"awards"`
	} `json:"education"`
	Experience []struct {
		Company   string `json:"company"`
		Location  string `json:"location"`
		Positions []struct {
			Title            string   `json:"title"`
			Duration         string   `json:"duration"`
			Responsibilities []string `json:"responsibilities"`
		} `json:"positions"`
	} `json:"experience"`
	Certifications []string `json:"certifications"`
	Additional     struct {
		Cfa       string   `json:"cfa"`
		Interests []string `json:"interests"`
		Languages []string `json:"languages"`
	} `json:"additional"`
}

func GetInitialResumeCreationPrompt(userId, extractedResumeText string) string {
	return fmt.Sprintf(`
You are a resume maintainance and enhancement bot that helps user track their work summaries across time and update it to the best of your ability such that it increases the user's chance of receiving a call back from companies they applied for.
Below (bounded by ~~~) is the initial resume text extracted from a resume file uploaded by the user (denoted by id %s)

~~~

%s

~~~

Please keep note of this initial resume state going forward and expect new work summaries to be provided from the user in the future. Can you also summarize this resume in a JSON document. JSON document should have the following keys:
- "id": the id of the user
- "profile": an ojbect that has "name", "website", "phone" and "email" keys
- "summary": a summary of the resume in a few lines
- "skills": an object that has keys that denate a skill category and then the value of these keys to be a lis tof strings
- "education": an array of objects. each object has "institution", "location", "degree", "major", "graduation", "gpa" and "awards" keys
- "experience": array of objects. Each object has "company", "location" and "postions" keys. "positions" is an array of objects that contain "title", "duration", and "responsibilities" keys

Please do not format the response in markdown and no backticks of any sort
`, userId, extractedResumeText)
}

func AddExperiencePrompt(userId string, newExperience string) string {
	return fmt.Sprintf(`
Based on the current state of resume for user %s, please add the folowing new experience (bounded by ~~~) to it

~~~

%s

~~~

Once you have updated the appropriate sections of the resume can you please provide the JSON equivalent of the updated resume as described. Only include JSON document with markdown formatting.
`, userId, newExperience)
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
