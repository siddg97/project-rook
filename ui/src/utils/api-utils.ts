import axios, { AxiosResponse } from 'axios';
import { ExperienceHistory } from '../models/ExperienceSummary.ts';

const client = axios.create();
const LOCAL_BASE_URL: string = 'http://localhost:3000';

export function isLocalEnv(): boolean {
  return window.location.host.includes('localhost');
}

export function uploadResume(file: File, uid: string) {
  const data = new FormData();
  data.append('file', file);

  return client.put(
    `${isLocalEnv() ? LOCAL_BASE_URL : ''}/v1/${uid}/resume`,
    data,
    {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    }
  );
}

export function submitExperience(
  experience: string,
  uid: string
): Promise<AxiosResponse<object>> {
  const data = {
    experience: experience,
  };
  return client.post(
    `${isLocalEnv() ? LOCAL_BASE_URL : ''}/v1/${uid}/resume`,
    data
  );
}

export function getExperience(
  userId: string
): Promise<AxiosResponse<ExperienceHistory>> {
  const baseUrl: string = isLocalEnv() ? LOCAL_BASE_URL : '';

  return client.get(`${baseUrl}/v1/${userId}/resume`, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
}
