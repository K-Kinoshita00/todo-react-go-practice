import Axios from 'axios'
import { getToken } from '../../providers/auth'

export const axios = Axios.create({
  baseURL: '',
})

axios.interceptors.request.use((config) => {
  const token = getToken()
  if (token != null) {
    config.headers.Accept = 'application/json'
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})
