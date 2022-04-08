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
  } from "@mantine/core";
  import { Link } from "react-router-dom";
  
  interface AuctionConfirmCreationProps {
    onOpen: boolean;
    onClose: () => void;
  }
  
  export default function AuctionConfirmCreation({
    onOpen,
    onClose
  }: AuctionConfirmCreationProps) {
    return (
        <Modal
        opened={onOpen}
        onClose={onClose}
        title="Aukciono kūrimas"
      >
          <Progress value={0}></Progress>
      </Modal>
    );
  }
  