import React, { FC, useEffect, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import { AppShell, Tabs, Grid, Title } from "@mantine/core";
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

const MyAuctions: FC<TitleProps> = ({}) => {
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
      <Tabs>
        <Tabs.Tab label="Aktyvūs Aukcionai" icon={<Photo size={14} />}>
          <Grid>
            {auctionsList?.auctions?.map((auctionItem) => {
              return (
                <>
                  {auctionItem.creatorId == decodedToken.username && (
                    <Grid.Col span={4}>
                      <AuctionCard
                        isFinished={auctionItem.isFinished}
                        auctionID={auctionItem.id}
                        auctionDate={auctionItem.auctionDate}
                        buyoutPrice={auctionItem.buyoutPrice}
                        auctionName={auctionItem.item.name}
                        category={auctionItem.item.category}
                        description={auctionItem.description}
                        bidIncrement={auctionItem.bidIncrement}
                      ></AuctionCard>
                    </Grid.Col>
                  )}
                </>
              );
            })}
          </Grid>
        </Tabs.Tab>
        <Tabs.Tab label="Baigęsi Aukcionai" icon={<MessageCircle size={14} />}>
          <Title>Bye bye</Title>
        </Tabs.Tab>
      </Tabs>
    </AppShell>
  );
};

export default MyAuctions;
