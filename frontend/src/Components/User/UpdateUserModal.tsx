import {
    Modal,
    Title,
    Divider,
    Button,
    Center,
    Text,
    TextInput,
    PasswordInput,
  } from "@mantine/core";
  import { useState, useEffect } from "react";
  import { useForm } from "@mantine/form";
  
  interface ItemProps {
    onOpen: boolean;
    onClose: () => void;
    userID: string;
  }
  
  export default function UpdateUserModal({ onOpen, onClose, userID}: ItemProps) {
    return (
      <Modal opened={onOpen} onClose={onClose} size="xl">
        <Title>Vartotojo informacija</Title>
        <Divider />
        <TextInput
        label="Vartotojo slapyvardis"
        value={userID}
        disabled
      />
        <PasswordInput
        label="Slaptažodis"/>
        <PasswordInput
        label="Pakartokite slaptažodį"/>
        <Center>
        <Button color="green" onClick={() => {
          onClose()
        }}>Atnaujinti</Button>
        </Center>
      </Modal>
    );
  }
  