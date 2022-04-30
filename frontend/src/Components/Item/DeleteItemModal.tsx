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
  
  export default function DeleteItem({ onOpen, onClose }: ItemProps) {
    return (
      <Modal opened={onOpen} onClose={onClose} size="xl">
        <Title>Inventoriaus panaikinimas</Title>
        <Divider />
      </Modal>
    );
  }
  