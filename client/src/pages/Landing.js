import React from "react";
import styled from "styled-components";
import Signup from "../components/Signup";

const Landing = () => {
  return (
    <Wrapper>
      <div className="body-1">
        <footer>
          <div className="footer-logo">
            <p className="footer-logo-heading">Rizzerator</p>
            <p className="footer-logo-content">the confidence model</p>
          </div>
        </footer>
      </div>
      <div className="middle">
        <Signup></Signup>
      </div>
    </Wrapper>
  );
};

export default Landing;

const Wrapper = styled.div`
  display: flex;
  flex-direction: row;
  .footer-logo {
    display: flex;
    flex-direction: column;
    justify-content: center;
    text-shadow: 2px 4px 4px rgba(0, 0, 0, 0.5);
  }
  .footer-logo-heading {
    font-size: 80px;
    margin: 0px;
    font-weight: 500;
  }

  .footer-logo-content {
    font-size: 30px;
    margin: 0;
    font-weight: 300;
    opacity: 70%;
  }
  footer {
    bottom: 0;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    align-items: center;
    font-family: Roboto, sans-serif;
    color: white;
  }
  .body-1 {
    display: flex;
    justify-content: center;
    align-items: center;
    justify-content: space-evenly;
    align-items: center;
    background-image: url("../../abstract.webp");

    background-size: cover;
    /* Set a specific height */
    min-height: 90vh;
    width: 50%;

    /* Create the parallax scrolling effect */
    background-attachment: fixed;
    background-repeat: no-repeat;
  }
  .middle {
    display: flex;
    width: 50%;
    justify-content: center;
    align-items: center;
  }
`;
