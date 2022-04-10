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
interface TitleProps {}

const Dashboard: FC<TitleProps> = ({}) => {
    const [opened, setOpened] = useState(false);
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
          <Title>Paieška</Title>
          <Select style={{width: 150}}
            width={20}
            label="Kategorija"
            placeholder="Kategorija"
            data={[
                {value: "-", label: "-"},
                { value: "Transportas", label: "Transportas" }
            ]}
          ></Select>
          <Select style={{width: 150}}
            width={20}
            label="Aukciono tipas"
            placeholder="Aukciono tipas"
            data={[
                {value: "-", label: "-"},
                { value: "Absoliutus", label: "Absoliutus" }
            ]}
          ></Select>
          <Button px={50} style={{top: 12}}>
            Ieškoti
            </Button>
           <Button color="teal" px={50} style={{top: 12, left: 400}} onClick={() => setOpened(true)}>
            Kurti aukcioną
            </Button>
        </Group>
        </div>
      <hr />
      <Container size="lg" style={{overflowY: 'scroll',overflowX: 'hidden', height: 800}}>
        <Grid gutter="xl">
          <AuctionGroup></AuctionGroup>
        </Grid>
      </Container>
      <AuctionCreateWizard onOpen={opened} onClose={() => setOpened(false)}></AuctionCreateWizard>
    </AppShell>
  );
};

export default Dashboard;
