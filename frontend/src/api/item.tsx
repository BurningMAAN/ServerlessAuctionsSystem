
export interface ItemCreateRequest {
    name: string;
    description: string;
    category: string;
    ownerID: string;
  }

export const createItem = async (item: ItemCreateRequest) => {
  const url =
    "https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/item";

  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(item),
  };

  try {
    fetch(url, requestOptions).then((response) => response.json());
  } catch (error) {
    console.log("failed to create item", error);
  }
};
