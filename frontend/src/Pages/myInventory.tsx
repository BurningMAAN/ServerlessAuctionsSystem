import React, { FC, useState } from "react";
import NavigationBar from "../Components/Skeleton/Navbar";
import {
  AppShell,
  Divider,
  Group,
  Select,
  Title,
  Container,
  Grid,
  Button,
} from "@mantine/core";
import ItemCreateWizard from "../Components/Item/AddItemWizard";
import ItemGroup from "../Components/Item/ItemGroup";
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
              {value: "Transportas", label: "Transportas" },
              {value: "Baldai", label: "Baldai"}, {value: "Elektronika", label: "Elektronika"}, {value: "Automobilių detalės", label: "Automobilių detalės"}, {value: "Drabužiai", label: "Drabužiai"}, {label: "Paveikslas", value: "Paveikslas"}
          ]}
          ></Select>
          <Button px={50} style={{top: 12}}>
            Ieškoti
            </Button>
          <Button color="green" px={50} style={{top: 12}} onClick={() => setOpen(true)}>Pridėti inventorių</Button>
        </Group>
        </div>
        <Divider/>
        <Container size="lg" style={{overflowY: 'scroll',overflowX: 'hidden', height: 800}}>
        <Grid gutter="xl">
          <ItemGroup></ItemGroup>
        </Grid>
      </Container>
        <ItemCreateWizard onOpen={open} onClose={() => setOpen(false)}></ItemCreateWizard>
    </AppShell>
  );
};

export default MyInventory;
