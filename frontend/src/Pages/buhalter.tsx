import React, { FC, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import AuctionCreateWizard from '../Components/Auction/CreateAuction/AuctionWizard';
import {
  AppShell,
  Container,
  Grid,
  Group,
  Select,
  Title,
  Button,
} from "@mantine/core";
import AuctionGroup from "../Components/Auction/AuctionGroup";
import { Redirect } from 'react-router-dom';
import jwtDecode, { JwtPayload } from "jwt-decode";
interface TitleProps {}

interface DecodedToken {
    role: string;
  }

const getToken = () => {
    let tokenas = "";
    const tokenString = sessionStorage.getItem("access_token");
    if (tokenString) {
      tokenas = tokenString;
    }
    return tokenas;
  };

const BugalterDashboard: FC<TitleProps> = ({}) => {
   const token = getToken();
  return (
    <AppShell
      padding="md"
      navbar={<NavigationBar></NavigationBar>}
      style={{
          overflow: 'hidden'
      }}
    //   header={<HeaderMiddle></HeaderMiddle>}
      fixed
    >
        <div>
        <Group spacing="sm">
           <Button color="teal" px={50} style={{top: 12, left: 400}} onClick={() => {
             if(token){

             }
           }}>
            Generuoti ataskaitas
            </Button>
        </Group>
        </div>
      <hr />
      <Container size="lg" style={{overflowY: 'scroll',overflowX: 'hidden', height: 800}}>
      </Container>
    </AppShell>
  );
};

export default BugalterDashboard;
