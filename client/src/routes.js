import React from "react";
import { Route, Routes, BrowserRouter } from "react-router-dom";

import MainPage from "./views/MainPage/MainPage";

// eslint-disable-next-line import/no-anonymous-default-export
export default () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<MainPage />} path="/" />
      </Routes>
    </BrowserRouter>
  );
};