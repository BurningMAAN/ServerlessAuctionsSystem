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
      photoURLs: string[];
    }
  ]
}

export default function AuctionGroup() {
  const [itemsList, setItemsList] = useState<ItemList>({} as ItemList);

  
  useEffect(() => {
    let tokenas:string = ""
    const token = sessionStorage.getItem("access_token");
    if(token){
      tokenas = token
    }

  const requestOptions = {
    method: "GET",
    headers: { "access_token": unescape(tokenas)},
  };
    const url =
      `${process.env.REACT_APP_API_URL}user/items`;

    const fetchData = async () => {
      try {
        const response = await fetch(url, requestOptions);
        const responseJSON = await response.json();
        setItemsList(responseJSON);
      } catch (error) {
        console.log("failed to get data from api", error);
      }
    };
    fetchData();
  }, []);

  return (
    <>
      {
      itemsList.items?.map((item) => {
        return (
          <Grid.Col span={4}>
            <ItemCard
              photoURLs={item.photoURLs}
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
