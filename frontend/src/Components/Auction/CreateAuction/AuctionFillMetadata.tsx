import {
  Card,
  Image,
  Text,
  Badge,
  Button,
  Group,
  useMantineTheme,
  Modal,
  Progress,
  Select,
  TextInput,
  Title,
} from "@mantine/core";
import {useState } from "react";
import { Link } from "react-router-dom";

interface AuctionMetadataProps {
  onOpen: boolean;
  onClose: () => void;
}

export default function AuctionMetadata({
  onOpen,
  onClose
}: AuctionMetadataProps) {
  const [opened, setOpened] = useState(false);
  return (
      <Modal
      opened={onOpen}
      onClose={onClose}
      title="Aukciono kūrimas"
    >
        <Title order={6}>Aukciono duomenys</Title>
        <Progress value={33}></Progress>
        <Select
          label="Daiktas"
          placeholder="Pasirinkti"
          data={[
              { value: 'Dviratis', label: 'Dviratis' },
              { value: 'Tankas', label: 'Tankas' },
          ]}
          />
        <Button style={{left: 330, top: 10}}>Tęsti</Button>
    </Modal>
  );
}
