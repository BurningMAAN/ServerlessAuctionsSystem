import NavigationBar from "../Components/Skeleton/Navbar";
import { useParams } from "react-router-dom";
import { AppShell, Grid, Divider } from "@mantine/core";
import { useState, useEffect } from "react";
import AuctionBiddingDashboard from "../Components/Auction/AuctionBidding/AuctionBidding";
import AuctionInformationDashboard from "../Components/Auction/AuctionBidding/AuctionInformationBlock";

interface AuctionProps {
  auctionName: string;
  auctionDate: string;
  category: string;
  description: string;
  auctionType: string;
  bidIncrement: number;
  creatorId: string;
  itemId: string;
  stage: string;
  startDate: string;
  endDate: string;
  isFinished: boolean;
  id: string;
}

interface AuctionItemProps {
  id: string;
  description: string;
  category: string;
  name: string;
  photoURLs: string[];
}

interface AuctionViewProps {}

export default function AuctionView({}: AuctionViewProps) {
  const { auctionID } = useParams<{ auctionID: string }>();
  const [auction, setAuction] = useState<AuctionProps>({} as AuctionProps);
  const [item, setItem] = useState<AuctionItemProps>({} as AuctionItemProps);

  const getData = async () => {
    const auctionData = await fetch(
      `${process.env.REACT_APP_API_URL}auctions/${auctionID}`
    ).then((res) => res.json());
    setAuction(auctionData);
    const itemData = await fetch(
      `${process.env.REACT_APP_API_URL}items/${auctionData.itemId}`
    ).then((res) => res.json());
    setItem(itemData);
  };

  useEffect(() => {
    getData();
    console.log('fetching')
  }, []);

  return (
    <AppShell padding="md" navbar={<NavigationBar></NavigationBar>} fixed>
      <Grid>
        <AuctionInformationDashboard
          name={item.name}
          description={item.description}
          photoURLs={item.photoURLs}
        ></AuctionInformationDashboard>
        <AuctionBiddingDashboard
          auctionID={auction.id}
        ></AuctionBiddingDashboard>
        <Grid.Col span={10}>
          <Divider />
        </Grid.Col>
      </Grid>
    </AppShell>
  );
}
