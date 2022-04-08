import {
  Button,
  Modal,
  Stepper,
  Center,
  Title,
  Select,
  NumberInput,
  Divider,
} from "@mantine/core";
import { useState } from "react";
import { DatePicker } from '@mantine/dates';

interface AuctionProps {
  onOpen: boolean;
  onClose: () => void;
}

export default function AuctionCreateWizard({ onOpen, onClose }: AuctionProps) {
  const [activeStep, setActiveStepStepper] = useState(0);
  const nextStep = () =>
    setActiveStepStepper((current) => (current < 3 ? current + 1 : current));
  const prevStep = () =>
    setActiveStepStepper((current) => (current > 0 ? current - 1 : current));

  const handleOnClose = (): void => {
    setActiveStepStepper(0);
    onClose();
  };
  return (
    <Modal opened={onOpen} onClose={handleOnClose} size="xl">
      <Stepper
       active={activeStep}
       color="green">
        <Stepper.Step label="Inventoriaus pasirinkimas"></Stepper.Step>
        <Stepper.Step label="Aukciono duomenys"></Stepper.Step>
        <Stepper.Step label="Patvirtinimas"></Stepper.Step>
      </Stepper>
      <Divider />
      {activeStep == 0 && (
        <>
         <Select
        label="Inventoriaus pasirinkimas"
        placeholder="Pasirinkti"
        data={[{ value: "Dvirka", label: "Dvirka" }]}
        required
      />
          <Divider/>
          <Center>
            <Button onClick={nextStep}>Toliau</Button>
          </Center>
        </>
      )}
      {activeStep == 1 && (
        <>
         <Select
        label="Aukciono tipas"
        placeholder="Pasirinkti"
        required
        data={[{ value: "absoliutus", label: "Absoliutus" }]}
      />
      <DatePicker placeholder="Pasirinkti" label="Aukciono data" required />
        <NumberInput
        label="Išpirkimo kaina"
        placeholder="Įvesti"
      />
        <Divider/>
        <Center>
        <Button onClick={prevStep}>Atgal</Button>
          <Button onClick={nextStep}>Toliau</Button>
        </Center>
      </>
      )}
      {activeStep == 2 && (
        <>
        <Title order={1}>Step 3</Title>
        <Divider/>
        <Center>
        <Button onClick={prevStep}>Atgal</Button>
          <Button color="green" onClick={handleOnClose}>Patvirtinti</Button>
        </Center>
      </>
      )}
    </Modal>
  );
}
