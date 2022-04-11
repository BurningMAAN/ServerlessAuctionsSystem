import { useState } from "react";
import { createStyles, Navbar, Group, Code, Image } from "@mantine/core";
import { Logout, Menu, Login } from "tabler-icons-react";
import sebas from "../../shared/sebas.png";
import { Redirect } from 'react-router-dom';

const getToken = () => {
  const tokenString = sessionStorage.getItem("access_token");
  console.log(tokenString);
  return tokenString;
};

const useStyles = createStyles((theme, _params, getRef) => {
  const icon = getRef("icon");
  return {
    header: {
      paddingBottom: theme.spacing.md,
      marginBottom: theme.spacing.md * 1.5,
      borderBottom: `1px solid ${
        theme.colorScheme === "dark"
          ? theme.colors.dark[4]
          : theme.colors.gray[2]
      }`,
    },

    footer: {
      paddingTop: theme.spacing.md,
      marginTop: theme.spacing.md,
      borderTop: `1px solid ${
        theme.colorScheme === "dark"
          ? theme.colors.dark[4]
          : theme.colors.gray[2]
      }`,
    },

    link: {
      ...theme.fn.focusStyles(),
      display: "flex",
      alignItems: "center",
      textDecoration: "none",
      fontSize: theme.fontSizes.sm,
      color:
        theme.colorScheme === "dark"
          ? theme.colors.dark[1]
          : theme.colors.gray[7],
      padding: `${theme.spacing.xs}px ${theme.spacing.sm}px`,
      borderRadius: theme.radius.sm,
      fontWeight: 500,

      "&:hover": {
        backgroundColor:
          theme.colorScheme === "dark"
            ? theme.colors.dark[6]
            : theme.colors.gray[0],
        color: theme.colorScheme === "dark" ? theme.white : theme.black,

        [`& .${icon}`]: {
          color: theme.colorScheme === "dark" ? theme.white : theme.black,
        },
      },
    },

    linkIcon: {
      ref: icon,
      color:
        theme.colorScheme === "dark"
          ? theme.colors.dark[2]
          : theme.colors.gray[6],
      marginRight: theme.spacing.sm,
    },

    linkActive: {
      "&, &:hover": {
        backgroundColor:
          theme.colorScheme === "dark"
            ? theme.fn.rgba(theme.colors[theme.primaryColor][8], 0.25)
            : theme.colors[theme.primaryColor][0],
        color:
          theme.colorScheme === "dark"
            ? theme.white
            : theme.colors[theme.primaryColor][7],
        [`& .${icon}`]: {
          color:
            theme.colors[theme.primaryColor][
              theme.colorScheme === "dark" ? 5 : 7
            ],
        },
      },
    },
  };
});

export default function NavigationBar() {
  const { classes, cx } = useStyles();
  const token = getToken();

  return (
    <Navbar height={700} width={{ sm: 300 }} p="md">
      <Image src={sebas}></Image>
      <Navbar.Section grow>
        <Group className={classes.header} position="apart"></Group>
        <a className={cx(classes.link)} href="/" key="Visi aukcionai"> <Menu className={classes.linkIcon} />
        <span>Visi aukcionai</span></a>
      {token && (
        <>
        <a className={cx(classes.link)} href="/myAuctions" key="Mano aukcionai"> <Menu className={classes.linkIcon} />
        <span>Mano aukcionai</span></a>
        <a className={cx(classes.link)} href="/myInventory" key="Mano inventorius"> <Menu className={classes.linkIcon} />
        <span>Mano inventorius</span></a>
        </>
      )}
      </Navbar.Section>

      <Navbar.Section className={classes.footer}>
        {!token && (
          <a href="/login" className={classes.link}>
            <Login className={classes.linkIcon} />
            <span>Prisijungti</span>
          </a>
        )}
        {token && (
          <a
            href="#"
            className={classes.link}
            onClick={(event) => {
              event.preventDefault()
              sessionStorage.removeItem('access_token')
              window.location.reload();
            }}
          >
            <Logout className={classes.linkIcon} />
            <span>Atsijungti</span>
          </a>
        )}
      </Navbar.Section>
    </Navbar>
  );
}
