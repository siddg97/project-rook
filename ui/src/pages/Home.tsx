import { ScrollShadow, Tab, Tabs } from '@nextui-org/react';
import {
  Conversation,
  PromptInputWithActions,
  SidebarWithGradient,
} from '../components';

function Home() {
  return (
    <div className='h-full w-full max-w-full'>
      <SidebarWithGradient
        header={
          <Tabs className='justify-center'>
            <Tab key='creative' title='Creative' />
            <Tab key='technical' title='Technical' />
            <Tab key='precise' title='Precise' />
          </Tabs>
        }
        title="Creative Uses for Kids' Art"
      >
        <div className='relative flex h-full flex-col'>
          <ScrollShadow className='flex h-full max-h-[60vh] flex-col gap-6 overflow-y-auto pb-8'>
            <Conversation />
            <Conversation />
          </ScrollShadow>
          <div className='mt-auto flex max-w-full flex-col gap-2'>
            <PromptInputWithActions />
            <p className='px-2 text-tiny text-default-400'>
              Acme AI can make mistakes. Consider checking important
              information.
            </p>
          </div>
        </div>
      </SidebarWithGradient>
    </div>
  );
}

export default Home;
