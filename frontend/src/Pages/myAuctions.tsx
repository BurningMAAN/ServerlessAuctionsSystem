import React, { FC } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import {
  AppShell,
  Tabs,
} from "@mantine/core";
import { Photo, MessageCircle} from 'tabler-icons-react';
interface TitleProps {}

const MyAuctions: FC<TitleProps> = ({}) => {
  return (
    <AppShell
      padding="md"
      navbar={<NavigationBar></NavigationBar>}
    //   header={<HeaderMiddle></HeaderMiddle>}
      fixed
    >
      <Tabs>
      <Tabs.Tab label="Aktyvūs Aukcionai" icon={<Photo size={14} />}>Aktyvūs aukcionai</Tabs.Tab>
      <Tabs.Tab label="Baigęsi Aukcionai" icon={<MessageCircle size={14} />}>Baigęsi aukcionai</Tabs.Tab>
    </Tabs>
    </AppShell>
  );
};

export default MyAuctions;
