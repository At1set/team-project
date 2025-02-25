import { createContext, useContext } from "react";

export const SomeContext = createContext(null);

export function useSomeContext() {
  return useContext(SomeContext)
}