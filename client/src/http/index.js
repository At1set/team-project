import axios from 'axios';
import Token from '../services/Token.js';

export const SERVER_URL = `http://193.227.240.146:8080`

const $api = axios.create({
  withCredentials: true,
  baseURL: SERVER_URL
})

function noAuthRoute(uri) {
  return /^(\/login\/|\/refresh\/$)/.test(uri)
}


$api.interceptors.request.use(async (config) => {
  console.log(
    `Отправляю запрос на ${config.url}, для него аутентификация : ${
      noAuthRoute(config.url) ? "НУЖНА" : "НЕ НУЖНА"
    }`
  )

  if (noAuthRoute(config.url) || config._isRetry) return config
  config._isRetry = true

  await Token.synchronize()

  config.headers.Authorization = `Bearer ${localStorage.getItem("token")}`
  return config
})

export default $api;