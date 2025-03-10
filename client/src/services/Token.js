import axios, { AxiosError } from "axios";
import { jwtDecode as jwt_decode } from "jwt-decode"
import { SERVER_URL } from "../http";


export default class Token {
  static async update() {
    try {
      console.log(`Обновляем токен: ${SERVER_URL}/auth/refresh`)
      const response = await axios.get(`${SERVER_URL}/auth/refresh`, {
        withCredentials: true,
      })
      console.log(`Получен ответ: ${response.data}`);
      localStorage.setItem("token", response.data.token)
      console.log("Новый токен успешно установлен!");
      return response
    } catch (error) {
      console.error("Ошибка обновления токена: ", error)
      localStorage.removeItem("token")
      console.log("Вы не авторизованы!")
      if (error instanceof AxiosError) throw TokenError.UpdateTokenFail(error)
      throw error
    }
  }

  static async synchronize() {
    console.log("Начинаю проверку авторизации пользователя")
    const token = localStorage.getItem("token")
    if (!token) {
      console.log("Вы не авторизованы!");
      throw TokenError.MissingToken()
    }

    const payload = jwt_decode(token)
    const expTime = payload.exp * 1000
    const curTime = new Date().getTime()

    const expTimeData = new Date(expTime)
    const formattedDate = expTimeData.toLocaleString("ru-RU", {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
    })
    console.log(`Токен истекает: ${formattedDate}`)
    await new Promise(resolve => setTimeout(resolve, 3000))

    const updateTime = 2 * 60 * 1000 // 2 minutes in ms

    if (expTime - curTime <= -updateTime) {
      console.log("Токен скоро истекает, обновляем...")
      await Token.update()
    }
    console.log("Токен успешно синхронизирован!");

    return payload
  }
}


export class TokenError extends Error {
  message
  axiosErr

  constructor(message, axiosErr = null) {
    super(message)
    this.axiosErr = axiosErr
  }

  static MissingToken() {
    return new TokenError("Token is missing!")
  }

  static UpdateTokenFail(axiosErr) {
    return new TokenError("An error occured while trying update the token!", axiosErr)
  }
}