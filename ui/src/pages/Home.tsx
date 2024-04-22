import { Button, Input } from '@nextui-org/react';
import { useState } from 'react';
import { uploadResume } from '../utils/api-utils';

function Home() {
  const [file, setFile] = useState<File | null>(null);
  
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const selectedFile: File = e.target.files[0];
      setFile(selectedFile);
    }
  };

  const handleFileUpload = async () => {
    if (file) {
      uploadResume(file);
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
