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
import { AppContext, initializeFirebase } from './contexts/AppContext.ts';

function App() {
  const firebaseApp = initializeFirebase();
  const navigate: NavigateFunction = useNavigate();

  return (
    <AppContext.Provider value={{firebaseApp}}>
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
                  key={r.key}
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
    </AppContext.Provider>
  );
}

export default App;
