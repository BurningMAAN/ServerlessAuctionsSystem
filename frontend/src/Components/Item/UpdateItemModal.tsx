import {
  Modal,
  Title,
  Divider,
} from "@mantine/core";
import { useState, useEffect } from "react";
import { useForm } from "@mantine/form";

interface ItemProps {
  onOpen: boolean;
  onClose: () => void;
}

export default function UpdateItem({ onOpen, onClose }: ItemProps) {
  return (
    <Modal opened={onOpen} onClose={onClose} size="xl">
      <Title>Inventoriaus atnaujinimas</Title>
      <Divider />
    </Modal>
  );
}
