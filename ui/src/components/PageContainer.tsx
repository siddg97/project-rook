import { ReactNode } from 'react';

type Props = {
  children: ReactNode;
};

function PageContainer({ children }: Props) {
  return (
    <div className='flex min-h-screen min-w-full justify-center justify-items-center items-center container mx-auto flex-grow'>
      {children}
    </div>
  );
}

export default PageContainer;
