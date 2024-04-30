import axios, { AxiosResponse } from 'axios';
import { ExperienceHistory } from '../models/ExperienceSummary.ts';

const client = axios.create();
const BASE_URL = 'http://localhost:3000';

export function uploadResume(file: File, uid: string) {
  const data = new FormData();
  data.append('file', file);

  client.put(`${BASE_URL}/v1/${uid}/resume`, data, {
    // TODO: Replace base URL with env
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
}

export function submitExperience(
  experience: string,
  uid: string
): Promise<AxiosResponse<object>> {
  const data = {
    experience: experience,
  };
  return client.post(`${BASE_URL}/v1/${uid}/resume`, data);
}

export function getExperience(
  uid: string
): Promise<AxiosResponse<ExperienceHistory>> {
  return client.get(`${BASE_URL}/v1/${uid}/resume`, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
}
