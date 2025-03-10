import React, { useState } from "react";

import LoginForm from "../../components/LoginForm/LoginForm";
import axios from "axios";

const Login = () => {
  const [state, setState] = useState(null)

  const [formData, setFormData] = useState({
    login: "",
    password: ""
  });

  function onChange(e) {
    setFormData(oldData => {return {
      ...oldData,
      [e.target.name]: e.target.value
    }})
  }

  async function submit(e) {
    e.preventDefault()
    
    try {
      const response = await axios.post("http://193.227.240.146:8088/api/v1/auth/login", 
      formData,
      {
        headers: {
          "Content-Type": "application/json"
        }
      }
    )
    localStorage.setItem("token", response.data.token)
    console.log(response);
    } catch (error) {
      console.log(error);
      setState("Error")
      setTimeout(() => {
        setState(null)
      }, 1000);
    }
  }

  return (
    <section className="LoginPage page-root">
      <h1>Страница логина</h1>
      <form className="loginForm">
        <div className="loginForm__title">
          <h1>Авторизация</h1>
        </div>
        <div className="loginForm__input loginForm__input_login">
          <label htmlFor="login">Login</label>
          <input
            type="text" 
            id="login" 
            name="login" 
            onChange={onChange}
            value={formData.login} 
          />
        </div>
        <div className="loginForm__input loginForm__input_password">
          <label htmlFor="password">Password</label>
          <input
            type="password" 
            id="password"
            name="password"
            onChange={onChange}
            value={formData.password}
          />
        </div>
        <div className="loginForm__submit">
          <button
            className={state === "Error" ? "loginForm__button error" : "loginForm__button"}
            type="submit"
            onClick={submit}
          >Войти</button>
        </div>
      </form>
      
    </section>
  )
}

export default Login