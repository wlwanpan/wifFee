import axios from 'axios'

export default function() {

  let axiosConfig = {
    baseURL: 'http://localhost:3000/',
    responseType: 'arraybuffer',
    withCredentials: false
  }

  return axios.create(axiosConfig)
}