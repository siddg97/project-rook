import { FC, ReactNode } from 'react';
import { useStore } from '../hooks/useStore.ts';
import { Navigate } from 'react-router-dom';

interface Props {
  children: ReactNode;
  path: string;
}

const AuthenticatedRoute: FC<Props> = ({ children, path }) => {
  const auth = useStore(state => state.auth);

  return auth.isLoggedIn ? (
    children
  ) : (
    <Navigate to={'/login'} state={{ originalPath: path }} />
  );
};

export default AuthenticatedRoute;
