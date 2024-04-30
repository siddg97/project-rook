export interface ExperienceHistory {
    promptHistory: Array<PromptHistoryEntry>
}

export interface PromptHistoryEntry {
    id: string,
    createdAt: string,
    value: string,
}