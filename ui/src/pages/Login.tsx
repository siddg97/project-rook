import { Button } from '@nextui-org/react';
import { AnimatePresence, motion } from 'framer-motion';
import { Icon } from '@iconify/react';
import { NavigateFunction, useLocation, useNavigate } from 'react-router-dom';
import { useStore } from '../hooks/useStore.ts';
import { GoogleAuthProvider, getAuth, signInWithPopup } from 'firebase/auth';

function Login() {
  const navigate: NavigateFunction = useNavigate();
  const location = useLocation();
  const firebase = useStore(state => state.firebase);
  const auth = useStore(state => state.auth);

  const onLogin = async () => {
    try {
      const firebaseAuth = getAuth(firebase.app);
      const result = await signInWithPopup(
        firebaseAuth,
        new GoogleAuthProvider()
      );

      // The signed-in user info.
      const user = result.user;
      auth.setAuthenticatedUser(user);

      // Check for original path that user requested
      const state = location.state;
      const redirectPath = state?.originalPath || '/';
      navigate(redirectPath);
    } catch (e) {
      console.log('login failed', e);
      auth.clearAuthenticatedUser();
    }
  };

  return (
    <div className='flex h-full w-full items-center justify-center'>
      <motion.div
        layout
        className='flex h-full w-full max-w-sm flex-col gap-4 rounded-large bg-content1 px-6 pb-6 pt-6 shadow-small'
      >
        <AnimatePresence>
          <div className='flex flex-col gap-2'>
            <div className='flex flex-col gap-2'>
              <Button
                startContent={
                  <Icon icon='flat-color-icons:google' width={24} />
                }
                onClick={onLogin}
                variant='flat'
              >
                Continue with Google
              </Button>
            </div>
          </div>
        </AnimatePresence>
      </motion.div>
    </div>
  );
}

export { Login };
