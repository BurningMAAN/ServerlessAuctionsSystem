import React, { FC, useEffect, useState } from "react";
import { Grid } from "@mantine/core";
import AuctionCard from "./AuctionItem";
import { AuctionProps } from "./AuctionItem";

export interface AuctionList {
  auctions: [
    {
      id: string;
      auctionDate: string;
      buyoutPrice: number;
      auctionType: string;
      bidIncrement: number;
      isFinished: boolean;
      description: string;
      photoURL: string;
      item: {
        category: string;
        name: string;
      };
    }
  ];
}

export default function AuctionGroup() {
  const [auctionsList, setAuctionsList] = useState<AuctionList>(
    {} as AuctionList
  );

  useEffect(() => {
    const url =
    `${process.env.REACT_APP_API_URL}auctionsList`;

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
    <>
      {auctionsList?.auctions?.map((auctionItem) => {
        return (
          <Grid.Col span={4}>
            <AuctionCard
              photoURL={auctionItem.photoURL}
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
        );
      })}
    </>
  );
}
