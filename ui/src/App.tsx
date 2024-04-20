import { NextUIProvider } from '@nextui-org/react';
import { ReactNode } from 'react';
import {
  NavigateFunction,
  Outlet,
  Route,
  Routes,
  useNavigate,
} from 'react-router-dom';
import { AuthenticatedRoute, PageContainer } from './components';
import { AppRoute, routes } from './routes.tsx';

function renderRoute(route: AppRoute): ReactNode {
  const { key, path, isProtected, component } = route;

  if (!isProtected) {
    return <Route path={path} key={key} element={component} />;
  }

  return (
    <Route
      path={path}
      key={key}
      element={<AuthenticatedRoute path={path}>{component}</AuthenticatedRoute>}
    />
  );
}

function App() {
  const navigate: NavigateFunction = useNavigate();

  return (
    <NextUIProvider navigate={navigate}>
      <main
        className={
          'dark bg-accent-gr bg-auto w-full min-h-screen text-foreground bg-background m-0'
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
            {routes.map(r => renderRoute(r))}
            {/*<Route path='*' element={<NoMatch />} />*/}
          </Route>
        </Routes>
      </main>
    </NextUIProvider>
  );
}

export default App;
