import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import App from './App.tsx'
import {BrowserRouter, Route, Routes} from "react-router";
import ByCountry from "./ByCountry.tsx";


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path={"/"} element={<App/>}/>
        <Route path={"/country/:iso2"} element={<ByCountry/>}/>
      </Routes>
    </BrowserRouter>
  </StrictMode>
)
