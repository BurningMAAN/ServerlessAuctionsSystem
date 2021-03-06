import {
  Button,
  Modal,
  Stepper,
  Center,
  List,
  Title,
  Select,
  NumberInput,
  Divider,
} from "@mantine/core";
import { useState, useEffect } from "react";
import { DatePicker, TimeInput } from "@mantine/dates";
import { useForm } from '@mantine/form';
import { Item } from "react-bootstrap/lib/Breadcrumb";

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
      `${process.env.REACT_APP_API_URL}user/items`;

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
    if(userItem?.auctionId == ""){
      selectionItems.push({ label: userItem.name, value: userItem.id });
    }
  });

  const createAuction = async (auction: CreateAuctionRequest) => {
    let tokenas: string = "";
    const token = sessionStorage.getItem("access_token");
    if (token) {
      tokenas = token;
    }
  
    const requestOptions = {
      method: "POST",
      headers: { access_token: unescape(tokenas) },
      body: JSON.stringify(auction)
    };
    const url =
      `${process.env.REACT_APP_API_URL}auction`;

    try {
      const response = await fetch(url, requestOptions);
      const responseJSON = await response.json();
    } catch (error) {
      console.log("failed to get data from api", error);
    }
  };

  const form = useForm({
    initialValues: {
      itemID: '',
      auctionType: '',
      auctionDate: '',
      bidIncrement: 0,
    },
    validate: {
      itemID: (value) => value.length >= 4 ? null : 'Daikto pavadinimas turi b??ti bent 4 simboli??',
      auctionType: (value) => value == "AbsoluteAuction"? null : "Pasirinkite aukciono tip??",
      auctionDate: (value) => {
        const inputDate = Date.parse(value)
        const currentDate = Date.now()
        return inputDate - currentDate >= 0 ? null : "Data turi b??ti ateityje arba ??ios dienos data"
      },
      bidIncrement: (value) => value > 0 ? null: "??veskite minimalaus statymo sum??"
    }
  })

  return (
    <Modal opened={onOpen} onClose={handleOnClose} size="xl">
      <form onSubmit={form.onSubmit((values) => {
        createAuction(values);
        handleOnClose();
        })}>
      <Stepper active={activeStep} color="green">
        <Stepper.Step label="Inventoriaus pasirinkimas"></Stepper.Step>
        <Stepper.Step label="Aukciono duomenys"></Stepper.Step>
      </Stepper>
      <Divider />
      {activeStep == 0 && (
        <>
          <Select
            label="Inventoriaus pasirinkimas"
            placeholder="Pasirinkti"
            data={selectionItems}
            required
            {...form.getInputProps('itemID')}
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
            data={[{ value: "AbsoluteAuction", label: "Absoliutus" }]}
            {...form.getInputProps('auctionType')}
          />
          <DatePicker placeholder="Pasirinkti" label="Aukciono data" required 
          onChange={(date) => {
            form.setFieldValue("auctionDate", date?.toISOString()!)
          }}
          />
          <TimeInput label="Laikas"
          placeholder="12:00"
          onChange={(time) => {
           let dateSelected = new Date(form.getInputProps("auctionDate").value)
           dateSelected.setHours(time.getHours(), time.getMinutes())
           form.setFieldValue("auctionDate", dateSelected.toISOString())
          }}
          required></TimeInput>
          <NumberInput label="Minimalus k??limas" placeholder="??vesti"
          {...form.getInputProps('bidIncrement')} />
          <Divider />
          <Center>
            <Button onClick={prevStep}>Atgal</Button>
            <Button color="green" type="submit">
              Patvirtinti
            </Button>
          </Center>
        </>
      )}
      </form>
    </Modal>
  );
}
