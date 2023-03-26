import React from "react";
import styled from "styled-components";

const Footer = () => {
  return (
    <Wrapper>
      <div className="body">
        <footer>
          <div className="footer-div-1">
            <div className="footer-logo">
              <p className="footer-logo-heading">Rizzerator</p>
              <p className="footer-logo-content">the confidence model</p>
            </div>
          </div>
          <div className="footer-div-2">
            <div className="footer-built">
              <p className="footer-built-heading">built with</p>
              <p className="footer-built-content">
                <a href="https://reactjs.org/" className="footer-built-link">
                  React
                </a>
                <br></br>
                <a href="https://go.dev/" className="footer-built-link">
                  Golang
                </a>
                <br></br>
                <a href="https://pytorch.org/" className="footer-built-link">
                  Pytorch
                </a>
              </p>
            </div>
          </div>
          <div className="footer-div-3">
            <div className="footer-about">
              <p className="footer-about-heading">about us</p>
              <p className="footer-about-content">
                The Rizzerator is a machine learning classification model
                designed to help improve your speech. Trained on hundreds of
                thousands of speeches it has found nuances in word choice,
                candence, and emotions to accurately rate the confidence (rizz)
                of an audio file. The Rizzerator will take you beyond your
                potential, allowing you to analyze and improve your speech has
                never been easier - with rizzz.
              </p>
            </div>
          </div>
        </footer>
        <div className="footer-copyright">
          <p>@ 2023 Rizzerator ® All rights reserved</p>
        </div>
      </div>
    </Wrapper>
  );
};

export default Footer;

const Wrapper = styled.div`
  footer {
    bottom: 0;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    align-items: center;
    font-family: Roboto, sans-serif;
    color: white;
  }
  .body {
    height: 437px;
    background-color: black;
  }
  .footer-copyright {
    margin-left: 41px;
    color: white;
    font-weight: 100;
    font-family: Roboto, sans-serif;
    color: rgba(255, 255, 255, 0.6);
    font-size: 14px;
  }
  .footer-div-1 {
    display: flex;
    flex-direction: column;
    justify-content: center;
    height: 281px;
    width: 269px;
  }

  .footer-div-2 {
    display: flex;
    flex-direction: column;
    justify-content: center;
    width: 226px;
    height: 386px;
  }

  .footer-div-3 {
    display: flex;
    flex-direction: column;
    justify-content: center;
    width: 427px;
    height: 386px;
  }
  .footer-logo {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }
  .footer-logo-heading {
    font-size: 48px;
    margin: 0px;
    font-weight: 500;
  }

  .footer-logo-content {
    font-size: 18px;
    margin: 0;
    font-weight: 300;
    color: rgba(255, 255, 255, 0.6);
  }

  .footer-terms-link {
    font-weight: 300;
    font-size: 14px;
    line-height: 200%;
    text-decoration: none;
    color: rgba(255, 255, 255, 0.6);
  }
  .footer-terms-link:hover {
    text-decoration: underline;
  }
  .footer-terms {
    width: 226px;
    height: 163px;
  }
  .footer-terms-content {
    font-weight: 300;
    font-size: 14px;
    line-height: 200%;

    color: rgba(255, 255, 255, 0.6);
  }
  .footer-terms-heading {
    margin-bottom: 5px;
  }
  .footer-about {
    width: 427px;
    height: 163px;
  }
  .footer-about-content {
    font-weight: 300;
    font-size: 14px;
    color: rgba(255, 255, 255, 0.6);
    line-height: 200%;
  }
  .footer-about-heading {
    margin-bottom: 5px;
  }

  .footer-socials {
    width: 268px;
    height: 80px;
  }

  .footer-socials-heading {
    margin-top: 0px;
    margin-bottom: 5px;
  }

  .footer-built-link {
    font-weight: 300;
    font-size: 14px;
    line-height: 200%;
    text-decoration: none;
    color: rgba(255, 255, 255, 0.6);
  }
  .footer-built-link:hover {
    text-decoration: underline;
  }
  .footer-built-content {
    font-weight: 300;
    font-size: 14px;
    color: rgba(255, 255, 255, 0.6);
    line-height: 200%;
  }
  .footer-built-heading {
    margin-bottom: 5px;
  }
  .footer-contact-content {
    font-weight: 300;
    font-size: 14px;
    color: rgba(255, 255, 255, 0.6);
    line-height: 200%;
  }
  .footer-contact-heading {
    margin-bottom: 5px;
  }
  .footer-socials-icons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    text-decoration: none;
    color: white;
  }

  .social-link {
    text-decoration: none;
    cursor: pointer;
    color: white;
  }
  .footer-copyright-small {
    margin-left: 41px;
    margin-top: 120px;
    color: white;
    font-weight: 100;
    font-family: Roboto, sans-serif;
    color: rgba(255, 255, 255, 0.6);
    font-size: 14px;
  }
`;
