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
  buyoutPrice: number;
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
}

interface AuctionViewProps {}

export default function AuctionView({}: AuctionViewProps) {
  const { auctionID } = useParams<{ auctionID: string }>();
  const [auction, setAuction] = useState<AuctionProps>({} as AuctionProps);
  const [item, setItem] = useState<AuctionItemProps>({} as AuctionItemProps);

  const getData = async () => {
    const auctionData = await fetch(
      `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctions/${auctionID}`
    ).then((res) => res.json());
    setAuction(auctionData);
    console.log(auction)
    const itemData = await fetch(
      `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/items/${auctionData.itemId}`
    ).then((res) => res.json());
    setItem(itemData);
  };

  useEffect(() => {
    getData();
  }, []);

  return (
    <AppShell padding="md" navbar={<NavigationBar></NavigationBar>} fixed>
      <Grid>
        <AuctionInformationDashboard
          name={item.name}
          description={item.description}
        ></AuctionInformationDashboard>
        <AuctionBiddingDashboard
          stage={auction.stage}
          auctionID={auction.id}
          startDate={auction.startDate}
          auctionType={auction.auctionType}
          currentMaxBid={auction.bidIncrement} // pakeisti i max bid ar dar kazka
          bidIncrement={auction.bidIncrement}
          creatorID={auction.creatorId}
          endDate={auction.endDate}
        ></AuctionBiddingDashboard>
        <Grid.Col span={10}>
          <Divider />
        </Grid.Col>
      </Grid>
    </AppShell>
  );
}
