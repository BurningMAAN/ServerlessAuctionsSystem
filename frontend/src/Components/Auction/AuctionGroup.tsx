import React, { FC, useEffect, useState } from "react";
import { Grid } from "@mantine/core";
import AuctionCard from "./AuctionItem";
import { AuctionProps } from "./AuctionItem";

export default function AuctionGroup() {
  const [auctionsList, setAuctionsList] = useState<AuctionProps[]>([]);

  useEffect(() => {
    const url =
      "http://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctionsList";

    const fetchData = async () => {
      try {
        const response = await fetch(url);
        const responseJSON = await response.json();
        setAuctionsList(responseJSON.auctions);
        console.log(responseJSON.auctions);
      } catch (error) {
        console.log("failed to get data from api", error);
      }
    };
    console.log("Updating data lists");
    fetchData();
  }, []);

  return (
    <>
      {auctionsList.map((auctionItem) => {
        return (
          <Grid.Col span={4}>
            <AuctionCard
              auctionDate={auctionItem.auctionDate}
              buyoutPrice={auctionItem.buyoutPrice}
            ></AuctionCard>
          </Grid.Col>
        );
      })}
    </>
  );
}
