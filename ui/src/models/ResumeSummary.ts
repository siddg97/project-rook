export interface ResumeSummary {
  id: string;
  profile: Profile;
  summary: string;
  skills: Record<string, string[]>;
  education: Education;
  experience: Experience[];
}

export interface Profile {
  name: string;
  website: any;
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
}

export interface Experience {
  company: string;
  location: any;
  positions: Position[];
}

export interface Position {
  title: string;
  duration: string;
  responsibilities: string[];
}
