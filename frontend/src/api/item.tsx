
export interface ItemCreateRequest {
    name: string;
    description: string;
    category: string;
    photoURLs: string[];
  }


export const createItem = async (item: ItemCreateRequest) => {
  const url =
    `${process.env.REACT_APP_API_URL}item`;

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

  console.log(JSON.stringify(item))
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
