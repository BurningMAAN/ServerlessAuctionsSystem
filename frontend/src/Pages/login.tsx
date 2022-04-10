import React, { useState } from "react";
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

export default function AuthenticationImage() {
  const { classes } = useStyles();
  const [clickedRegister, setClickedRegister] = useState(false);
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
       response.json()
       console.log(response)
       
      });
    } catch (error) {
      console.log("failed to create user", error);
    }

    setClickedRegister(false);
  };

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
          />
          <PasswordInput
            label="Slaptažodis"
            placeholder="Slaptažodis"
            mt="md"
            size="md"
          />
          {/* <Checkbox label="Keep me logged in" mt="xl" size="md" /> */}
          <Button fullWidth mt="xl" size="md">
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
          <Button fullWidth mt="xl" size="md" onClick={() => registerUser(userInformation)}>
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
    </div>
  );
}
