import React from "react";
import styled from "styled-components";
import { useState, useRef } from "react";
import { Redirect } from "react-router-dom";
const Signin = () => {
  const [authUser, setAuthUser] = useState({ email: "", password: "" });
  const errorMessage = useRef(null);

  const handleChange = (e) => {
    errorMessage.current.innerHTML = "Please enter an email and password";
    errorMessage.current.hidden = true;
    const name = e.target.name;
    const value = e.target.value;
    setAuthUser({ ...authUser, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (authUser.email && authUser.password) {
      window.alert("Sign in  successful");
      window.location.replace("http://localhost:3000/rizz");
      //   const newAuthUser = {
      //     email: authUser.email,
      //     password: authUser.password,
      //   };
      //   const url = "/api/signin";
      //   const options = {
      //     method: "POST",
      //     mode: "cors",
      //     headers: {
      //       Accept: "application/json",
      //       "Content-Type": "application/json;charset=UTF-8",
      //     },
      //     body: JSON.stringify(newAuthUser),
      //   };
      //   try {
      //     const response = await fetch(url, options);
      //     const result = await response.json();
      //     if (result.error) {
      //       errorMessage.current.innerHTML = "Try signing in again!";
      //       errorMessage.current.hidden = false;
      //     } else {
      //     }
      //   } catch (err) {
      //     console.log(err);
      //   }
    } else {
      errorMessage.current.innerHTML = "Please enter an email and password";
      errorMessage.current.hidden = false;
    }
  };
  return (
    <Wrapper>
      <div className="footer-logo">
        <p className="footer-logo-heading">Rizzerator</p>
        <p className="footer-logo-content">the confidence model</p>
      </div>
      <div className="body">
        <form onSubmit={handleSubmit}>
          <div className="signin-inputs">
            <input
              type="email"
              placeholder="email"
              className="signin-input"
              name="email"
              value={authUser.email}
              onChange={handleChange}
            />
            <input
              type="password"
              placeholder="password"
              className="signin-input"
              name="password"
              value={authUser.password}
              onChange={handleChange}
            />
          </div>
          <div className="signin-button">
            <button type="submit">sign in</button>
            <p hidden ref={errorMessage}>
              Please enter an email and password
            </p>
          </div>
        </form>
      </div>
    </Wrapper>
  );
};

export default Signin;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 100px;
  margin-bottom: 100px;
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
  .body {
    width: 430px;
    height: 300px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-evenly;
  }

  p {
    color: black
    font-size: 14px;
    font-family: Roboto, sans-serif;
    font-weight: 400;
  }
  .subtitle {
    font-family: Roboto, sans-serif;
    font-size: 24px;
    font-weight: 300;
    color: rgba(0, 0, 0, 0.8);
    margin-bottom: 29px;
    margin-top: 10px;
  }
  h1 {
    font-family: Roboto, sans-serif;
    font-size: 48px;
    font-weight: 500;
    margin-top: 0;
    margin-bottom: 0;
  }
  .signin-inputs {
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    width: 387px;
    height: 176px;
  }

  input::placeholder {
    font-size: 18px;
  }

  .signin-input {
    font-size: 18px;
    padding-inline-start: 30px;
    padding-inline-end: 30px;
    width: 387px;
    height: 46px;
    border-radius: 50px;
    border-color: black;
    border-width: 1px;
    box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  }
  form {
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    width: 387px;
    height: 285px;
  }
  button {
    width: 269px;
    height: 70px;
    cursor: pointer;
    border-radius: 50px;
    border-style: none;
    color: white;
    background: rgb(27, 27, 76);
    font-family: Roboto, sans-serif;
    font-size: 24px;
    font-weight: 300;
  }
  .signin-button {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
  }

  @media only screen and (max-width: 430px) {
    .body {
      width: 200px;
      height: 300px;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: space-evenly;
    }
    .signin-inputs {
      display: flex;
      flex-direction: column;
      justify-content: space-evenly;
      width: 300px;
      height: 176px;
    }
    .signin-input {
      font-size: 18px;
      padding-inline-start: 30px;
      padding-inline-end: 30px;
      width: 300px;
      height: 46px;
      border-radius: 50px;
      border-color: black;
      border-width: 1px;
      box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
    }
    form {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: space-evenly;
      width: 300px;
      height: 285px;
    }
  }
  @media only screen and (max-width: 300px) {
    .signin-inputs {
      display: flex;
      flex-direction: column;
      justify-content: space-evenly;
      width: 200px;
      height: 176px;
    }
    .signin-input {
      font-size: 18px;
      padding-inline-start: 30px;
      padding-inline-end: 30px;
      width: 200px;
      height: 40px;
      border-radius: 50px;
      border-color: black;
      border-width: 1px;
      box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
    }
    form {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: space-evenly;
      width: 200px;
      height: 285px;
    }
    button {
      width: 180px;
      height: 70px;
      cursor: pointer;
      border-radius: 50px;
      border-style: none;
      color: white;
      background: rgb(185, 1, 1);
      background: linear-gradient(
        360deg,
        rgba(185, 1, 1, 1) 0%,
        rgba(219, 0, 0, 1) 100%
      );
      font-family: Roboto, sans-serif;
      font-size: 24px;
      font-weight: 300;
    }
  }
`;
