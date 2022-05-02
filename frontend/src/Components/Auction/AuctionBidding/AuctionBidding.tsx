import { Button, Center, Title, Text, Grid, Divider } from "@mantine/core";
import { useState, useEffect } from "react";
import ProgressCircle from "../../General/ProgressCircle";
import jwtDecode, { JwtPayload } from "jwt-decode";
import { useInterval } from 'usehooks-ts'

interface AuctionInput{
  auctionID: string;
}

interface GetAuctionResponse{
  auctionType: string;
  bidIncrement: number;
  startDate: string;
  creatorID: string;
  auctionID: string;
  stage: string;
  endDate: string;
}

interface DecodedToken {
  username: string;
}

interface Bid{
  bids: BidProps[] | null;
}

interface BidProps{
  auctionId: string;
  value: number;
  timestamp: string;
  userId: string;
}

const getToken = () => {
  let tokenas = "";
  const tokenString = sessionStorage.getItem("access_token");
  if (tokenString) {
    tokenas = tokenString;
  }
  return tokenas;
};

const placeBid = async (auctionID: string, bid: number) => {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json", "access_token": unescape(getToken())},
    body: JSON.stringify({value: bid})
  };
  return await fetch(
    `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auction/${auctionID}/bid`, requestOptions
  )
}

export default function AuctionBiddingDashboard({
  auctionID,
}: AuctionInput) {
  // Cia programuoju naujai
  const [auction, setAuction] = useState<GetAuctionResponse>({} as GetAuctionResponse)
  const getData = async () => {
    const auctionData = await fetch(
      `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctions/${auctionID}`
    ).then((res) => res.json());
    setAuction(auctionData);
  };

  useEffect(() => {
    getData();
  });
  const [timeLeft, setTimeLeft] = useState(30);
  const [days, setDays] = useState(0);
  const [hours, setHours] = useState(0);
  const [minutes, setMinutes] = useState(0);
  const [seconds, setSeconds] = useState(0);
  const [bids, setBids] = useState<Bid>({} as Bid)

  let refreshTime: number | null = 300
  useInterval(() => {
    if (auction.stage === "STAGE_ACCEPTING_BIDS") {
        const targetDate = new Date(auction.startDate);
        const now = new Date();
        const difference = targetDate.getTime() - now.getTime();

        const d = Math.floor(difference / (1000 * 60 * 60 * 24));
        setDays(d);

        const h = Math.floor(
          (difference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
        );
        setHours(h);

        const m = Math.floor((difference % (1000 * 60 * 60)) / (1000 * 60));
        setMinutes(m);

        const s = Math.floor((difference % (1000 * 60)) / 1000);
        setSeconds(s);
    } else {
      refreshTime = null
    }
  }, refreshTime);

  let timerInterval: number | null = 1000
  if (auction.stage === "STAGE_AUCTION_FINISHED"){
    timerInterval = null
  }
  useInterval(() => {
    const endDateTime = new Date(auction.endDate);
    const now = new Date();
    const difference = endDateTime.getTime() - now.getTime();
    const s = Math.floor((difference % (1000 * 60)) / 1000);
    console.log(s)
    if (s <= 0){
      setTimeLeft(0)
    }
    if(auction.stage == "STAGE_AUCTION_FINISHED"){
      setTimeLeft(0)
      return
    }
    if (auction.stage === "STAGE_AUCTION_ONGOING") {
      if (timeLeft <= 0) {
        return
      }
      setTimeLeft(s - 1);
    }
  }, timerInterval);

  
  const getLatestBids = async (auctionID: string) => {
    const requestOptions = {
      method: "GET",
      headers: {"access_token": unescape(getToken())},
    };
    const itemData = await fetch(
      `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auction/${auctionID}/bids`, requestOptions
    ).then((res) => res.json());
    setBids(itemData);
  }

  let getLatestBidsDelay: number | null = 300
   useInterval(() => {
     if(auctionID){
      getLatestBids(auctionID)
      getLatestBidsDelay = null
      return
     }
      getLatestBids(auctionID)
  }, getLatestBidsDelay)

  const token = getToken();
  const decodedToken = jwtDecode<DecodedToken>(token);
  return (
    <Grid.Col span={4}>
      <Center>
        {auction.auctionType == "AbsoluteAuction" && <Title>Absoliutus aukcionas</Title>}
        {auction.auctionType == "reserved" && <Title>Rezervinis aukcionas</Title>}
      </Center>
      <Divider />
      <Center>
        <h4>Paskutinis statymas</h4>
      </Center>
      <Center>
        <h1>{bids.bids == null && bids.bids == undefined && '0' || bids.bids![0].value} €</h1>
      </Center>
      <Center>
        <ProgressCircle progressValue={timeLeft}></ProgressCircle>
      </Center>
      <Center>
        <Text>Minimalus kėlimas: {auction.bidIncrement} €</Text>
      </Center>
      <Center>
        {(auction.stage != "STAGE_AUCTION_FINISHED" && timeLeft !== 0 && token && decodedToken.username != auction.creatorID && (
          <Button
            color="green"
            onClick={() => {
              let bidValue: number;
              if(bids.bids == null || bids.bids == undefined){
                bidValue = 0 + auction.bidIncrement
              }else{
                bidValue = bids.bids![0].value + auction.bidIncrement
              }
              placeBid(auctionID, bidValue)
              .then((response) => {
                if(response.status === 201){
                  setTimeLeft(30);
                }
              })
              .catch((error) => {
                console.log(error)
              })
            }}
          >
            + {auction.bidIncrement}
          </Button>
        )) ||
          (token && decodedToken.username != auction.creatorID && (
            <Button color="grey" disabled>
              Aukcionas baigėsi
            </Button>
          )) ||
          (!token && decodedToken.username != auction.creatorID && (
            <Button color="grey" disabled>
              Tik registruotiems nariams
            </Button>
          ))}
      </Center>
      <Center>
        {auction.stage === "STAGE_ACCEPTING_BIDS" && (
          <Title order={6}>
            Aukcionas prasideda už {days} dienų {hours} valandų {minutes}{" "}
            minučių {seconds} sekundžių
          </Title>
        )}
        {auction.stage === "STAGE_AUCTION_ONGOING" && (
          <Title order={6}>
            Aukcionas šiuo metu vyksta
          </Title>
        )}
        {auction.stage === "STAGE_AUCTION_FINISHED" &&  (
          <Title order={6}>
            Aukcionas baigtas
          </Title>
        )}
      </Center>
    </Grid.Col>
  );
}
