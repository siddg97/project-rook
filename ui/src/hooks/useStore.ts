import { FirebaseApp } from 'firebase/app';
import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import { immer } from 'zustand/middleware/immer';
import { User } from 'firebase/auth';
import { initializeFirebase } from '../utils';
import { isLocalEnv } from '../utils';

interface AppStore {
  env: {
    local: boolean;
  };
  firebase: {
    app: FirebaseApp;
  };
  auth: {
    isLoggedIn: boolean;
    authenticatedUser: User | null;

    setAuthenticatedUser: (user: User) => void;
    clearAuthenticatedUser: () => void;
  };
}

export const useStore = create<AppStore>()(
  devtools(
    immer(_set => ({
      env: {
        local: isLocalEnv(),
      },
      firebase: {
        app: initializeFirebase(),
      },
      auth: {
        isLoggedIn: false,
        authenticatedUser: null,

        setAuthenticatedUser: user =>
          _set(state => {
            state.auth.isLoggedIn = true;
            state.auth.authenticatedUser = user;
          }),
        clearAuthenticatedUser: () =>
          _set(state => {
            state.auth.isLoggedIn = false;
            state.auth.authenticatedUser = null;
          }),
      },
    })),
    {
      name: 'rook',
    }
  )
);
