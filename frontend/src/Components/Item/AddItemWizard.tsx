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
  Textarea,
  MantineTheme,
  TextInput,
} from "@mantine/core";
import { Upload, Photo, X, Icon as TablerIcon } from "tabler-icons-react";
import { Dropzone, DropzoneStatus, IMAGE_MIME_TYPE } from "@mantine/dropzone";
import {ItemCreateRequest, createItem} from "../../api/item";

interface ItemCreateProps {
  onOpen: boolean;
  onClose: () => void;
}


export default function ItemCreateWizard({ onOpen, onClose }: ItemCreateProps) {
  const theme = useMantineTheme();
  const [item, setItem] = useState<ItemCreateRequest>({} as ItemCreateRequest);

  return (
    <Modal opened={onOpen} onClose={onClose} size="xl">
      <Title order={1} onClick={() => console.log(item)}>
        Inventoriaus pridėjimas
      </Title>
      <TextInput
        label="Pavadinimas"
        description="Pavadinimas"
        onChange={(event) =>
          setItem({
            name: event.currentTarget.value,
            description: item.description,
            category: item.category,
          } as ItemCreateRequest)
        }
      />
      <Select
        label="Kategorija"
        placeholder="Pasirinkti"
        required
        onChange={(selectedItem) => {
          setItem({
            name: item.name,
            description: item.description,
            category: selectedItem,
          } as ItemCreateRequest);
        }}
        data={[{ value: "Transportas", label: "Transportas" }]}
      />
      <Textarea
        placeholder="Aprašymas"
        label="Aprašymas"
        onChange={(event) =>
          setItem({
            name: item.name,
            description: event.currentTarget.value,
            category: item.category,
          } as ItemCreateRequest)
        }
        required
      />
      <Title order={3}>Nuotraukos</Title>
      <Dropzone
        onDrop={() => console.log("accepted files")}
        onReject={() => console.log("rejected files")}
        maxSize={3 * 1024 ** 2}
        accept={IMAGE_MIME_TYPE}
      >
        {(status) => dropzoneChildren(status, theme)}
      </Dropzone>
      <Divider />
      <Center>
        <Button color="green" onClick={() => createItem(item)}>Patvirtinti</Button>
      </Center>
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
