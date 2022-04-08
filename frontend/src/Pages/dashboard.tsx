import React, { FC, useState } from "react";
import AuctionCard from "../Components/Auction/AuctionItem";
import {
  AuctionGroup,
  GetAuctionList,
} from "../Components/Auction/AuctionGroup";
import NavigationBar from "../Components/Skeleton/Navbar";
import AuctionCreateWizard from '../Components/Auction/CreateAuction/AuctionWizard';
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
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
          <Grid.Col span={4}>
            <AuctionItem
              auctionName="Daiktas"
              auctionDescription="Aprasymas"
            ></AuctionItem>
          </Grid.Col>
        </Grid>
      </Container>
      <AuctionCreateWizard onOpen={opened} onClose={() => setOpened(false)}></AuctionCreateWizard>
    </AppShell>
  );
};

export default Dashboard;
