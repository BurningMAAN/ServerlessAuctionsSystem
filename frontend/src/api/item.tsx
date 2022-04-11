
export interface ItemCreateRequest {
    name: string;
    description: string;
    category: string;
  }


export const createItem = async (item: ItemCreateRequest) => {
  const url =
    "https://garckgt6p0.execute-api.us-east-1.amazonaws.com/Stage/item";

    let tokenas:string = ""
    const token = sessionStorage.getItem("access_token");
    if(token){
      tokenas = token
    }

  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json", "access_token": unescape(tokenas)},
    body: JSON.stringify(item),
  };

  try {
    fetch(url, requestOptions).then((response) => {
      console.log(requestOptions)
     return response.json()
    }).then((responseJSON) => {
      console.log(responseJSON)
      console.log(token)
    });
  } catch (error) {
    console.log("failed to create item", error);
  }
};
