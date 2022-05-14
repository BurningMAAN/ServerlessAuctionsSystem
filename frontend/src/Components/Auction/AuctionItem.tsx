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
import { useParams, Redirect } from "react-router-dom";

export interface AuctionProps {
  auctionName: string;
  auctionDate: string;
  category: string;
  buyoutPrice: number;
  description: string;
  bidIncrement: number;
  auctionID: string;
  isFinished: boolean;
  photoURL: string;
}

export default function AuctionCard({
  auctionDate,
  buyoutPrice,
  auctionName,
  category,
  description,
  bidIncrement,
  auctionID,
  isFinished,
  photoURL,
}: AuctionProps) {
  const theme = useMantineTheme();
  const secondaryColor =
    theme.colorScheme === "dark" ? theme.colors.dark[1] : theme.colors.gray[7];
  return (
    <div style={{ width: 340, margin: "auto" }}>
      <Card shadow="sm" p="lg">
        <Card.Section>
          <Image
            src={`${process.env.REACT_APP_S3_URL}/${photoURL}`}
            height={160}
            alt="Norway"
          />
        </Card.Section>

        <Group
          position="apart"
          style={{ marginBottom: 5, marginTop: theme.spacing.sm }}
        >
          <Text weight={500}>{auctionName}</Text>
          {buyoutPrice && (
            <Badge color="green" variant="light">
              Pirkti dabar
            </Badge>
          )}
          <Badge color="pink" variant="light">
            {category}
          </Badge>
        </Group>

        <Text size="sm" style={{ color: secondaryColor, lineHeight: 1.5 }}>
          <b>Aukciono pradžia</b>:{" "}
          {(isFinished && "Aukcionas baigtas") || auctionDate}
          <br />
          <b>Minimalus kėlimas</b>: {bidIncrement}
          <br />
          <b>Aukciono tipas</b>: Absoliutus
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
