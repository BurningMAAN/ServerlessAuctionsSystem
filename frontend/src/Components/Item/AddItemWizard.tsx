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
  Textarea, 
  MantineTheme,
  TextInput,
} from "@mantine/core";
import { Upload, Photo, X, Icon as TablerIcon } from 'tabler-icons-react';
import { Dropzone, DropzoneStatus, IMAGE_MIME_TYPE } from '@mantine/dropzone';

interface ItemCreateProps {
  onOpen: boolean;
  onClose: () => void;
}

export default function ItemCreateWizard({ onOpen, onClose }: ItemCreateProps) {
    const theme = useMantineTheme();
  return <Modal opened={onOpen} onClose={onClose} size="xl">
      <Title order={1}>Inventoriaus pridėjimas</Title>
      <TextInput label="Pavadinimas" description="Pavadinimas" />
         <Select
        label="Kategorija"
        placeholder="Pasirinkti"
        required
        data={[{ value: "Transportas", label: "Transportas" }]}
      />
       <Textarea
      placeholder="Aprašymas"
      label="Aprašymas"
      required
    />
      <Title order={3}>Nuotraukos</Title>
      <Dropzone
      onDrop={() => console.log('accepted files')}
      onReject={() => console.log('rejected files')}
      maxSize={3 * 1024 ** 2}
      accept={IMAGE_MIME_TYPE}
    >
         {(status) => dropzoneChildren(status, theme)}
    </Dropzone>
        <Divider/>
        <Center>
        <Button color="green">Patvirtinti</Button>
        </Center>
  </Modal>;
}

export const dropzoneChildren = (status: DropzoneStatus, theme: MantineTheme) => (
    <Group position="center" spacing="xl" style={{ minHeight: 220, pointerEvents: 'none' }}>
      <ImageUploadIcon status={status} style={{ color: getIconColor(status, theme) }} size={80} />
  
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
      ? theme.colors[theme.primaryColor][theme.colorScheme === 'dark' ? 4 : 6]
      : status.rejected
      ? theme.colors.red[theme.colorScheme === 'dark' ? 4 : 6]
      : theme.colorScheme === 'dark'
      ? theme.colors.dark[0]
      : theme.colors.gray[7];
  }