export interface ExperienceHistory {
  resume: Resume;
  promptHistory: PromptHistory[];
}

export interface Resume {
  userId: string;
  resumeId: string;
  resumeText: string;
}

export interface PromptHistory {
  id: string;
  createdAt: string;
  role: string;
  text: string;
}
