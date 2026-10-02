import axios from "axios"

export interface Remote {
  mfName: string
  componentName: string
  entryUrl :string
}


export const getRemotes = async (page: string): Promise<Remote[]> => {
    return axios.get(`http://localhost:8080/remotes/${page}`)
        .then(r => r.data.remotes || [])
}


export const createRemotes = async (remote: Remote): Promise<{ok: boolean, mess: string}> => {
    return axios.post(`http://localhost:8080/remotes`, remote)
        .then(() => ({
            ok: true,
            mess: 'Виджет успешно добавлен!',
        }))
        .catch(err => ({
            ok: false,
            mess: 'Произошла ошибка при добавлении виджета!' + err.message
        }))
}