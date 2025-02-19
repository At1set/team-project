import React from "react";
import { useAuthContext } from "../context/SomeConext";
import { Navigate, useLocation } from "react-router-dom";

const RequiredAuth = ({children}) => {
  const location = useLocation();
  const { isAuth } = useAuthContext();

  if (!isAuth) {
    return <Navigate to="/" state={{redirectedFrom: location}}/>
  }

  return children;
}

export default RequiredAuth