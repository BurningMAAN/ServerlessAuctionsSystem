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
  isFinished: boolean;
  id: string;
}

interface AuctionItemProps {
  id: string;
  description: string;
  category: string;
  name: string;
}

interface AuctionWorkerData {
  AuctionID: string;
  Status: string;
  EndDate: string;
}

interface AuctionViewProps {}

export default function AuctionView({}: AuctionViewProps) {
  const { auctionID } = useParams<{ auctionID: string }>();
  const [auction, setAuction] = useState<AuctionProps>({} as AuctionProps);
  const [auctionWorker, setAuctionWorker] = useState<AuctionWorkerData>({} as AuctionWorkerData)
  const [item, setItem] = useState<AuctionItemProps>({} as AuctionItemProps);

  const getData = async () => {
    const auctionData = await fetch(
      `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctions/${auctionID}`
    ).then((res) => res.json());
    console.log("parsing auction");
    setAuction(auctionData);
    const itemData = await fetch(
      `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/items/${auctionData.itemId}`
    ).then((res) => res.json());
    setItem(itemData);
    const workerData = await fetch(
      `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctions/${auctionID}/worker`
    ).then((res) => res.json());
    setAuctionWorker(workerData)
  };

  useEffect(() => {
    console.log(auction.auctionType)
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
          status={auctionWorker.Status}
          auctionID={auction.id}
          startDate={auctionWorker.EndDate}
          auctionType={auction.auctionType}
          currentMaxBid={auction.bidIncrement} // pakeisti i max bid ar dar kazka
          bidIncrement={auction.bidIncrement}
          creatorID={auction.creatorId}
        ></AuctionBiddingDashboard>
        <Grid.Col span={10}>
          <Divider />
        </Grid.Col>
      </Grid>
    </AppShell>
  );
}
