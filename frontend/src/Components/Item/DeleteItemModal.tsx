import {
    Modal,
    Title,
    Divider,
    Button,
    Center,
    Text,
  } from "@mantine/core";
  import { useState, useEffect } from "react";
  import { useForm } from "@mantine/form";
  
  interface ItemProps {
    onOpen: boolean;
    onClose: () => void;
    itemID: string;
    itemName: string;
  }
  
  export default function DeleteItem({ onOpen, onClose, itemName, itemID }: ItemProps) {
    return (
      <Modal opened={onOpen} onClose={onClose} size="xl">
        <Title>Inventoriaus panaikinimas</Title>
        <Divider />
        <Text><b>Primename</b>: Pašalinti inventorių galima tik kai inventorius nėra priskirtas prie aukciono.</Text>
        Ar tikrai norite pašalinti inventorių: <b>{itemName}</b>?
        <br/>
        <Center>
        <Button color="red" onClick={() => {
          console.log("panaikinta")
          onClose()
        }}>Pašalinti</Button>
        </Center>
      </Modal>
    );
  }
  