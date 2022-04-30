import React, { FC, useEffect, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import { AppShell, Tabs, Grid, Title, Container, Button } from "@mantine/core";
import AuctionCard from "../Components/Auction/AuctionItem";
import jwtDecode, { JwtPayload } from "jwt-decode";
import { Photo, MessageCircle } from "tabler-icons-react";
interface TitleProps {}

export interface AuctionList {
  auctions: [
    {
      id: string;
      auctionDate: string;
      buyoutPrice: number;
      auctionType: string;
      bidIncrement: number;
      creatorId: string;
      isFinished: boolean;
      description: string;
      item: {
        category: string;
        name: string;
      };
    }
  ];
}

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

const UserDashboard: FC<TitleProps> = ({}) => {
  const token = getToken();
  let decodedToken: DecodedToken = {} as DecodedToken;
  if (token) {
    decodedToken = jwtDecode<DecodedToken>(token);
  }
  const [auctionsList, setAuctionsList] = useState<AuctionList>(
    {} as AuctionList
  );
  useEffect(() => {
    const url =
      "https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctionsList";

    const fetchData = async () => {
      try {
        const response = await fetch(url);
        const responseJSON = await response.json();
        console.log(responseJSON);
        setAuctionsList(responseJSON);
      } catch (error) {
        console.log("failed to get data from api", error);
      }
    };
    console.log("Updating data lists");
    fetchData();
  }, []);
  return (
    <AppShell padding="md" navbar={<NavigationBar></NavigationBar>} fixed>
      <Container>
        <Grid>
          <Grid.Col span={4}>
              <Button>Atnaujinti vartotojo informaciją</Button>
          </Grid.Col>
          <Grid.Col span={4}>
              <Button>Generuoti aukcionų ataskaitą</Button>
              </Grid.Col>
        </Grid>
      </Container>
    </AppShell>
  );
};

export default UserDashboard;
