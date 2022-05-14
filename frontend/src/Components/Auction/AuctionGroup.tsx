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
      stage: string;
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
       let auctionDateParsed = new Date(auctionItem.auctionDate)
       let formatted = formatDate(auctionDateParsed)
        return (
          <Grid.Col span={4}>
            <AuctionCard
            stage={auctionItem.stage}
              photoURL={auctionItem.photoURL}
              auctionID={auctionItem.id}
              auctionDate={formatted}
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

function padTo2Digits(num: number) {
  return num.toString().padStart(2, '0');
}

function formatDate(date: Date) {
  return (
    [
      date.getFullYear(),
      padTo2Digits(date.getMonth() + 1),
      padTo2Digits(date.getDate()),
    ].join('-') +
    ' ' +
    [
      padTo2Digits(date.getHours()),
      padTo2Digits(date.getMinutes()),
    ].join(':')
  );
}
