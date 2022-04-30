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
  
  export default function DeleteAuction({ onOpen, onClose }: ItemProps) {
    return (
      <Modal opened={onOpen} onClose={onClose} size="xl">
        <Title>Aukciono panaikinmas</Title>
        <Divider />
      </Modal>
    );
  }
  