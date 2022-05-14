import {
  Card,
  Image,
  Text,
  Badge,
  Button,
  Group,
  useMantineTheme,
} from "@mantine/core";
import { Link } from "react-router-dom";
import UpdateItem from "./UpdateItemModal";
import { useState } from "react";
import DeleteItem from "./DeleteItemModal";

export interface ItemProps {
  id: string;
  description: string;
  category: string;
  name: string;
  photoURLs: string[];
}

export default function ItemCard({
  name,
  category,
  description,
  id,
}: ItemProps) {
  const theme = useMantineTheme();
  const secondaryColor =
    theme.colorScheme === "dark" ? theme.colors.dark[1] : theme.colors.gray[7];

  const [updateOpen, setUpdateOpen] = useState(false)
  const [deleteOpen, setDeleteOpen] = useState(false)
  return (
    <div style={{ width: 340, margin: "auto" }}>
      <Card shadow="sm" p="lg">
        <Card.Section>
          <Image
            src="https://cdn.shopify.com/s/files/1/0773/9113/products/RoeblingProfile_5000x.jpg?v=1629750752"
            height={160}
            alt="Norway"
          />
        </Card.Section>

        <Group
          position="apart"
          style={{ marginBottom: 5, marginTop: theme.spacing.sm }}
        >
          <Text weight={500}>{name}</Text>
          <Badge color="pink" variant="light">
            {category}
          </Badge>
        </Group>
        <Button
        variant="light"
        color="blue"
        fullWidth
        style={{marginTop: 14}}
        onClick={() => setUpdateOpen(true)}>
          Atnaujinti
        </Button>
        <Button
        variant="light"
        color="red"
        fullWidth
        style={{marginTop: 14}}
        onClick={() => setDeleteOpen(true)}>
          Panaikinti
        </Button>
      </Card>
      <UpdateItem name={name} description={description} category={category} id={id}onOpen={updateOpen} onClose={() => setUpdateOpen(false)}/>
      <DeleteItem onOpen={deleteOpen} onClose={() => setDeleteOpen(false)} itemID={id} itemName={name}></DeleteItem>
    </div>
  );
}
