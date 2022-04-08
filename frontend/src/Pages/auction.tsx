import React, { FC, useEffect, useState } from "react";
import AuctionCard from "../Components/Auction/AuctionItem";
import {
  AuctionGroup,
  GetAuctionList,
} from "../Components/Auction/AuctionGroup";
import NavigationBar from "../Components/Skeleton/Navbar";
import HeaderMiddle from "../Components/Skeleton/Header";
import ProgressCircle from "../Components/General/ProgressCircle";
import { Carousel } from "react-bootstrap";
import {
  AppShell,
  Container,
  Grid,
  Menu,
  UnstyledButton,
  Group,
  Stack,
  Select,
  Title,
  Table,
  Divider,
  Badge,
  ScrollArea,
  Button,
  Image,
  Center,
  createStyles,
  Text,
  RingProgress,
} from "@mantine/core";
import { ChevronDown } from "tabler-icons-react";
import AuctionItem from "../Components/Auction/AuctionItem";

interface AuctionProps {
  data: { name: string; email: string; company: string }[];
}

const useStyles = createStyles((theme) => ({
  header: {
    position: "sticky",
    top: 0,
    backgroundColor:
      theme.colorScheme === "dark" ? theme.colors.dark[7] : theme.white,
    transition: "box-shadow 150ms ease",

    "&::after": {
      content: '""',
      position: "absolute",
      left: 0,
      right: 0,
      bottom: 0,
      borderBottom: `1px solid ${
        theme.colorScheme === "dark"
          ? theme.colors.dark[3]
          : theme.colors.gray[2]
      }`,
    },
  },

  scrolled: {
    boxShadow: theme.shadows.sm,
  },
}));

export default function AuctionView({}: AuctionProps) {
  const { classes, cx } = useStyles();
  const [scrolled, setScrolled] = useState(false);

  const [timeLeft, setTimeLeft] = useState(0);

  useEffect(() => {
    if (timeLeft == 0) {
      return;
    }

    if (timeLeft == 1) {
      setTimeLeft(30);
    }
    const intervalId = setInterval(() => {
      setTimeLeft(timeLeft - 1);
    }, 1000);
    return () => clearInterval(intervalId);
  });
  return (
    <AppShell
      padding="md"
      navbar={<NavigationBar></NavigationBar>}
      //   header={<HeaderMiddle></HeaderMiddle>}
      fixed
    >
      <Grid>
        <Grid.Col span={7}>
          <Title>BMW 525D E39</Title>
          <Divider />
          <Carousel style={{height: 450, width: 650, top: 25}} onClick={() => setTimeLeft(30)}>
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
          <Text style={{top: 200}} size="sm">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur nec tempus est, id varius nisi. Donec et aliquet elit. Suspendisse potenti. Donec viverra, eros at mollis finibus, diam neque vehicula dolor, sed aliquet arcu arcu ac turpis. Nulla congue risus id consectetur lacinia. Nulla metus nibh, gravida eget lacus feugiat, tristique ornare massa. Donec vel augue eu justo dapibus porta. Aenean mauris risus, blandit in auctor in, mattis ac ex. Etiam pharetra sollicitudin libero vel porttitor. Aliquam semper pellentesque turpis at ornare. Sed at volutpat tortor. Nunc nec augue vestibulum, auctor purus a, cursus urna. Integer vulputate orci nulla, vitae cursus lacus fringilla id. Aliquam quis ullamcorper risus, non imperdiet dui. Maecenas id hendrerit metus, ac suscipit tellus. Cras gravida pharetra metus ac consequat.
</Text>
        </Grid.Col>
        <Grid.Col span={4}>
          <Center>
            <Title>Absoliutus aukcionas</Title>
          </Center>
          <Divider />
          <Center>
            <h4>Paskutinis statymas</h4>
          </Center>
          <Center>
            <h1>250.00 €</h1>
          </Center>
          <Center>
            <ProgressCircle progressValue={timeLeft}></ProgressCircle>
          </Center>
          <Center>
            <Text>Statymo suma: 25.00 €</Text>
          </Center>
          <Center>
            {(timeLeft !== 0 && <Button color="green">+ 25</Button>) || (
              <Button color="grey" disabled>
                + 25
              </Button>
            )}
          </Center>
          <Table>
            <thead>
              <tr>
                <th>Miestas</th>
                <th>Statymas</th>
                <th>Data</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Kaunas</td>
                <td>250.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
              <tr>
                <td>Kaunas</td>
                <td>225.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
              <tr>
                <td>Kaunas</td>
                <td>200.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
              <tr>
                <td>Kaunas</td>
                <td>175.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
              <tr>
                <td>Kaunas</td>
                <td>150.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
              <tr>
                <td>Kaunas</td>
                <td>125.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
              <tr>
                <td>Kaunas</td>
                <td>100.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
              <tr>
                <td>Kaunas</td>
                <td>75.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
              <tr>
                <td>Kaunas</td>
                <td>50.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
              <tr>
                <td>Kaunas</td>
                <td>25.00 €</td>
                <td>2021-12-12 17:38</td>
              </tr>
            </tbody>
            <caption>Paskutiniai 10 statymų</caption>
          </Table>
        </Grid.Col>
        <Grid.Col span={10}>
            <Divider/>
        </Grid.Col>
      </Grid>
    </AppShell>
  );
}
