import { Button, Center, Title, Text, Grid, Divider, Container } from "@mantine/core";
import { useState, useEffect } from "react";
import { Carousel } from "react-bootstrap";

interface AuctionInformationDashboardProps {
  description: string;
  name: string;
  photoURLs: string[];
}

export default function AuctionInformationDashboard({
  description,
  name,
  photoURLs,
}: AuctionInformationDashboardProps) {
  return (
    <Grid.Col span={8}>
      <Title>{name}</Title>
      <Divider />
      <Carousel variant="dark" style={{ maxHeight: 450, maxWidth: 650, top: 25, minHeight: 450, minWidth: 650 }}>
      {photoURLs?.map((photoURL) => {
        return (
          <Carousel.Item>
          <img
            style={{maxHeight: 450, maxWidth: 650, minHeight: 450, minWidth: 650}}
            className="d-block w-50"
            src={`${process.env.REACT_APP_S3_URL}/${photoURL}`}
          />
        </Carousel.Item>
        )
      })}
      </Carousel>
      <Container>
      <Title py={50}>Informacija</Title>
      <Text style={{ top: 200 }} size="sm">
        {description}
      </Text>
      </Container>
    </Grid.Col>
  );
}
