import { ReactNode } from 'react';
import { useStore } from '../hooks/useStore.ts';
import { SidebarWithGradient } from './sidebar';

type Props = {
  children: ReactNode;
};

function PageContainer({ children }: Props) {
  const auth = useStore(state => state.auth);

  return (
    <div className='flex min-h-screen min-w-full justify-center justify-items-center items-center container mx-auto flex-grow'>
      {!auth.isLoggedIn ? (
        children
      ) : (
        <SidebarWithGradient>{children}</SidebarWithGradient>
      )}
    </div>
  );
}

export default PageContainer;
