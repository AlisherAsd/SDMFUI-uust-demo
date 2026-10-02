import axios from "axios"
import { toast } from "vue3-toastify";

export const getRemotes = async (page: string) => {
    return axios.get(`http://localhost:8080/remotes/${page}`)
        .then(r => r.data.remotes || [])
        .catch(err => {
             toast("Ошибка при построении страницы", {
                type: "error"
            });
        })
}