import { Button, Spinner } from '@nextui-org/react';
import { ChangeEvent, useState } from 'react';
import { uploadResume } from '../utils';
import { useStore } from '../hooks/useStore';

function Home() {
  const { auth } = useStore();

  const [isUploadingResume, setIsUploadingResume] = useState<boolean>(false);

  const handleFileChange = async (e: ChangeEvent<HTMLInputElement>) => {
    const hasOneFile: boolean = e.target.files?.length === 1;
    if (hasOneFile) {
      setIsUploadingResume(true);
      const selectedFile: File = e.target.files[0];
      const uid = auth.authenticatedUser.uid;
      const resp = await uploadResume(selectedFile, uid);
      setIsUploadingResume(false);
      console.log(resp.data);
    }
  };

  return (
    <div className='w-full h-full'>
      {isUploadingResume ? (
        <Spinner />
      ) : (
        <Button color='primary' fullWidth={false}>
          <label
            htmlFor='upload-resume'
            className='cursor-pointer inline-block w-full'
          >
            Upload Resume
          </label>
          <input
            id='upload-resume'
            type='file'
            className='hidden'
            onChange={handleFileChange}
          />
        </Button>
      )}
    </div>
  );
}

export { Home };
