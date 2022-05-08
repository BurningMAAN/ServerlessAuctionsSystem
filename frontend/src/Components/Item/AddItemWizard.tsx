import { useState } from "react";
import {
  Button,
  Modal,
  Center,
  Title,
  Select,
  Text,
  Group,
  Divider,
  useMantineTheme,
  Grid,
  Textarea,
  MantineTheme,
  Image,
  TextInput,
} from "@mantine/core";
import { Upload, Photo, X, Icon as TablerIcon } from "tabler-icons-react";
import { Dropzone, DropzoneStatus, IMAGE_MIME_TYPE } from "@mantine/dropzone";
import {ItemCreateRequest, createItem} from "../../api/item";
import { useForm } from '@mantine/form';
import AWS from 'aws-sdk'
import uuid from "uuid";

const S3_BUCKET ='auctioneer-images-bucket';
const REGION ='us-east-1';


AWS.config.update({
    accessKeyId: 'AKIARPU6BCUFUG5PQY6N',
    secretAccessKey: '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'
})

const myBucket = new AWS.S3({
    params: { Bucket: S3_BUCKET},
    region: REGION,
})

interface ItemCreateProps {
  onOpen: boolean;
  onClose: () => void;
}


export default function ItemCreateWizard({ onOpen, onClose }: ItemCreateProps) {
  const theme = useMantineTheme();
  const [item, setItem] = useState<ItemCreateRequest>({} as ItemCreateRequest);
  const [uploadedImages, setUploadedImages] = useState<string[]>([])
  const [photoImages, setPhotoImages] = useState<File[]>([])
  const [imageNames, setImageNames] = useState<string[]>([])
  let testImages = new FormData();

  const form = useForm({
    initialValues: {
      name: '',
      description: '',
      category: '',
      photoURLs: []
    },
    validate: {
      name: (value) => value.length >= 4 ? null : 'Daikto pavadinimas turi būti bent 4 simbolių',
      description: (value) => value.length > 10 ? null : 'Daikto aprašymas turi būti bent 10 simbolių',
      category: (value) => value == 'Transportas' ? null : 'Pasirinkite tinkamą kategoriją'
    }
  })
  return (
    <Modal opened={onOpen} onClose={() => {
      setUploadedImages([])
      onClose();
    }} size="xl">
      <form onSubmit={form.onSubmit((values) => {
          //  setUploadedImages([])
          let imageUUID = uuid.v4()
         photoImages.map((uploadedImage) => {
            const params = {
              ACL: 'public-read',
              Body: uploadedImage,
              Bucket: S3_BUCKET,
              Key: imageUUID
          };
  
          myBucket.putObject(params)
              .send((err) => {
                  if (err) console.log(err)
              })
          })
          imageNames.push(imageUUID)
           createItem(values)
           onClose();
        })}>
      <Title order={1} >
        Inventoriaus pridėjimas
      </Title>
      <TextInput
        label="Pavadinimas"
        description="Pavadinimas"
        {...form.getInputProps('name')}
      />
      <Select
        label="Kategorija"
        placeholder="Pasirinkti"
        required
        {...form.getInputProps('category')}
        data={[{ value: "Transportas", label: "Transportas" }]}
      />
      <Textarea
        placeholder="Aprašymas"
        label="Aprašymas"
        {...form.getInputProps('description')}
        required
      />
      <Title order={3}>Nuotraukos</Title>
      {uploadedImages.length > 0 && (
        <Grid>
          {
            uploadedImages.map((image) => {
              return (
                <Grid.Col span={4}><Image
                width={200}
                height={120}
                style={{maxHeight: 'auto', maxWidth: '100%', objectFit: 'fill'}}
                src={image}
                alt="With default placeholder"
              /></Grid.Col>
              )
            })
          }
        </Grid>
      )}
      <Divider/>
      <Dropzone
        onDrop={ async (images)=> {
          images.map((image) => {
            const url = URL.createObjectURL(image)
            setUploadedImages([...uploadedImages, url])
            testImages.append('id', image)
            console.log(testImages.get('id'))
            photoImages.push(image)
          })
        }}
        onReject={() => console.log("rejected files")}
        maxSize={3 * 1024 ** 2}
        accept={IMAGE_MIME_TYPE}
      >
        {(status) => dropzoneChildren(status, theme)}
      </Dropzone>
      <Center>
        <Button color="green" type="submit">Patvirtinti</Button>
      </Center>
      </form>
    </Modal>
  );
}

export const dropzoneChildren = (
  status: DropzoneStatus,
  theme: MantineTheme
) => (
  <Group
    position="center"
    spacing="xl"
    style={{ minHeight: 220, pointerEvents: "none" }}
  >
    <ImageUploadIcon
      status={status}
      style={{ color: getIconColor(status, theme) }}
      size={80}
    />

    <div>
      <Text size="xl" inline>
        Tempkite nuotrauką(-as) norint įkelti inventoriaus nuotrauką(-as)
      </Text>
    </div>
  </Group>
);

function ImageUploadIcon({
  status,
  ...props
}: React.ComponentProps<TablerIcon> & { status: DropzoneStatus }) {
  if (status.accepted) {
    return <Upload {...props} />;
  }

  if (status.rejected) {
    return <X {...props} />;
  }

  return <Photo {...props} />;
}

function getIconColor(status: DropzoneStatus, theme: MantineTheme) {
  return status.accepted
    ? theme.colors[theme.primaryColor][theme.colorScheme === "dark" ? 4 : 6]
    : status.rejected
    ? theme.colors.red[theme.colorScheme === "dark" ? 4 : 6]
    : theme.colorScheme === "dark"
    ? theme.colors.dark[0]
    : theme.colors.gray[7];
}
