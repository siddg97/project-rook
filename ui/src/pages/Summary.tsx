import { useEffect, useState } from 'react';
import { useStore } from '../hooks/useStore';
import { getExperience } from '../utils/api-utils';
import { ExperienceHistory, PromptHistory } from '../models/ExperienceSummary';
import {
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from '@nextui-org/react';

function Summary() {
  const [experienceHistory, setExperienceHistory] =
    useState<ExperienceHistory | null>(null);
  const {
    auth,
    env: { local },
  } = useStore();

  useEffect(() => {
    const uid = auth.authenticatedUser ? auth.authenticatedUser.uid : null;
    if (uid) {
      getExperience(uid, local)
        .then(response => {
          const history = {
            promptHistory: response.data.promptHistory.sort(
              (h1: PromptHistory, h2: PromptHistory) =>
                h2.createdAt.localeCompare(h1.createdAt)
            ),
          };

          setExperienceHistory(response.data);
        })
        .catch(err => console.log(`Failed to get experience. Reason: ${err}`));
    }
  }, []);

  return (
    <div className='w-full h-full'>
      <div className='relative flex h-full flex-col'>
        {experienceHistory ? (
          <Table isStriped>
            <TableHeader>
              <TableColumn key='createdAt'>DATE</TableColumn>
              <TableColumn key='value'>PROMPT</TableColumn>
            </TableHeader>
            <TableBody
              emptyContent={'No history to display.'}
              items={experienceHistory?.promptHistory}
            >
              {(item: PromptHistory) => (
                <TableRow key={item.id}>
                  <TableCell>{item.createdAt}</TableCell>
                  <TableCell>{item.text}</TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        ) : null}
      </div>
    </div>
  );
}

export { Summary };
