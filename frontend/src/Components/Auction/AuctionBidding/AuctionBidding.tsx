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
  );
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
  const [activeBiddingState, setActiveBiddingState] = useState(
    "bids_before_auction_start"
  );

  if(isFinished){
    setActiveBiddingState("bids_auction_finish");
  }
  
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
    if (activeBiddingState === "bids_auction_start") {
      if (timeLeft == 0) {
        setActiveBiddingState("bids_auction_finish");
        return () => finishAuction(auctionID)
      }

      const intervalId = setInterval(() => {
        setTimeLeft(timeLeft - 1);
      }, 1000);
      return () => clearInterval(intervalId);
    }
  });

  const token = getToken();
  const decodedToken = jwtDecode<DecodedToken>(token);
  return (
    <Grid.Col span={4}>
      <Center>
        {auctionType == "absolute" && <Title>Absoliutus aukcionas</Title>}
        {auctionType == "reserved" && <Title>Rezervinis aukcionas</Title>}
      </Center>
      <Divider />
      <Center>
        <h4>Paskutinis statymas</h4>
      </Center>
      <Center>
        <h1>{currentMaxBid} €</h1>
      </Center>
      <Center>
        <ProgressCircle progressValue={timeLeft}></ProgressCircle>
      </Center>
      <Center>
        <Text>Minimalus kėlimas: {bidIncrement} €</Text>
      </Center>
      <Center>
        {(timeLeft !== 0 && token && decodedToken.username != creatorID && (
          <Button
            color="green"
            onClick={() => {
              console.log("atliktas statymas");
              setTimeLeft(30);
            }}
          >
            + {bidIncrement}
          </Button>
        )) ||
          (token && decodedToken.username != creatorID && (
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
        {activeBiddingState === "bids_before_auction_start" && (
          <Title order={6}>
            Aukcionas prasideda už {days} dienų {hours} valandų {minutes}{" "}
            minučių {seconds} sekundžių
          </Title>
        )}
        {activeBiddingState === "bids_auction_start" && (
          <Title order={6}>
            Aukcionas šiuo metu vyksta
          </Title>
        )}
        {activeBiddingState === "bids_auction_finish" && (
          <Title order={6}>
            Aukcionas baigtas
          </Title>
        )}
      </Center>
    </Grid.Col>
  );
}
