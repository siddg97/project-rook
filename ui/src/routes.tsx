import Home from './pages/Home.tsx';
import Login from './pages/Login.tsx';

export const routes = [
  {
    path: '/',
    navText: 'Home',
    component: <Home />,
  },
  {
    path: 'login',
    navText: 'Login',
    component: <Login />,
  },
];
