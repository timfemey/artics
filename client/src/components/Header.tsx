import { Component } from "solid-js";
import Recommendations from "./Recommendations";

const Header: Component<{}> = () => {
  return (
    <>
      <header class="header">
        <div class="content">
          {/* <h1 class="heading">
            <span class="small">welcome in the world of</span>
            blog
            <span class="no-fill">writing</span>
          </h1> */}
          <a href="/editor" class="btn">
            write a blog
          </a>
        </div>
      </header>
      <Recommendations />
    </>
  );
};

export default Header;
