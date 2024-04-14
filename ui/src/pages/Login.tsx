import { Button } from '@nextui-org/react';
import { AnimatePresence, motion } from 'framer-motion';
import { Icon } from '@iconify/react';

function Login() {
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

export default Login;
