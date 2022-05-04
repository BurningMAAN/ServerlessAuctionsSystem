import {
  Modal,
  Title,
  Divider,
  TextInput,
  Select,
  Textarea,
} from "@mantine/core";
import { useState, useEffect } from "react";
import jwtDecode, { JwtPayload } from "jwt-decode";
import { useForm } from "@mantine/form";

interface ItemProps {
  id: string;
  onOpen: boolean;
  onClose: () => void;
}

interface GetItem {
  id: string;
  description: string;
  category: string;
  name: string;
}

interface DecodedToken {
  username: string;
}

export default function UpdateItem({ id, onOpen, onClose }: ItemProps) {
  const [item, setItem] = useState<GetItem>({} as GetItem)
  const getItem = async() => {
    let tokenas:string = ""
    const token = sessionStorage.getItem("access_token");
    if(token){
      tokenas = token
    }
    const decodedToken = jwtDecode<DecodedToken>(tokenas);

  const requestOptions = {
    method: "PATCH",
    headers: { "access_token": unescape(tokenas)},
  };
    const url =
      `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/users/${decodedToken.username}/items/${id}`;

    const fetchData = async () => {
      try {
        const response = await fetch(url, requestOptions);
        const responseJSON = await response.json();
        console.log(responseJSON);
        setItem(responseJSON);
      } catch (error) {
        console.log("failed to get data from api", error);
      }
    };
    console.log("Updating data lists");
    fetchData();
  }
  const form = useForm({
    initialValues: {
      name: item.name,
      description: item.description,
      category: item.category,
      body: new FormData()
    },
    validate: {
      name: (value) => value.toString().length >= 4 ? null : 'Daikto pavadinimas turi būti bent 4 simbolių',
      description: (value) => value.length > 10 ? null : 'Daikto aprašymas turi būti bent 10 simbolių',
      category: (value) => value == 'Transportas' ? null : 'Pasirinkite tinkamą kategoriją'
    }
  })
  getItem()
  return (
    <Modal opened={onOpen} onClose={onClose} size="xl">
      <Title>Inventoriaus atnaujinimas</Title>
      <Divider />
      <TextInput
        label="Pavadinimas"
        description="Pavadinimas"
        placeholder={item.name}
        {...form.getInputProps('name')}
      />
      <Select
        label="Kategorija"
        placeholder="Pasirinkti"
        required
        {...form.getInputProps('category')}
        data={[{ value: "Transportas", label: "Transportas" }]}
      />
      <Textarea
        placeholder="Aprašymas"
        label="Aprašymas"
        {...form.getInputProps('description')}
        required
      />
    </Modal>
  );
}
