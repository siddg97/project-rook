import Home from './pages/Home.tsx';
import Login from './pages/Login.tsx';

export const routes = [
  {
    key: 'home',
    path: '/',
    navText: 'Home',
    component: <Home />,
  },
  {
    key: 'login',
    path: 'login',
    navText: 'Login',
    component: <Login />,
  },
];
