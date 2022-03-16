import React from "react";
import ReactDOM from "react-dom";
import App from "./App";

import { StoreProvider } from "./stores/context";
import * as stores from "./stores";

ReactDOM.render(
  <React.StrictMode>
    <StoreProvider value={stores}>
      <App />
    </StoreProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
