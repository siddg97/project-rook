import { ReactNode } from 'react';
import { AddExperience, Home, Login, Summary } from './pages';

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
    key: 'summary',
    path: '/summary',
    navText: 'Summary',
    component: <Summary />,
    isProtected: true,
  },
  {
    key: 'addExperience',
    path: '/experience/new',
    navText: 'Add experience',
    component: <AddExperience />,
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
