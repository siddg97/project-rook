import { useGetResumeSummary } from '../hooks/useGetResumeSummary.tsx';
import { useStore } from '../hooks/useStore';
import { ResumeSummary } from '../models/ResumeSummary.ts';
import { randomElementFrom } from '../utils';
import {
  Accordion,
  AccordionItem,
  Card,
  CardBody,
  CardHeader,
  Chip,
  Divider,
  Spacer,
  Spinner,
} from '@nextui-org/react';

function Summary() {
  const { auth } = useStore();

  const {
    data: resumeSummaryResponse,
    error: resumeSummaryResponseError,
    isFetching: isFetchingResumeSummary,
    isError: isErrorResumeSummary,
    isLoading: isLoadingResumeSummary,
  } = useGetResumeSummary({ userId: auth.authenticatedUser.uid });

  if (isLoadingResumeSummary || isFetchingResumeSummary) {
    return <Spinner />;
  }

  if (isErrorResumeSummary) {
    return (
      <div className='w-full h-full'>
        {JSON.stringify(resumeSummaryResponseError)}
      </div>
    );
  }

  const promptHistory = (
    <Accordion variant='bordered'>
      {resumeSummaryResponse?.data?.promptHistory.map(
        (promptHistory, index) => (
          <AccordionItem
            key={index}
            title={`${promptHistory.role} - ${promptHistory.id} - ${promptHistory.createdAt}`}
          >
            {promptHistory.text}
          </AccordionItem>
        )
      )}
    </Accordion>
  );

  const latestResumeDetails: ResumeSummary = JSON.parse(
    resumeSummaryResponse.data.promptHistory[
      resumeSummaryResponse.data.promptHistory.length - 1
    ].text
  ) as ResumeSummary;

  const resumeProfile = (
    <Card isBlurred isHoverable className='w-full'>
      <CardHeader className='flex flex-col items-start'>
        <p className='text-xl text-default-900'>
          {latestResumeDetails.profile.name}
        </p>
        <p className='text-small text-default-500'>
          {latestResumeDetails.profile.email}
        </p>
        <p className='text-small text-default-500'>
          {latestResumeDetails.profile.phone}
        </p>
      </CardHeader>
      <Divider />
      <CardBody>
        <div className='flex flex-col'>
          <p className='text-default-600'>{latestResumeDetails.summary}</p>
        </div>
      </CardBody>
    </Card>
  );

  const resumeSkills = (
    <Card isBlurred isHoverable className='w-full'>
      <CardHeader className='text-lg text-default-900'>Skills</CardHeader>
      <Divider />
      <CardBody>
        {Object.entries(latestResumeDetails.skills).map(
          ([category, skills]) => {
            const categoryColor = randomElementFrom([
              'warning',
              'success',
              'danger',
              'primary',
              'secondary',
            ]);
            return (
              <div key={`category-${category}`} className='pt-2'>
                <p className='text-lg text-default-500'>{category}</p>
                <div className='my-auto flex flex-row gap-2 flex-wrap'>
                  {skills.map(skill => (
                    <Chip
                      key={`skill-${category}-${skill}`}
                      color={categoryColor}
                      variant='dot'
                    >
                      {skill}
                    </Chip>
                  ))}
                </div>
              </div>
            );
          }
        )}
      </CardBody>
    </Card>
  );

  const resumeExperience = (
    <Card isBlurred isHoverable className='w-full'>
      <CardHeader className='text-lg text-default-900'>Experience</CardHeader>
      <Divider />
      <CardBody>{JSON.stringify(latestResumeDetails.experience)}</CardBody>
    </Card>
  );

  const resumeEducation = (
    <Card isBlurred isHoverable className='w-full'>
      <CardHeader className='text-lg text-default-900'>Education</CardHeader>
      <Divider />
      <CardBody>{JSON.stringify(latestResumeDetails.education)}</CardBody>
    </Card>
  );

  const latestResume = (
    <>
      {resumeProfile}
      {resumeSkills}
      {resumeExperience}
      {resumeEducation}
    </>
  );

  return (
    <div className='w-full min-h-full'>
      {latestResume}
      <Spacer y={6} />
      {promptHistory}
    </div>
  );
}

export { Summary };
