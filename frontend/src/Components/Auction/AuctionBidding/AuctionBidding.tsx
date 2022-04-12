import { Button, Center, Title, Text, Grid, Divider } from "@mantine/core";
import { useState, useEffect } from "react";
import ProgressCircle from "../../General/ProgressCircle";
import jwtDecode, { JwtPayload } from "jwt-decode";

interface AuctionBiddingProps {
  auctionType: string;
  currentMaxBid: number;
  bidIncrement: number;
  startDate: string;
  creatorID: string;
  auctionID: string;
  isFinished: boolean;
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

interface PlaceBidRequest{
  auctionId: string;
  value: number;
}

const getToken = () => {
  let tokenas = "";
  const tokenString = sessionStorage.getItem("access_token");
  if (tokenString) {
    tokenas = tokenString;
  }
  return tokenas;
};

const finishAuction = async (auctionID: string) => {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json", "access_token": unescape(getToken())},
  };
  const finishedAuction = await fetch(
    `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auctions/${auctionID}/finish`, requestOptions
  ).catch((error) => console.log(error));
}

const placeBid = async (auctionID: string, bid: number) => {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json", "access_token": unescape(getToken())},
    body: JSON.stringify({value: bid})
  };
  await fetch(
    `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auction/${auctionID}/bid`, requestOptions
  ).catch((error) => console.log(error));
}

export default function AuctionBiddingDashboard({
  auctionType,
  currentMaxBid,
  bidIncrement,
  startDate,
  creatorID,
  isFinished,
  auctionID
}: AuctionBiddingProps) {
  const [timeLeft, setTimeLeft] = useState(30);
  const [days, setDays] = useState(0);
  const [hours, setHours] = useState(0);
  const [minutes, setMinutes] = useState(0);
  const [seconds, setSeconds] = useState(0);
  const [bids, setBids] = useState<Bid>({} as Bid)
  const [nextBidValue, setNextBidValue] = useState(0);
  const [activeBiddingState, setActiveBiddingState] = useState(
    "bids_before_auction_start"
  );

  useEffect(() => {
    if (activeBiddingState === "bids_before_auction_start") {
      const interval = setInterval(() => {
        const targetDate = new Date(startDate);
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

        if (d <= 0 && h <= 0 && m <= 0 && s <= 0) {
          setActiveBiddingState("bids_auction_start");
          return () => clearInterval(interval);
        }
      }, 1000);

      return () => clearInterval(interval);
    }
  });

  useEffect(() => {
    if(isFinished){
      setTimeLeft(0)
      return
    }
    if (activeBiddingState === "bids_auction_start") {
      if (timeLeft == 0) {
        setTimeLeft(30)
        return
        // setActiveBiddingState("bids_auction_finish");
        // return () => finishAuction(auctionID)
      }

      const intervalId = setInterval(() => {
        setTimeLeft(timeLeft - 1);
      }, 1000);
      return () => clearInterval(intervalId);
    }
  });

  
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

  useEffect(() => {
    const intervalId = setInterval(() => {
      getLatestBids(auctionID)
    }, 1000);
    return () => clearInterval(intervalId);
  });

  const token = getToken();
  const decodedToken = jwtDecode<DecodedToken>(token);
  return (
    <Grid.Col span={4}>
      <Center>
        {auctionType == "AbsoluteAuction" && <Title>Absoliutus aukcionas</Title>}
        {auctionType == "reserved" && <Title>Rezervinis aukcionas</Title>}
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
        <Text>Minimalus kėlimas: {bidIncrement} €</Text>
      </Center>
      <Center>
        {!isFinished && (timeLeft !== 0 && token && decodedToken.username != creatorID && (
          <Button
            color="green"
            onClick={() => {
              if(bids.bids == null || bids.bids == undefined){
                console.log("nera statymu")
                setNextBidValue(0 + bidIncrement)
              }else{
                console.log("kitas statymas")
                setNextBidValue(bids.bids![0].value + bidIncrement)
              }
              console.log('bid placed: ', nextBidValue)
              placeBid(auctionID, nextBidValue)
              setTimeLeft(30);
            }}
          >
            + {bidIncrement}
          </Button>
        )) ||
          (isFinished && token && decodedToken.username != creatorID && (
            <Button color="grey" disabled>
              Aukcionas baigėsi
            </Button>
          )) ||
          (!token && decodedToken.username != creatorID && (
            <Button color="grey" disabled>
              Tik registruotiems nariams
            </Button>
          ))}
      </Center>
      <Center>
        {!isFinished && activeBiddingState === "bids_before_auction_start" && (
          <Title order={6}>
            Aukcionas prasideda už {days} dienų {hours} valandų {minutes}{" "}
            minučių {seconds} sekundžių
          </Title>
        )}
        {!isFinished && activeBiddingState === "bids_auction_start" && (
          <Title order={6}>
            Aukcionas šiuo metu vyksta
          </Title>
        )}
        {isFinished && (
          <Title order={6}>
            Aukcionas baigtas
          </Title>
        )}
      </Center>
    </Grid.Col>
  );
}
