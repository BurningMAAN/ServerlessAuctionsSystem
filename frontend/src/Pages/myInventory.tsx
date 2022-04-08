import React, { FC } from "react";
import AuctionCard from "../Components/Auction/AuctionItem";
import {
  AuctionGroup,
  GetAuctionList,
} from "../Components/Auction/AuctionGroup";
import NavigationBar from "../Components/Skeleton/Navbar";
import HeaderMiddle from "../Components/Skeleton/Header";
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
  Button,
} from "@mantine/core";
import { ChevronDown } from "tabler-icons-react";
import AuctionItem from "../Components/Auction/AuctionItem";
interface TitleProps {}

const MyInventory: FC<TitleProps> = ({}) => {
  return (
    <AppShell
      padding="md"
      navbar={<NavigationBar></NavigationBar>}
    //   header={<HeaderMiddle></HeaderMiddle>}
      fixed
    >
    </AppShell>
  );
};

export default MyInventory;
