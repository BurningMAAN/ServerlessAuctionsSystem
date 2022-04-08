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
import { useState } from 'react';
import AuctionMetadata from './AuctionSelectItem';

interface AuctionProps {
  onOpen: boolean;
  onClose: () => void;
}

export default function AuctionCreateWizard({
  onOpen,
  onClose
}: AuctionProps) {
  return <AuctionMetadata onOpen={onOpen} onClose={onClose}></AuctionMetadata>;
}
