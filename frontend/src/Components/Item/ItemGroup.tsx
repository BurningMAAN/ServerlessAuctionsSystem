import React, { FC, useEffect, useState } from "react";
import { Grid } from "@mantine/core";
import ItemCard from "./ItemCard";

export interface ItemList {
  items: [
    {
      id: string;
      description: string;
      category: string;
      name: string;
    }
  ]
}

export default function AuctionGroup() {
  const [itemsList, setItemsList] = useState<ItemList>({} as ItemList);

  useEffect(() => {
    const url =
      "https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/itemsList";

    const fetchData = async () => {
      try {
        const response = await fetch(url);
        const responseJSON = await response.json();
        console.log(responseJSON);
        setItemsList(responseJSON);
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
      itemsList.items?.map((item) => {
        return (
          <Grid.Col span={4}>
            <ItemCard
              name={item.name}
              id={item.id}
              category={item.category}
              description={item.description}
            ></ItemCard>
          </Grid.Col>
        );
      })}
    </>
  );
}
