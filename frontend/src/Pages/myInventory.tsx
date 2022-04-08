import React, { FC, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import {
  AppShell,
  Divider,
  Group,
  Select,
  Title,
  Button,
} from "@mantine/core";
import ItemCreateWizard from "../Components/Item/AddItemWizard";
interface TitleProps {}

const MyInventory: FC<TitleProps> = ({}) => {
  const [open, setOpen] = useState(false);
  return (
    <AppShell
      padding="md"
      navbar={<NavigationBar></NavigationBar>}
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
          <Button px={50} style={{top: 12}}>
            Ieškoti
            </Button>
          <Button color="green" px={50} style={{top: 12}} onClick={() => setOpen(true)}>Pridėti inventorių</Button>
        </Group>
        </div>
        <Divider/>
        <ItemCreateWizard onOpen={open} onClose={() => setOpen(false)}></ItemCreateWizard>
    </AppShell>
  );
};

export default MyInventory;
