import axios from "axios";
import { toast } from "vue3-toastify";

export interface Context {
  user?: {
    username?: string;
    isAuthorization?: boolean;
  };
}

export const getContext = async (): Promise<Context> => {
  return axios
    .get(`http://localhost:8081/bff/context`)
    .then((r) => r.data.context || {})
    .catch((err) => {
      toast("Ошибка при данные о пользователе: " + err.error, {
        type: "error",
      });
    });
};
