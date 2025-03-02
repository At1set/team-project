import React from "react";
import { NavLink } from "react-router-dom";
import { Outlet } from "react-router-dom";

import "./_NavigationBlock.scss"

const NavigationBlock = () => {
  return (
    <>
      <Outlet />
      <nav className="navigationBlock">
        <span>Доступные страницы: </span>
        <ul>
          <li><NavLink to={"/"} className={({isActive}) => isActive ? "_active" : null}>Домашняя (/)</NavLink></li>
          <li><NavLink to={"/login"} className={({isActive}) => isActive ? "_active" : null}>Авторизация (/login)</NavLink></li>
          <li><NavLink to={"/admin"} className={({isActive}) => isActive ? "_active" : null}>Авторизация для суперпользователя (/admin)</NavLink></li>
        </ul>
      </nav>
    </>
  )
}

export default NavigationBlock