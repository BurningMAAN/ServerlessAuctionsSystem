import React, { FC } from 'react';
import {CardGroup, Row, Col} from 'react-bootstrap';
import AuctionCard from './AuctionItem';

export interface GetAuctionList {
    auctionName: string
    auctionDescription: string
}

interface AuctionGroupChildrenProps {
    data: GetAuctionList[]
  }

export const AuctionGroup: FC<AuctionGroupChildrenProps> = ({data}) => {
  return (
    <CardGroup>
        {
            data.map((auctionItem) => {
                <Col>
                    <AuctionCard auctionName={auctionItem.auctionName} auctionDescription={auctionItem.auctionDescription}></AuctionCard>
                </Col>
            })
        }
    </CardGroup>
  );
};
