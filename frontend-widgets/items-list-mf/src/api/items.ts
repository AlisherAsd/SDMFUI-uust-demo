import axios from "axios";

export interface Item {
  id: number;
  title: string;
  description: string;
  price: number;
}

export type ItemCreate = Omit<Item, "id">;

export const getItems = async (): Promise<Item[]> => {
  return axios
    .get("http://localhost:8082/items-service/list")
    .then((r) => r.data.items || []);
};

export const addItem = async (item: ItemCreate): Promise<Item[]> => {
  return axios.post("http://localhost:8082/items-service/item", item);
};
