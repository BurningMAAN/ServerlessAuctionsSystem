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
  import { Link } from "react-router-dom";
  import {useState } from "react";
  import AuctionMetadata from './AuctionFillMetadata';
  
  interface AuctionSelectItemProps {
    onOpen: boolean;
    onClose: () => void;
  }
  
  export default function AuctionSelectItem({
    onOpen,
    onClose
  }: AuctionSelectItemProps) {
    const [opened, setOpened] = useState(false);
    return (
        <Modal
        opened={onOpen}
        onClose={onClose}
        title="Aukciono kūrimas"
      >
          <Title order={6}> Inventoriaus pasirinkimas</Title>
          <Progress value={0}></Progress>
          <Select
            label="Daiktas"
            placeholder="Pasirinkti"
            data={[
                { value: 'Dviratis', label: 'Dviratis' },
                { value: 'Tankas', label: 'Tankas' },
            ]}
            />
          <Button style={{left: 330, top: 10}} onClick={() => setOpened(true)}>Tęsti</Button>
          <AuctionMetadata onOpen={opened} onClose={onClose}></AuctionMetadata>
      </Modal>
    );
  }
  