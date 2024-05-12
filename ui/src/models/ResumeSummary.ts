export interface ResumeSummary {
  id: string;
  profile: Profile;
  summary: string;
  skills: Record<string, string[]>;
  education: Education[];
  experience: Experience[];
}

export interface Profile {
  name: string;
  website: string;
  phone: string;
  email: string;
}

export interface Education {
  institution: string;
  location: string;
  degree: string;
  major: string;
  graduation: string;
  gpa: string;
  awards: string[];
}

export interface Experience {
  company: string;
  location: string;
  positions: Position[];
}

export interface Position {
  title: string;
  duration: string;
  responsibilities: string[];
}
