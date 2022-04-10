import { useEffect, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import ProgressCircle from "../Components/General/ProgressCircle";
import { Carousel } from "react-bootstrap";
import {
  AppShell,
  Grid,
  Title,
  Table,
  Divider,
  Button,
  Center,
  createStyles,
  Text,
} from "@mantine/core";
import axios from 'axios';
import AuctionBiddingDashboard from "../Components/Auction/AuctionBidding/AuctionBidding";
import AuctionInformationDashboard from "../Components/Auction/AuctionBidding/AuctionInformationBlock";

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

  const [timeLeft, setTimeLeft] = useState(30);

  useEffect(() => {
    if (timeLeft == 0) {
      console.log('aukcionas baigesi')
      return;
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
      fixed
    >
      <Grid>
        <AuctionInformationDashboard auctionID="example"></AuctionInformationDashboard>
        <AuctionBiddingDashboard
        auctionType="absolute"
        timeLeft={timeLeft}
        setTimeLeft={setTimeLeft}
        currentMaxBid={30}
        bidIncrement={15}
        ></AuctionBiddingDashboard>
        <Grid.Col span={10}>
          <Divider />
        </Grid.Col>
      </Grid>
    </AppShell>
  );
}
