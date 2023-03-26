import React from "react";
import styled from "styled-components";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <Wrapper>
      <ul className="links">
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/Rizzpository">Rizzpository</Link>
        </li>
        <li>
          <Link to="/signin">Sign In</Link>
        </li>
      </ul>
    </Wrapper>
  );
};

export default Navbar;

const Wrapper = styled.nav`
  display: flex;
  justify-content: space-between;
  background-color: black;
  height: 78px;
  font-weight: 200;
  font-family: Roboto, sans-serif;

  a {
    color: white;
    text-decoration: none;
  }
  li {
    list-style-type: none;
  }
  * {
    outline: none;
  }
  .links {
    display: flex;
    justify-content: space-around;
    align-items: center;
    width: 100%;
    padding: 0px;
    margin: 0px;
  }
`;
