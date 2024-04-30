import { useEffect, useState } from "react";
import { useStore } from "../hooks/useStore";
import { getExperience } from "../utils/api-utils";
import { ExperienceHistory, PromptHistoryEntry } from "../models/ExperienceSummary";
import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow, getKeyValue } from "@nextui-org/react";

function Summary() {
  const [experienceHistory, setExperienceHistory] = useState<ExperienceHistory | null>(null);
  const { auth } = useStore();
  
  useEffect(() => {
    const uid = auth.authenticatedUser ? auth.authenticatedUser.uid : null;
    if (uid) {
      getExperience(uid)
        .then((response) => {
          const history = {
            /*eslint-disable @typescript-eslint/no-explicit-any*/
            promptHistory: response.data
              .promptHistory
              .map((history: any) => {
                return {
                  id: history.id,
                  createdAt: history.createdAt,
                  value: history.text,
                };
              })
              .sort(
                (h1: PromptHistoryEntry, h2: PromptHistoryEntry) => h2.createdAt.localeCompare(h1.createdAt)
              )
          };
          
          setExperienceHistory(history);
        })
        .catch((err) => console.log(`Failed to get experience. Reason: ${err}`));
    }
  }, []);

  return (
    <div className='w-full h-full'>
      <div className='relative flex h-full flex-col'>
        {
          experienceHistory ?
            <Table isStriped>
              <TableHeader>
                <TableColumn key="createdAt">DATE</TableColumn>
                <TableColumn key="value">PROMPT</TableColumn>
              </TableHeader>
              <TableBody emptyContent={"No history to display."} items={experienceHistory?.promptHistory}>
                {(item: PromptHistoryEntry) => (
                  <TableRow key={item.id}>
                    <TableCell>{item.createdAt}</TableCell>
                    <TableCell>{item.value}</TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          : null
        }
      </div>
    </div>
  );
}

export { Summary };
