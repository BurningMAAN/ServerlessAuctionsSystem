import {
    Button,
    Center,
    Title,
    Text,
    Grid,
    Divider,
  } from "@mantine/core";
  import {useState, useEffect} from 'react'
  import ProgressCircle from "../../General/ProgressCircle";
  import jwtDecode, { JwtPayload } from "jwt-decode";

  interface AuctionBiddingProps {
    auctionType: string;
    currentMaxBid: number;
    bidIncrement: number;
    startDate: Date;
    creatorID: string;
  }

  interface DecodedToken{
    username: string;
  }
  
  const getToken = () => {
    let tokenas = ""
    const tokenString = sessionStorage.getItem('access_token');
    if(tokenString){
      tokenas = tokenString
    }
    return tokenas
  };

  export default function AuctionBiddingDashboard({ auctionType, currentMaxBid, bidIncrement, startDate , creatorID}: AuctionBiddingProps) {
    const [timeLeft, setTimeLeft] = useState(30);

    // const AuctionTimer = useEffect(() => {
    //   if (timeLeft == 0) {
    //     console.log('aukcionas baigesi')
    //     return;
    //   }
  
    //   const intervalId = setInterval(() => {
    //     setTimeLeft(timeLeft - 1);
    //   }, 1000);
    //   return () => clearInterval(intervalId);
    // });

    const token = getToken()
    const decodedToken = jwtDecode<DecodedToken>(token);
  return (
    <Grid.Col span={4}>
      <Center>
          {auctionType == "absolute" && <Title>Absoliutus aukcionas</Title>}
          {auctionType == "reserved" && <Title>Rezervinis aukcionas</Title>}
      </Center>
      <Divider/>
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
            {(timeLeft !== 0 && token && decodedToken.username != creatorID && <Button color="green" onClick={() => {
              console.log('atliktas statymas')
              setTimeLeft(30)
            }}>+ {bidIncrement}</Button>) || token && (
              <Button color="grey" disabled>
                Aukcionas baigėsi
              </Button>
            ) || !token && (
              <Button color="grey" disabled>
                Tik registruotiems nariams
              </Button>
            )}
          </Center>
    </Grid.Col>
    );
  }
  