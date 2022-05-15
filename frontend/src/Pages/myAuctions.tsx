import React, { FC, useEffect, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import { AppShell, Tabs, Grid, Title } from "@mantine/core";
import AuctionCard from "../Components/Auction/AuctionItem";
import jwtDecode, { JwtPayload } from "jwt-decode";
import { Photo, MessageCircle } from "tabler-icons-react";
import MyAuctionCard from "../Components/Auction/AuctionCard"
interface TitleProps {}

export interface AuctionList {
  auctions: [
    {
      id: string;
      auctionDate: string;
      auctionType: string;
      bidIncrement: number;
      creatorId: string;
      isFinished: boolean;
      stage: string;
      photoURL: string;
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
  const [activeAuctionsList, setActiveAuctionsList] = useState<AuctionList>(
    {} as AuctionList
  );
  const [endedAuctionsList, setEndedAuctionsList] = useState<AuctionList>(
    {} as AuctionList
  );
  const [wonAuctionsList, setWonAuctionsList] = useState<AuctionList>(
    {} as AuctionList
  );
  useEffect(() => {
    const url =
      `${process.env.REACT_APP_API_URL}auctions/search`;

    const getActive = async () => {
      try {
        const requestOptions = {
          method: "POST",
          body: JSON.stringify({stage: 'STAGE_ACCEPTING_BIDS'})
        };
        const response = await fetch(url, requestOptions);
        const responseJSON = await response.json();
        console.log(responseJSON);
        setActiveAuctionsList(responseJSON);
      } catch (error) {
        console.log("failed to get data from api", error);
      }
    };
    const getEnded = async () => {
      try {
        const requestOptions = {
          method: "POST",
          body: JSON.stringify({stage: 'STAGE_AUCTION_FINISHED'})
        };
        const response = await fetch(url, requestOptions);
        const responseJSON = await response.json();
        console.log(responseJSON);
        setEndedAuctionsList(responseJSON);
      } catch (error) {
        console.log("failed to get data from api", error);
      }
    };
    const getWon = async () => {
      try {
        const requestOptions = {
          method: "POST",
          body: JSON.stringify({winnerName: decodedToken.username})
        };
        const response = await fetch(url, requestOptions);
        const responseJSON = await response.json();
        console.log(responseJSON);
        setWonAuctionsList(responseJSON);
      } catch (error) {
        console.log("failed to get data from api", error);
      }
    };
    getActive();
    getWon();
    getEnded();
  }, []);
  return (
    <AppShell padding="md" navbar={<NavigationBar></NavigationBar>} fixed>
      <Tabs>
        <Tabs.Tab label="Aktyvūs Aukcionai" icon={<Photo size={14} />}>
          <Grid>
            {activeAuctionsList?.auctions?.map((auctionItem) => {
              return (
                <>
                  {auctionItem.creatorId == decodedToken.username && (
                    <Grid.Col span={4}>
                      <MyAuctionCard
                      stage={auctionItem.stage}
                        auctionID={auctionItem.id}
                        auctionDate={auctionItem.auctionDate}
                        auctionName={auctionItem.item.name}
                        category={auctionItem.item.category}
                        bidIncrement={auctionItem.bidIncrement}
                        photoURL={auctionItem.photoURL}
                      ></MyAuctionCard>
                    </Grid.Col>
                  )}
                </>
              );
            })}
          </Grid>
        </Tabs.Tab>
        <Tabs.Tab label="Baigti Aukcionai" icon={<MessageCircle size={14} />}>
        {endedAuctionsList?.auctions?.map((auctionItem) => {
              return (
                <>
                  {auctionItem.creatorId == decodedToken.username && (
                    <Grid.Col span={4}>
                      <MyAuctionCard
                      stage={auctionItem.stage}
                        auctionID={auctionItem.id}
                        auctionDate={auctionItem.auctionDate}
                        auctionName={auctionItem.item.name}
                        category={auctionItem.item.category}
                        bidIncrement={auctionItem.bidIncrement}
                        photoURL={auctionItem.photoURL}
                      ></MyAuctionCard>
                    </Grid.Col>
                  )}
                </>
              );
            })}
        </Tabs.Tab>
        <Tabs.Tab label="Laimėti Aukcionai" icon={<MessageCircle size={14} />}>
        {wonAuctionsList?.auctions?.map((auctionItem) => {
              return (
                <>
                  {auctionItem.creatorId == decodedToken.username && (
                    <Grid.Col span={4}>
                      <MyAuctionCard
                      stage={auctionItem.stage}
                        auctionID={auctionItem.id}
                        auctionDate={auctionItem.auctionDate}
                        auctionName={auctionItem.item.name}
                        category={auctionItem.item.category}
                        bidIncrement={auctionItem.bidIncrement}
                        photoURL={auctionItem.photoURL}
                      ></MyAuctionCard>
                    </Grid.Col>
                  )}
                </>
              );
            })}
        </Tabs.Tab>
      </Tabs>
    </AppShell>
  );
};

export default MyAuctions;
