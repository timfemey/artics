import { Component } from "solid-js";
import { useCounter } from "../context";

const Navbar: Component<{}> = () => {
  //@ts-ignore
  const [loggedIn, { increment, decrement }] = useCounter();
  return (
    <nav class="navbar">
      <h1>Artics</h1>
      <ul class="links-container">
        <li class="link-item">
          <a href="/" class="link">
            {loggedIn() ? "Profile" : "Sign Up"}
          </a>
        </li>
        <li class="link-item">
          <a href="https://femi-port.web.app/" target="_blank" class="link">
            About Creator
          </a>
        </li>
      </ul>
    </nav>
  );
};

export default Navbar;
