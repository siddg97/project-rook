import axios, { AxiosResponse } from 'axios';
import { ExperienceHistory } from '../models/ExperienceSummary.ts';

const client = axios.create();
const LOCAL_BASE_URL: string = 'http://localhost:3000';

export function uploadResume(file: File, uid: string, local: boolean) {
  const data = new FormData();
  data.append('file', file);

  client.put(`${local ? LOCAL_BASE_URL : ''}/v1/${uid}/resume`, data, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
}

export function submitExperience(
  experience: string,
  uid: string,
  local: boolean
): Promise<AxiosResponse<object>> {
  const data = {
    experience: experience,
  };
  return client.post(`${local ? LOCAL_BASE_URL : ''}/v1/${uid}/resume`, data);
}

export function getExperience(
  uid: string,
  local: boolean
): Promise<AxiosResponse<ExperienceHistory>> {
  return client.get(`${local ? LOCAL_BASE_URL : ''}/v1/${uid}/resume`, {
    headers: {
      'Content-Type': 'application/json',
    },
  });
}
