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
  import {X} from "tabler-icons-react";
  import { showNotification } from '@mantine/notifications';
  
  interface ItemProps {
    onOpen: boolean;
    onClose: () => void;
    itemID: string;
    itemName: string;
  }
  
  export default function DeleteItem({ onOpen, onClose, itemName, itemID }: ItemProps) {
    const deleteItem = async (itemID: string) => {
      let tokenas: string = "";
      const token = sessionStorage.getItem("access_token");
      if (token) {
        tokenas = token;
      }
    
      const requestOptions = {
        method: "DELETE",
        headers: { access_token: unescape(tokenas) },
      };
      const url =
        `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/item/${itemID}`;
  
      try {
        const response = await fetch(url, requestOptions);
        const responseJSON = await response.json();
        if(response.status == 200){
          showNotification({
            title: 'Inventoriaus pašalinimas',
            message: 'Inventorius sėkmingai pašalintas',
          })
          onClose()
        } else{
          showNotification({
            title: 'Inventoriaus pašalinimas',
            message: 'Inventoriaus nepavyko pašalinti',
            color: 'red',
            icon: <X/>,
          })
        }
      } catch (error) {
        console.log("failed to delete item", error);
      }
    };
    return (
      <Modal opened={onOpen} onClose={onClose} size="xl">
        <Title>Inventoriaus panaikinimas</Title>
        <Divider />
        <Text><b>Primename</b>: Pašalinti inventorių galima tik kai inventorius nėra priskirtas prie aukciono.</Text>
        Ar tikrai norite pašalinti inventorių: <b>{itemName}</b>?
        <br/>
        <Center>
        <Button color="red" onClick={() => {
          deleteItem(itemID)
        }}>Pašalinti</Button>
        </Center>
      </Modal>
    );
  }
  