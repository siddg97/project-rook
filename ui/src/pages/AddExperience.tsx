import {
  Button,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  Textarea,
} from '@nextui-org/react';
import { useState } from 'react';
import { useStore } from '../hooks/useStore';
import { submitExperience } from '../utils';

function AddExperience() {
  const [experience, setExperience] = useState<string>('');
  const [isSubmittingExperience, setIsSubmittingExperience] =
    useState<boolean>(false);
  const [showExperienceSubmittedModal, setShowExperienceSubmittedModal] =
    useState<boolean>(false);
  const [showExperienceSubmitFailedModal, setShowExperienceSubmitFailedModal] =
    useState<boolean>(false);
  const { auth } = useStore();

  const handleSubmitExperience = async () => {
    const uid = auth.authenticatedUser ? auth.authenticatedUser.uid : null;
    if (uid) {
      setIsSubmittingExperience(true);
      try {
        const resp = await submitExperience(experience, uid);
        console.log(`Submitted experience. Response: ${JSON.stringify(resp)}`);
        setIsSubmittingExperience(false);
        setShowExperienceSubmittedModal(true);
      } catch (err) {
        console.log(`Failed to submit experience due to: ${err}`);
        setIsSubmittingExperience(false);
        setShowExperienceSubmitFailedModal(true);
      }
    }
  };

  return (
    <div className='w-full h-full'>
      <div className='relative flex h-full flex-col'>
        <Modal
          isOpen={showExperienceSubmittedModal}
          onClose={() => setShowExperienceSubmittedModal(false)}
        >
          <ModalContent>
            <ModalBody>Successfully submitted experience!</ModalBody>
            <ModalFooter>
              <Button
                color='danger'
                variant='light'
                onPress={() => setShowExperienceSubmittedModal(false)}
              >
                Close
              </Button>
            </ModalFooter>
          </ModalContent>
        </Modal>
        <Modal
          isOpen={showExperienceSubmitFailedModal}
          onClose={() => setShowExperienceSubmitFailedModal(false)}
        >
          <ModalContent>
            <ModalBody>Failed to submit experience</ModalBody>
            <ModalFooter>
              <Button
                color='danger'
                variant='light'
                onPress={() => setShowExperienceSubmitFailedModal(false)}
              >
                Close
              </Button>
            </ModalFooter>
          </ModalContent>
        </Modal>
        <Textarea
          isRequired
          label='Experience'
          labelPlacement='outside'
          placeholder='Describe the experience (e.g. Rested & vested)'
          onValueChange={(value: string) => setExperience(value)}
        />
        <Button
          color='primary'
          onPress={handleSubmitExperience}
          isLoading={isSubmittingExperience}
        >
          Submit
        </Button>
      </div>
    </div>
  );
}

export { AddExperience };
