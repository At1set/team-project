import React from "react";
import { Outlet, useLocation } from "react-router-dom";

const Wrapper = () => {
  const style = {}
  const location = useLocation()
  
  if (["/", "/login", "/admin"].includes(location.pathname)) style.backgroundColor = "#909090"
  
  return (
    <div className="wrapper" style={{...style}}>
      <Outlet />
    </div>
  )
}

export default Wrapper