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

interface AuctionProps {
  auctionID?: string;
  auctionName: string;
  auctionDescription: string;
}

export default function AuctionCard({
    auctionID,
  auctionName,
  auctionDescription,
}: AuctionProps) {
  const theme = useMantineTheme();
  const secondaryColor =
    theme.colorScheme === "dark" ? theme.colors.dark[1] : theme.colors.gray[7];
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
          <Text weight={500}>{auctionName}</Text>
          <Badge color="green" variant="light">
            Buy Now
          </Badge>
          <Badge color="pink" variant="light">
            Kategorija
          </Badge>
        </Group>

        <Text size="sm" style={{ color: secondaryColor, lineHeight: 1.5 }}>
          <b>Aukciono pradžia</b>: data
          <br />
          <b>Aprašymas</b>: {auctionDescription}
        </Text>

        <Button
          variant="light"
          color="blue"
          fullWidth
          style={{ marginTop: 14 }}
        >
          <Link to={`/auctions/${auctionID}`}>Peržiūrėti</Link>
        </Button>
      </Card>
    </div>
  );
}
