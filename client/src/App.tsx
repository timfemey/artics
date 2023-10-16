import { createSignal } from "solid-js";
import "./App.css";

function App() {
  const [loggedIn, setLoggedIn] = createSignal(false);
  return (
    <>
      <nav class="navbar">
        <h1>Artics</h1>
        <ul class="links-container">
          <li class="link-item">
            <a href="/" class="link">
              {loggedIn() ? "Profile" : "Sign Up"}
            </a>
          </li>
          <li class="link-item">
            <a href="/editor" class="link">
              About Creator
            </a>
          </li>
        </ul>
      </nav>
    </>
  );
}

export default App;
