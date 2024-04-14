import { ReactNode } from 'react';

type Props = {
  children: ReactNode;
};

function PageContainer({ children }: Props) {
  return (
    <div className='flex min-h-screen w-full justify-center justify-items-center items-center container mx-auto max-w-7xl px-6 flex-grow'>
      {children}
    </div>
  );
}

export default PageContainer;
