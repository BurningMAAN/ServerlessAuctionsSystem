import React, { useState} from "react";
import {
  Paper,
  createStyles,
  TextInput,
  PasswordInput,
  Checkbox,
  Button,
  Title,
  Center,
  Box,
  Text,
  Group,
  Anchor,
  Notification,
} from "@mantine/core";
import { Redirect } from 'react-router-dom';
import {X} from "tabler-icons-react";
import { showNotification } from '@mantine/notifications';
import { useForm } from '@mantine/form';

const useStyles = createStyles((theme) => ({
  wrapper: {
    minHeight: 900,
    backgroundSize: "cover",
    backgroundImage:
      "url(https://images.pexels.com/photos/3183197/pexels-photo-3183197.jpeg)",
  },

  form: {
    borderRight: `1px solid ${
      theme.colorScheme === "dark" ? theme.colors.dark[7] : theme.colors.gray[3]
    }`,
    minHeight: 900,
    maxWidth: 450,
    paddingTop: 80,

    [`@media (max-width: ${theme.breakpoints.sm}px)`]: {
      maxWidth: "100%",
    },
  },

  title: {
    color: theme.colorScheme === "dark" ? theme.white : theme.black,
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
  },

  logo: {
    color: theme.colorScheme === "dark" ? theme.white : theme.black,
    width: 120,
    display: "block",
    marginLeft: "auto",
    marginRight: "auto",
  },
}));

interface CreateUserRequest {
  username: string;
  password: string;
  email: string;
}

interface AuthorizeUserRequest {
  username: string;
  password: string;
}

interface AuthorizerUserResponse {
  access_token: string;
}

export default function AuthenticationImage() {
  const [loginNotification, setLoginNotification] = useState(false)
  const { classes } = useStyles();
  const [token, setToken] = useState<AuthorizerUserResponse>(
    {} as AuthorizerUserResponse
  );
  const [clickedRegister, setClickedRegister] = useState(false);
  const [loginInformation, setLoginInformation] =
    useState<AuthorizeUserRequest>({} as AuthorizeUserRequest);
  const [userInformation, setUserInformation] = useState<CreateUserRequest>(
    {} as CreateUserRequest
  );

  const form = useForm({
    initialValues: {
      username: '',
      password: '',
      email: '',
    },
    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Neteisingas el.pašto formatas'),
      username: (value) => value.length >= 6 ? null : 'Vartotojo vardas turi būti bent 6 simbolių',
      password: (value) => value.length > 6 ? null : 'Vartotojo slapyvardis turi būti ilgesnis nei 6 simboliai'
    }
  })

  const registerUser = async (user: CreateUserRequest) => {
    const url =
      `${process.env.REACT_APP_API_URL}user`;

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    };

    try {
      fetch(url, requestOptions).then((response) => {
        response.json();
        if(response.status == 201){
          setClickedRegister(false);
        }

      });
    } catch (error) {
      console.log("failed to create user", error);
      setLoginNotification(true)
    }
  };

  const authorizeUser = async (user: AuthorizeUserRequest) => {
    const url =
      `${process.env.REACT_APP_API_URL}authorize`;

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    };

    try {
      const response = await fetch(url, requestOptions);
      const responseJSON = await response.json();
      if(response.status == 200){
        setToken(responseJSON);
      } else{
        showNotification({
          title: 'Autorizacija',
          message: 'Nepavyko autorizuoti vartotojo. Pasitikrinkite, ar įvesti duomenys yra teisingi',
          color: 'red',
          icon: <X/>,
        })
      }
    } catch (error) {
      console.log("failed to get data from api", error);
    }
  };

  sessionStorage.setItem('access_token', JSON.stringify(token.access_token));
  return (
    <div className={classes.wrapper}>
      {!clickedRegister && (
        <Paper className={classes.form} radius={0} p={30}>
          <Title
            order={2}
            className={classes.title}
            align="center"
            mt="md"
            mb={50}
          >
            Prisijungimas
          </Title>

          <TextInput
            label="Vartotojo vardas"
            placeholder="Vartotojas"
            size="md"
            onChange={(event) => {
              setLoginInformation({
                username: event.currentTarget.value,
                password: loginInformation.password,
              } as AuthorizeUserRequest);
            }}
          />
          <PasswordInput
            label="Slaptažodis"
            placeholder="Slaptažodis"
            mt="md"
            size="md"
            onChange={(event) => {
              setLoginInformation({
                username: loginInformation.username,
                password: event.currentTarget.value,
              } as AuthorizeUserRequest);
            }}
          />
          {/* <Checkbox label="Keep me logged in" mt="xl" size="md" /> */}
          <Button
            fullWidth
            mt="xl"
            size="md"
            onClick={() => authorizeUser(loginInformation)}
          >
            Prisijungti
          </Button>

          <Text align="center" mt="md">
            Neturite paskyros?{" "}
            <Anchor<"a">
              href="#"
              weight={700}
              onClick={(event) => {
                event.preventDefault();
                setClickedRegister(true);
              }}
            >
              Registruotis
            </Anchor>
          </Text>
        </Paper>
      )}
      {clickedRegister && (
        <Paper className={classes.form} radius={0} p={30}>
        <Box sx={{ maxWidth: 300 }} mx="auto">
        <form onSubmit={form.onSubmit((values) => {
          registerUser(form.values)
        })}>
          <TextInput 
          required
          label="Vartotojo vardas"
          placeholder="Vartotojo vardas"
          {...form.getInputProps('username')}
          />
          <PasswordInput
          required
          label="Slaptažodis"
          placeholder="Slaptažodis"
          {...form.getInputProps('password')} />
          <TextInput
            required
            label="Elektroninis paštas"
            placeholder="Elektroninis paštas"
            {...form.getInputProps('email')}
          />
  
          <Group position="right" mt="md">
            <Center>
            <Button type="submit">Registruotis</Button>
            </Center>
          </Group>
        </form>
      </Box>
      <Text align="center" mt="md">
            Turite paskyrą?{" "}
             <Anchor<"a">
               href="#"
             weight={700}
              type="submit"
              onClick={(event) => {
                event.preventDefault();
                 setClickedRegister(false);
               }}
             >
               Prisijungti
             </Anchor>
           </Text>
      </Paper>
      )}
      {token.access_token && <Redirect to="/"></Redirect>}
    </div>
  );
}
