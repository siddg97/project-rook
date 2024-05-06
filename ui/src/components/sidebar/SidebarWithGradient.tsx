import {
  Avatar,
  Button,
  ScrollShadow,
  Spacer,
  useDisclosure,
} from '@nextui-org/react';
import { Icon } from '@iconify/react';
import { getAuth } from 'firebase/auth';
import { ReactNode } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { sectionItems } from '../../constants';
import { useStore } from '../../hooks/useStore.ts';
import { Sidebar } from './Sidebar.tsx';
import { SidebarDrawer } from './SidebarDrawer.tsx';

interface Props {
  children?: ReactNode;
  header?: ReactNode;
}

const pathToPageTitleMappings: Record<string, string> = {
  '/': 'Upload resume',
  '/experience/new': 'Add experience',
  '/summary': 'Summary',
};

function SidebarWithGradient({ children, header }: Props) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const navigate = useNavigate();
  const location = useLocation();
  const auth = useStore(state => state.auth);
  const firebase = useStore(state => state.firebase);

  const rookIcon = (
    <div className='flex items-center gap-2 px-2'>
      <div className='flex h-8 w-8 items-center justify-center rounded-full border-small border-foreground/20'>
        <Icon
          width={20}
          height={20}
          icon='tabler:chess-rook-filled'
          className='text-foreground'
        />
      </div>
      <span className='text-small font-medium uppercase text-foreground'>
        Rook
      </span>
    </div>
  );

  const userDetails = (
    <div className='flex flex-col gap-4'>
      <div className='flex items-center gap-3 px-2'>
        <Avatar size='sm' src={auth.authenticatedUser?.photoURL as string} />
        <div className='flex flex-col'>
          <p className='text-small text-foreground'>
            {auth.authenticatedUser?.displayName}
          </p>
          <p className='text-tiny text-default-500'>
            {auth.authenticatedUser?.email as string}
          </p>
        </div>
      </div>
    </div>
  );

  const onLogOut = async () => {
    try {
      const firebaseAuth = getAuth(firebase.app);
      await firebaseAuth.signOut();
    } catch (e) {
      console.log('logout failed', e);
      console.log('redirecting to login page and clearing state anyways');
    } finally {
      auth.clearAuthenticatedUser();
      navigate('/login');
    }
  };

  const drawerContent = (
    <div className='relative flex min-h-screen w-72 flex-1 flex-col bg-gradient-to-b from-default-100 via-danger-100 to-secondary-100 p-6'>
      {rookIcon}
      <Spacer y={8} />
      {userDetails}

      <ScrollShadow className='-mr-6 h-full min-h-full py-6 pr-6'>
        <Sidebar
          defaultSelectedKey={
            sectionItems.find(i => i.href === location.pathname)?.key || 'home'
          }
          iconClassName='text-default-600 group-data-[selected=true]:text-foreground'
          itemClasses={{
            base: 'data-[selected=true]:bg-default-400/20 data-[hover=true]:bg-default-400/10',
            title:
              'text-default-600 group-data-[selected=true]:text-foreground',
          }}
          items={sectionItems}
          sectionClasses={{
            heading: 'text-default-600 font-medium',
          }}
          variant='flat'
        />
        <Button
          className='justify-start text-default-600 data-[hover=true]:text-black'
          startContent={
            <Icon
              className='rotate-180 text-default-600'
              icon='solar:minus-circle-line-duotone'
              width={24}
            />
          }
          variant='light'
          onClick={onLogOut}
        >
          Log Out
        </Button>
      </ScrollShadow>
    </div>
  );

  return (
    <div className='flex min-h-full w-full'>
      <SidebarDrawer
        className='flex-none'
        isOpen={isOpen}
        onOpenChange={onOpenChange}
      >
        {drawerContent}
      </SidebarDrawer>
      <div className='flex w-full min-h-full flex-col gap-y-4 p-4 sm:max-w-[calc(100%_-_288px)]'>
        <header className='flex h-16 min-h-16 items-center justify-between gap-2 overflow-x-scroll rounded-medium border-small border-divider px-4 py-2'>
          <div className='flex max-w-full items-center gap-2'>
            <Button
              isIconOnly
              className='flex sm:hidden'
              size='sm'
              variant='light'
              onPress={onOpen}
            >
              <Icon
                className='text-default-500'
                height={24}
                icon='solar:hamburger-menu-outline'
                width={24}
              />
            </Button>
            <h2 className='truncate text-medium font-medium text-default-700'>
              {pathToPageTitleMappings[location.pathname]}
            </h2>
          </div>
          {header}
        </header>
        <main className='flex h-full'>
          <div className='flex h-full w-full flex-col gap-4 rounded-medium border-small border-divider p-6'>
            {children}
          </div>
        </main>
      </div>
    </div>
  );
}

export { SidebarWithGradient };
