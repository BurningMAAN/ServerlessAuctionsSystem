import { Button, Center, Title, Text, Grid, Divider } from "@mantine/core";
import { useState, useEffect } from "react";
import { Carousel } from "react-bootstrap";

interface AuctionProps {
  auctionName: string;
  auctionDate: string;
  category: string;
  buyoutPrice: number;
  description: string;
  bidIncrement: number;
  itemId: string;
  id: string;
}

interface AuctionItemProps {
  id: string;
  description: string;
  category: string;
  name: string;
}

interface AuctionInformationDashboardProps {
  auctionID: string;
}

export default function AuctionInformationDashboard({
 auctionID
}: AuctionInformationDashboardProps) {
  const [auction, setAuction] = useState<AuctionProps>({} as AuctionProps);
  const [item, setItem] = useState<AuctionItemProps>({} as AuctionItemProps);

  const auctionURL = `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctions/${auctionID}`;

  const getAuction = async () => {
    try {
      const response = await fetch(auctionURL);
      const responseJSON = await response.json();
      setAuction(responseJSON);
      console.log(responseJSON);
    } catch (error) {
      console.log("failed to get auction data from api", error);
    }
  };

  const itemURL = `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctions/${auction.itemId}`;

  const getItem = async () => {
    try {
      const response = await fetch(itemURL);
      const responseJSON = await response.json();
      setItem(responseJSON);
    } catch (error) {
      console.log("failed to get item data from api", error);
    }
  };
  getAuction();
  getItem();
  
  return (
    <Grid.Col span={7}>
      <Title>{item.name}</Title>
      <Divider />
      <Carousel style={{ height: 450, width: 650, top: 25 }}>
        <Carousel.Item>
          <img
            className="d-block w-100"
            src="https://img.autogidas.lt/10_1_7133435/bmw-530-2000-2004.jpg"
          />
        </Carousel.Item>
        <Carousel.Item>
          <img
            className="d-block w-100"
            src="https://img.autogidas.lt/10_1_7133435/bmw-530-2000-2004.jpg"
          />
        </Carousel.Item>
        <Carousel.Item>
          <img
            className="d-block w-100"
            src="https://img.autogidas.lt/10_1_7133435/bmw-530-2000-2004.jpg"
            alt="Third slide"
          />
        </Carousel.Item>
        <Carousel.Item>
          <img
            className="d-block w-100"
            src="https://img.autogidas.lt/10_1_7133435/bmw-530-2000-2004.jpg"
            alt="Third slide"
          />
        </Carousel.Item>
        <Carousel.Item>
          <img
            className="d-block w-100"
            src="https://img.autogidas.lt/10_1_7133435/bmw-530-2000-2004.jpg"
            alt="Third slide"
          />
        </Carousel.Item>
        <Carousel.Item>
          <img
            className="d-block w-100"
            src="https://img.autogidas.lt/10_1_7133435/bmw-530-2000-2004.jpg"
            alt="Third slide"
          />
        </Carousel.Item>
      </Carousel>
      <Title py={50}>Informacija</Title>
      <Text style={{ top: 200 }} size="sm">
        {item.description}
      </Text>
    </Grid.Col>
  );
}
