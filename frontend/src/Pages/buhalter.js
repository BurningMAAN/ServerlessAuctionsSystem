import React, { FC, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import AuctionCreateWizard from '../Components/Auction/CreateAuction/AuctionWizard';
import {
  AppShell,
  Container,
  Grid,
  Group,
  Select,
  Title,
  Button,
} from "@mantine/core";
import AuctionGroup from "../Components/Auction/AuctionGroup";
import { Redirect } from 'react-router-dom';
import {generateData} from "../api/generateData";
import jwtDecode, { JwtPayload } from "jwt-decode";
import CsvDownload from 'react-json-to-csv'

const getToken = () => {
    let tokenas = "";
    const tokenString = sessionStorage.getItem("access_token");
    if (tokenString) {
      tokenas = tokenString;
    }
    return tokenas;
  };

const BugalterDashboard =  () => {
  const [generate, setGenerate] = useState(false)
   const token = getToken();
   const downloadTxtFile = () => {
    const element = document.createElement("a");
    const file = new Blob([JSON.stringify(data)], {
      type: "text/plain"
    });
    element.href = URL.createObjectURL(file);
    element.download = "ataskaita.json";
    document.body.appendChild(element);
    element.click();
  };

  const [data, setData] = useState('')
  const getData = async () => {
    const auctionData = await fetch(
      `https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/generateData`
    ).then((res) => res.json());
    setData(auctionData)
    console.log(data)
  };
  return (
    <AppShell
      padding="md"
      navbar={<NavigationBar></NavigationBar>}
      style={{
          overflow: 'hidden'
      }}
    //   header={<HeaderMiddle></HeaderMiddle>}
      fixed
    >
        <div>
        <Group spacing="sm">

           <Button color="teal" px={50} style={{top: 12, left: 400}} onClick={() => {
             getData()
             downloadTxtFile()
             }}>
            Generuoti ataskaitas
            </Button>
        </Group>
        </div>
      <hr />
      <Container size="lg" style={{overflowY: 'scroll',overflowX: 'hidden', height: 800}}>
      </Container>
      {/* {generate && <CsvDownload data={`{"name": "stringas"}`}/>} */}
    </AppShell>
  );
};

export default BugalterDashboard;
