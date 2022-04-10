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
      description: string;
      item:{
        category: string;
        name: string;
      }
    }
  ]
}

export default function AuctionGroup() {
  const [auctionsList, setAuctionsList] = useState<AuctionList>({} as AuctionList);

  useEffect(() => {
    const url =
      "http://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctionsList";

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
      {
      auctionsList?.auctions?.map((auctionItem) => {
        return (
          <Grid.Col span={4}>
            <AuctionCard
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
