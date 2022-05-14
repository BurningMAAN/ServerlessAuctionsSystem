import React, { FC, useEffect, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import { showNotification } from '@mantine/notifications';
import {
  AppShell,
  Tabs,
  Grid,
  Title,
  Container,
  Button,
  Divider,
  TextInput,
  PasswordInput,
  NumberInput,
  Text,
} from "@mantine/core";
import AuctionCard from "../Components/Auction/AuctionItem";
import jwtDecode, { JwtPayload } from "jwt-decode";
import { Photo, MessageCircle } from "tabler-icons-react";
import UpdateUserModal from "../Components/User/UpdateUserModal";
import { useForm } from "@mantine/form";

interface DecodedToken {
  username: string;
}

const getToken = () => {
  let tokenas = "";
  const tokenString = sessionStorage.getItem("access_token");
  if (tokenString) {
    tokenas = tokenString;
  }
  return tokenas;
};

export default function UserDashboard() {
  const token = getToken();
  let decodedToken: DecodedToken = {} as DecodedToken;
  if (token) {
    decodedToken = jwtDecode<DecodedToken>(token);
  }

  const updateUser = async() => {
    let tokenas:string = ""
    const token = sessionStorage.getItem("access_token");
    if(token){
      tokenas = token
    }
    const decodedToken = jwtDecode<DecodedToken>(tokenas);


  const requestOptions = {
    method: "PATCH",
    headers: { "access_token": unescape(tokenas)},
    body: JSON.stringify(form)
  };
    let url = `${process.env.REACT_APP_API_URL}/user`;
    const fetchData = async () => {
      try {
        const response = await fetch(url, requestOptions);
        const responseJSON = await response.json();
        console.log(responseJSON);
      } catch (error) {
        console.log("failed to get data from api", error);
      }
    };
    fetchData();
  }

  const form = useForm({
    initialValues: {
      creditBalance: 0.0,
      userId: decodedToken.username
    },
    validate: {
      creditBalance: (value) => value > 0 ? null : 'Kreditų papildymas turi būti didesnis nei nulis',
      // name: (value) => value.toString().length >= 4 ? null : 'Daikto pavadinimas turi būti bent 4 simbolių',
      // description: (value) => value.length > 10 ? null : 'Daikto aprašymas turi būti bent 10 simbolių',
      // category: (value) => value == 'Transportas' || 'Baldai' || 'Elektronika' || 'Automobilių detalės' || 'Drabužiai' || 'Paveikslai' ? null : 'Pasirinkite tinkamą kategoriją'
    }
  })

  const [openEditModal, setOpenEditModal] = useState(false);
  return (
    <AppShell padding="md" navbar={<NavigationBar></NavigationBar>} fixed>
      <Container>
        <Title>Vartotojo informacija</Title>
        <Divider />
        <Grid>
          <Grid.Col span={4}>
            <Title order={5}>Vartotojo duomenys</Title>
            <br />
            <TextInput
              label="Prisijungimo slapyvardis"
              value="testing"
              disabled
              style={{ width: "200px" }}
            ></TextInput>
            <PasswordInput label="Slaptažodis" style={{ width: "200px" }} />
            <PasswordInput
              label="Pakartokite slaptažodį"
              style={{ width: "200px" }}
            />
            <TextInput
            label="Elektroninis paštas"
            placeholder="Elektroninis paštas"
            value="example@gmail.com"
            style={{width: "200px"}}
          />
          <br/>
            <Button color="green">Atnaujinti</Button>
          </Grid.Col>
          <Grid.Col span={4}>
            <Title order={5}>Mokėjimai</Title>
            <NumberInput label="Kreditų skaičius"
            {...form.getInputProps('creditBalance')}></NumberInput>
            <br/>
            <Button color="green" 
            onClick={() => updateUser()}>Papildyti</Button>
          </Grid.Col>
        </Grid>
        <UpdateUserModal
          onOpen={openEditModal}
          onClose={() => setOpenEditModal(false)}
          userID="test"
        ></UpdateUserModal>
      </Container>
    </AppShell>
  );
};

