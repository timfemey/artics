/* @refresh reload */
import { render } from "solid-js/web";

import "./index.css";
import App from "./App";
import { CounterProvider } from "./context";
import { Router } from "@solidjs/router";

const root = document.getElementById("root");

render(
  () => (
    <CounterProvider loggedIn={false}>
      <Router>
        <App />
      </Router>
    </CounterProvider>
  ),
  root!
);
