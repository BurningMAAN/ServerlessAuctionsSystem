
export interface ItemCreateRequest {
    name: string;
    description: string;
    category: string;
  }


export const createItem = async (item: ItemCreateRequest) => {
  const url =
    "https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/item";

    const token = sessionStorage.getItem("access_token");
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json", "access_token": `${token}` },
    body: JSON.stringify(item),
  };

  try {
    fetch(url, requestOptions).then((response) => {
     response.json()
     console.log(response)
     
    });
  } catch (error) {
    console.log("failed to create item", error);
  }
};
