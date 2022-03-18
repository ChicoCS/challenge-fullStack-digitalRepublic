import React from "react";
import { Route, Routes, BrowserRouter } from "react-router-dom";

import MainPage from "./views/MainPage/MainPage";
import MainPageResult from "./views/MainPage/MainPageResult";

// eslint-disable-next-line import/no-anonymous-default-export
export default () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<MainPage />} path="/" />
        <Route element={<MainPageResult />} path="/paint/result" />
      </Routes>
    </BrowserRouter>
  );
};