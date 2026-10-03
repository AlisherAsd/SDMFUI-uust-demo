import axios from "axios";

export interface Item {
  id: number;
  title: string;
  description: string;
  price: number;
}

export const getItems = async (): Promise<Item[]> => {
  return axios
    .get("http://localhost:8082/items-service/list")
    .then((r) => r.data.items || []);
};
