import { useQuery } from '@tanstack/react-query';
import { getExperience } from '../utils';

export interface GetResumeSummaryParams {
  userId: string;
}

export const useGetResumeSummary = (params: GetResumeSummaryParams) => {
  return useQuery({
    queryKey: ['resume', params.userId],
    queryFn: () => getExperience(params.userId),
  });
};
