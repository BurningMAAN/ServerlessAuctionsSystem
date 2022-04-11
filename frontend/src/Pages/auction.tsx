import { useEffect, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import ProgressCircle from "../Components/General/ProgressCircle";
import { Carousel } from "react-bootstrap";
import { useParams } from "react-router-dom";
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
import axios from "axios";
import AuctionBiddingDashboard from "../Components/Auction/AuctionBidding/AuctionBidding";
import AuctionInformationDashboard from "../Components/Auction/AuctionBidding/AuctionInformationBlock";

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

interface AuctionViewProps{}
export default function AuctionView({}: AuctionViewProps) {
  const { auctionID } = useParams<{ auctionID: string }>();
  const { classes, cx } = useStyles();
  const [scrolled, setScrolled] = useState(false);
  console.log(auctionID)

  return (
    <AppShell padding="md" navbar={<NavigationBar></NavigationBar>} fixed>
      <Grid>
        <AuctionInformationDashboard
          auctionID={auctionID}
        ></AuctionInformationDashboard>
        <AuctionBiddingDashboard
          auctionType="absolute"
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
