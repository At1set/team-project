import React from "react"

import { BrowserRouter, Routes, Route } from 'react-router-dom';

import ErrorPage from "../pages/Error/Error"
import Wrapper from "./Wrapper"
import HomePage from "../pages/Home/Home"
import AdminPage from "../pages/Admin/Admin"
import LoginPage from "../pages/Login/Login"
import NavigationBlock from "./NavigationBlock/NavigationBlock";

const AppRouter = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Wrapper />} >
          <Route element={<NavigationBlock />}>
            <Route index element={<HomePage />} />
            <Route path="admin" element={<AdminPage />} />
            <Route path="login" element={<LoginPage />} />
          </Route>
          <Route path="*" element={<ErrorPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default AppRouter