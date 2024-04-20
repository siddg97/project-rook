import { ScrollShadow } from '@nextui-org/react';
import { Conversation, PromptInputWithActions } from '../components';

function Home() {
  return (
    <div className='w-full h-full'>
      <div className='relative flex h-full flex-col'>
        <ScrollShadow className='flex h-full max-h-[60vh] flex-col gap-6 overflow-y-auto pb-8'>
          <Conversation />
          <Conversation />
        </ScrollShadow>
        <div className='mt-auto flex max-w-full flex-col gap-2'>
          <PromptInputWithActions />
          <p className='px-2 text-tiny text-default-400'>
            Acme AI can make mistakes. Consider checking important information.
          </p>
        </div>
      </div>
    </div>
  );
}

export { Home };
