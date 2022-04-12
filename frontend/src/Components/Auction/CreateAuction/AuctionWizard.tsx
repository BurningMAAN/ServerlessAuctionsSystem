import {
  Button,
  Modal,
  Stepper,
  Center,
  Title,
  Select,
  NumberInput,
  Divider,
} from "@mantine/core";
import { useState, useEffect } from "react";
import { DatePicker } from "@mantine/dates";

interface AuctionProps {
  onOpen: boolean;
  onClose: () => void;
}

interface ItemList {
  items: [
    {
      id: string;
      description: string;
      category: string;
      name: string;
      auctionId: string;
    }
  ];
}

interface CreateAuctionRequest {
  itemID: string;
  auctionDate: string;
  buyoutPrice: number;
  auctionType: string;
  bidIncrement: number;
}

interface SelectItemProps {
  label: string;
  value: string;
}

export default function AuctionCreateWizard({ onOpen, onClose }: AuctionProps) {
  const [activeStep, setActiveStepStepper] = useState(0);
  const [auctionMetadata, setAuctionMetadata] = useState<CreateAuctionRequest>({} as CreateAuctionRequest);
  const nextStep = () =>
    setActiveStepStepper((current) => (current < 3 ? current + 1 : current));
  const prevStep = () =>
    setActiveStepStepper((current) => (current > 0 ? current - 1 : current));

  const handleOnClose = (): void => {
    setActiveStepStepper(0);
    onClose();
  };

  const [userItemsList, setUserItemsList] = useState<ItemList>({} as ItemList);
  useEffect(() => {
    let tokenas: string = "";
    const token = sessionStorage.getItem("access_token");
    if (token) {
      tokenas = token;
    }

    const requestOptions = {
      method: "GET",
      headers: { access_token: unescape(tokenas) },
    };
    const url =
      "https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/user/items";

    const getUserItems = async () => {
      try {
        const response = await fetch(url, requestOptions);
        const responseJSON = await response.json();
        setUserItemsList(responseJSON);
      } catch (error) {
        console.log("failed to get data from api", error);
      }
    };
    getUserItems();
  }, []);

  const selectionItems: SelectItemProps[] = [];
  userItemsList.items?.map((userItem) => {
    if(!userItem.auctionId){
      selectionItems.push({ label: userItem.name, value: userItem.id });
    }
  });

  const createAuction = async () => {
    let tokenas: string = "";
    const token = sessionStorage.getItem("access_token");
    if (token) {
      tokenas = token;
    }
  
    const requestOptions = {
      method: "POST",
      headers: { access_token: unescape(tokenas) },
      body: JSON.stringify(auctionMetadata)
    };
    const url =
      "https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/auction";
  
    try {
      const response = await fetch(url, requestOptions);
      const responseJSON = await response.json();
    } catch (error) {
      console.log("failed to get data from api", error);
    }
  };

  const [pickedDate, setPickedDate] = useState<Date|null>(new Date());

  return (
    <Modal opened={onOpen} onClose={handleOnClose} size="xl">
      <Stepper active={activeStep} color="green">
        <Stepper.Step label="Inventoriaus pasirinkimas"></Stepper.Step>
        <Stepper.Step label="Aukciono duomenys"></Stepper.Step>
        <Stepper.Step label="Patvirtinimas"></Stepper.Step>
      </Stepper>
      <Divider />
      {activeStep == 0 && (
        <>
          <Select
            label="Inventoriaus pasirinkimas"
            placeholder="Pasirinkti"
            data={selectionItems}
            required
            onChange={(selectedItem) => {
              setAuctionMetadata({
                auctionDate: pickedDate?.toISOString(),
                buyoutPrice: auctionMetadata.buyoutPrice,
                auctionType: auctionMetadata.auctionType,
                bidIncrement:auctionMetadata.bidIncrement,
                itemID: selectedItem,
              } as CreateAuctionRequest)
            }}
          />
          <Divider />
          <Center>
            <Button onClick={nextStep}>Toliau</Button>
          </Center>
        </>
      )}
      {activeStep == 1 && (
        <>
          <Select
            label="Aukciono tipas"
            placeholder="Pasirinkti"
            required
            data={[{ value: "absoliutus", label: "Absoliutus" }]}
            onSelect={(selectedItem) => {
              setAuctionMetadata({
                auctionDate: auctionMetadata.auctionDate,
                buyoutPrice: auctionMetadata.buyoutPrice,
                auctionType: selectedItem.currentTarget.value,
                bidIncrement:auctionMetadata.bidIncrement,
                itemID: auctionMetadata.itemID,
              } as CreateAuctionRequest)
            }}
          />
          <DatePicker placeholder="Pasirinkti" label="Aukciono data" required 
          onChange={(date) =>  setPickedDate(date)}/>
          <NumberInput label="Išpirkimo kaina" placeholder="Įvesti"
          onChange={(inputValue) => {
            setAuctionMetadata({
              auctionDate: auctionMetadata.auctionDate,
              buyoutPrice: inputValue?.valueOf(),
              auctionType: auctionMetadata.auctionType,
              bidIncrement:auctionMetadata.bidIncrement,
              itemID: auctionMetadata.itemID,
            } as CreateAuctionRequest)
          }} />
          <NumberInput label="Minimalus kėlimas" placeholder="Įvesti"
          onChange={(inputValue) => {
            setAuctionMetadata({
              auctionDate: auctionMetadata.auctionDate,
              buyoutPrice: auctionMetadata.buyoutPrice,
              auctionType: auctionMetadata.auctionType,
              bidIncrement:inputValue?.valueOf(),
              itemID: auctionMetadata.itemID,
            } as CreateAuctionRequest)
          }}  />
          <Divider />
          <Center>
            <Button onClick={prevStep}>Atgal</Button>
            <Button onClick={nextStep}>Toliau</Button>
          </Center>
        </>
      )}
      {activeStep == 2 && (
        <>
          <Title order={1}>Step 3</Title>
          <Divider />
          <Center>
            <Button onClick={prevStep}>Atgal</Button>
            <Button color="green" onClick={() => {
              createAuction();
              handleOnClose();
            }}>
              Patvirtinti
            </Button>
          </Center>
        </>
      )}
    </Modal>
  );
}
