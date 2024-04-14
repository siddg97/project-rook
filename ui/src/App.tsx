import { NextUIProvider } from '@nextui-org/react';
import {
  NavigateFunction,
  Outlet,
  Route,
  Routes,
  useNavigate,
} from 'react-router-dom';
import { PageContainer } from './components';
import { routes } from './routes.tsx';

function App() {
  const navigate: NavigateFunction = useNavigate();

  return (
    <NextUIProvider navigate={navigate}>
      <main
        className={
          'dark bg-accent-gr bg-auto w-full min-h-screen text-foreground bg-background'
        }
      >
        <Routes>
          <Route
            path='/'
            element={
              <PageContainer>
                <Outlet />
              </PageContainer>
            }
          >
            {routes.map(r => (
              <Route
                index={r.path === '/'}
                path={r.path}
                element={r.component}
              />
            ))}
            {/*<Route path='*' element={<NoMatch />} />*/}
          </Route>
        </Routes>
      </main>
    </NextUIProvider>
  );
}

export default App;
