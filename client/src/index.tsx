/* @refresh reload */
import { render } from "solid-js/web";

import "./index.css";
import App from "./App";
import { CounterProvider } from "./context";

const root = document.getElementById("root");

render(
  () => (
    <CounterProvider loggedIn={false}>
      <App />
    </CounterProvider>
  ),
  root!
);
