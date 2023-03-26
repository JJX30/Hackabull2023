import React from "react";
import Home from "./pages/Home";
import Landing from "./pages/Landing";
import Navbar from "./components/Navbar";
import Footer from "./components/Footer";
import Signin from "./components/Signin";
import { Routes, Route } from "react-router-dom";
import Rizzpository from "./pages/Rizzpository";

function App() {
  return (
    <>
      <Navbar></Navbar>
      <Routes>
        <Route path="/" element={<Landing></Landing>} />
        <Route path="/rizz" element={<Home></Home>}></Route>
        <Route path="/signin" element={<Signin></Signin>}></Route>
        <Route
          path="/rizzpository"
          element={<Rizzpository></Rizzpository>}
        ></Route>
      </Routes>
      <Footer></Footer>
    </>
  );
}

export default App;
