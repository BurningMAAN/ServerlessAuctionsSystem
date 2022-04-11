import React, { useState} from "react";
import {
  Paper,
  createStyles,
  TextInput,
  PasswordInput,
  Checkbox,
  Button,
  Title,
  Text,
  Anchor,
} from "@mantine/core";
import { Redirect } from 'react-router-dom';
import { At } from "tabler-icons-react";

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

  const registerUser = async (user: CreateUserRequest) => {
    const url =
      "https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/user";

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    };

    try {
      fetch(url, requestOptions).then((response) => {
        response.json();
        console.log(response);
      });
    } catch (error) {
      console.log("failed to create user", error);
    }

    setClickedRegister(false);
  };

  const authorizeUser = async (user: AuthorizeUserRequest) => {
    const url =
      "https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/authorize";

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    };

    try {
      const response = await fetch(url, requestOptions);
      const responseJSON = await response.json();
      console.log(responseJSON);
      setToken(responseJSON);
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
          <Title
            order={2}
            className={classes.title}
            align="center"
            mt="md"
            mb={50}
          >
            Registracija
          </Title>

          <TextInput
            label="Vartotojo slapyvardis"
            placeholder="Vartotojas"
            size="md"
            onChange={(event) =>
              setUserInformation({
                username: event.currentTarget.value,
                password: userInformation.password,
                email: userInformation.email,
              } as CreateUserRequest)
            }
          />
          <PasswordInput
            label="Slaptažodis"
            placeholder="Slaptažodis"
            mt="md"
            size="md"
            onChange={(event) =>
              setUserInformation({
                username: userInformation.username,
                password: event.currentTarget.value,
                email: userInformation.email,
              } as CreateUserRequest)
            }
          />
          <PasswordInput
            label="Patvirtinti slaptažodį"
            placeholder="Patvirtinti slaptažodį"
            mt="md"
            size="md"
          />
          <TextInput
            label="Elektroninis paštas"
            placeholder="Elektroninis paštas"
            onChange={(event) =>
              setUserInformation({
                username: userInformation.username,
                password: userInformation.password,
                email: event.currentTarget.value,
              } as CreateUserRequest)
            }
            icon={<At size={14} />}
          />
          <Button
            fullWidth
            mt="xl"
            size="md"
            onClick={() => registerUser(userInformation)}
          >
            Registruotis
          </Button>

          <Text align="center" mt="md">
            Turite paskyrą?{" "}
            <Anchor<"a">
              href="#"
              weight={700}
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
