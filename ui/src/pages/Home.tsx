import { Button, Input } from '@nextui-org/react';
import { useState } from 'react';
import { uploadResume } from '../utils/api-utils';
import { useStore } from '../hooks/useStore';

function Home() {
  const [file, setFile] = useState<File | null>(null);
  const { auth } = useStore();
  
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const selectedFile: File = e.target.files[0];
      setFile(selectedFile);
    }
  };

  const handleFileUpload = async () => {
    if (file) {
      const uid = auth.authenticatedUser ? auth.authenticatedUser.uid : null;
      if (uid) {
        uploadResume(file, uid);
      }
    }
  }

  return (
    <div className='w-full h-full'>
      <div className='relative flex h-full flex-col'>
        <Input type="file" label="Choose a file" onChange={handleFileChange}/>
        {file && (
        <section>
          File details:
          <ul>
            <li>Name: {file.name}</li>
            <li>Type: {file.type}</li>
            <li>Size: {file.size} bytes</li>
          </ul>
        </section>
      )}
        <Button color="primary" onClick={handleFileUpload}>
          Upload Resume
        </Button>
      </div>
    </div>
  );
}

export { Home };
