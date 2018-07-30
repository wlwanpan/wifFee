import axios from 'axios'

export default function() {

  let axiosConfig = {
    baseURL: 'http://localhost:3000/',
    withCredentials: false
  }

  return axios.create(axiosConfig)
}