import "./App.css";
import Navbar from "./components/Navbar";
import Header from "./components/Header";
import { Route, Routes } from "@solidjs/router";
import Editor from "./components/Editor";

function App() {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" component={Header} />
        <Route path="/editor" component={Editor} />
      </Routes>
    </>
  );
}

export default App;
