import { Button, Center, Title, Text, Grid, Divider } from "@mantine/core";
import { useState, useEffect } from "react";
import { Carousel } from "react-bootstrap";

interface AuctionInformationDashboardProps {
  description: string;
  name: string;
}

export default function AuctionInformationDashboard({
  description,
  name
}: AuctionInformationDashboardProps) {
  return (
    <Grid.Col span={8}>
      <Title>{name}</Title>
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
        {description}
      </Text>
    </Grid.Col>
  );
}
