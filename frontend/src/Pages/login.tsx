import React from 'react';
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
} from '@mantine/core';

const useStyles = createStyles((theme) => ({
  wrapper: {
    minHeight: 900,
    backgroundSize: 'cover',
    backgroundImage:
      'url(https://images.pexels.com/photos/3183197/pexels-photo-3183197.jpeg)',
  },

  form: {
    borderRight: `1px solid ${
      theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.colors.gray[3]
    }`,
    minHeight: 900,
    maxWidth: 450,
    paddingTop: 80,

    [`@media (max-width: ${theme.breakpoints.sm}px)`]: {
      maxWidth: '100%',
    },
  },

  title: {
    color: theme.colorScheme === 'dark' ? theme.white : theme.black,
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
  },

  logo: {
    color: theme.colorScheme === 'dark' ? theme.white : theme.black,
    width: 120,
    display: 'block',
    marginLeft: 'auto',
    marginRight: 'auto',
  },
}));

export default function AuthenticationImage() {
  const { classes } = useStyles();
  return (
    <div className={classes.wrapper}>
      <Paper className={classes.form} radius={0} p={30}>
        <Title order={2} className={classes.title} align="center" mt="md" mb={50}>
          Prisijungimas
        </Title>

        <TextInput label="Vartotojo vardas" placeholder="Vartotojas" size="md" />
        <PasswordInput label="Slaptažodis" placeholder="Slaptažodis" mt="md" size="md" />
        {/* <Checkbox label="Keep me logged in" mt="xl" size="md" /> */}
        <Button fullWidth mt="xl" size="md">
          Prisijungti
        </Button>

        <Text align="center" mt="md">
          Neturite paskyros?{' '}
          <Anchor<'a'> href="#" weight={700} onClick={(event) => event.preventDefault()}>
            Registruotis
          </Anchor>
        </Text>
      </Paper>
    </div>
  );
}
