import {
  Modal,
  Title,
  Divider,
  Button,
  Center,
  Text,
} from "@mantine/core";
import {X, ChevronDown} from "tabler-icons-react";
import { showNotification } from '@mantine/notifications';

interface ItemProps {
  onOpen: boolean;
  onClose: () => void;
  auctionID: string;
}

export default function DeleteAuction({ onOpen, onClose, auctionID }: ItemProps) {
  const deleteAuction = async (auctionID: string) => {
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
      `${process.env.REACT_APP_API_URL}auctions/${auctionID}`;

    try {
      const response = await fetch(url, requestOptions);
      if(response.status == 204){
        showNotification({
          title: 'Aukciono pašalinimas',
          color: 'green',
          icon: <ChevronDown/>,
          message: 'Aukcionas sėkmingai pašalintas',
        })
        onClose()
      } else{
        showNotification({
          title: 'Aukciono pašalinimas',
          message: 'Aukciono nepavyko pašalinti',
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
      <Title>Aukciono panaikinimas</Title>
      <Divider />
      <Text><b>Primename</b>: Pašalinti aukcioną galima tik kai jis nėra prasidėjęs ar pasibaigęs.</Text>
      Ar tikrai norite pašalinti aukcioną?
      <br/>
      <Center>
      <Button color="red" onClick={() => {
        deleteAuction(auctionID)
      }}>Pašalinti</Button>
      </Center>
    </Modal>
  );
}
