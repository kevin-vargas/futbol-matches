import React from 'react'
import { Routes, Route, BrowserRouter } from "react-router-dom";
import Principal from "../components/principal/principal";
import CreateMatch from "../components/match/createMatch";
import Error from "../components/error/error";
import Footer from "../components/footer/footer";
import Register from "../components/register/register";
import Header from "../components/header/header";
import Login from "../components/login/login";

export const Router = () => {
    return (
      <BrowserRouter>
          <Header />
          <Routes>
              <Route path="/" element={<Login />} />
              <Route path="/register" element={<Register />} />
              <Route path="/principal" element={<Principal />} />
              <Route path="/create-match" element={<CreateMatch />} />
              <Route path="*" element={<Error />} />
          </Routes>
          <Footer />
      </BrowserRouter>
    );
}
