import { ReactNode } from 'react';
import Home from './pages/Home.tsx';
import Login from './pages/Login.tsx';

export interface AppRoute {
  key: string;
  path: string;
  navText: string;
  component: ReactNode;
  isProtected: boolean;
}

export const routes: AppRoute[] = [
  {
    key: 'home',
    path: '/',
    navText: 'Home',
    component: <Home />,
    isProtected: true,
  },
  {
    key: 'login',
    path: 'login',
    navText: 'Login',
    component: <Login />,
    isProtected: false,
  },
];
