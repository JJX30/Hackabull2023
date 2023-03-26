import React, { useRef, useState, useEffect } from "react";
import styled from "styled-components";

const Signup = () => {
  const [user, setUser] = useState({ email: "", password: "" });
  const [confirm, setConfirm] = useState("");
  const errorMessage = useRef(null);
  const passwordReq = useRef(null);
  const passwordReqNum = useRef(null);
  const passwordReqLen = useRef(null);
  const password = useRef(null);

  useEffect(() => {
    // add when mounted
    document.addEventListener("mousedown", handleClickoutside);
    // return function to be called when unmounted
    return () => {
      document.removeEventListener("mousedown", handleClickoutside);
    };
  }, []);
  const handleClickoutside = (e) => {
    if (password.current.contains(e.target)) {
      passwordReq.current.style.display = "block";
      return;
    }
    // outside click
    passwordReq.current.style.display = "none";
  };

  useEffect(() => {
    if (user.password.length >= 8) {
      passwordReqLen.current.style.color = "green";
    } else if (user.password.length < 8) {
      passwordReqLen.current.style.color = "red";
    }
    if (/\d/.test(user.password)) {
      passwordReqNum.current.style.color = "green";
    } else if (!/\d/.test(user.password)) {
      passwordReqNum.current.style.color = "red";
    }
  }, [user]);
  const handleChange = (e) => {
    errorMessage.current.hidden = true;
    const name = e.target.name;
    const value = e.target.value;

    setUser({ ...user, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (user.email && user.password && confirm) {
      if (
        passwordReqLen.current.style.color === "red" ||
        passwordReqNum.current.style.color === "red"
      ) {
        setUser({ email: user.email, password: "" });
        setConfirm("");
        errorMessage.current.innerHTML = "Password does not meet requirements";
        errorMessage.current.hidden = false;
      } else if (user.password !== confirm) {
        setUser({ email: user.email, password: "" });
        setConfirm("");
        errorMessage.current.innerHTML = "Passwords do not match";
        errorMessage.current.hidden = false;
      } else {
        const newUser = { email: user.email, password: user.password };
        const url = "";

        const options = {
          method: "POST",
          mode: "cors",
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json;charset=UTF-8",
          },
          body: JSON.stringify(newUser),
        };

        try {
          const response = await fetch(url, options);
          const result = await response.json();
          if (result.status === 404) {
            setUser({ email: user.email, password: "" });
            setConfirm("");
            errorMessage.current.innerHTML = "Email already in use";
            errorMessage.current.hidden = false;
          } else if (result.status === 200) {
            alert("Account Succesfully created!");
            // send to
          }
        } catch (err) {
          console.log(err);
        }
      }
    } else {
      setUser({ email: "", password: "" });
      setConfirm("");
      errorMessage.current.innerHTML = "You're missing something";
      errorMessage.current.hidden = false;
    }
  };

  return (
    <Wrapper>
      <div className="body">
        <h1>Get started for free!</h1>
        <p className="subtitle">and improve abilites</p>
        <form onSubmit={handleSubmit}>
          <div className="signup-inputs">
            <input
              type="email"
              placeholder="email"
              className="signup-input"
              name="email"
              value={user.email}
              onChange={handleChange}
            />
            <input
              ref={password}
              type="password"
              placeholder="password"
              className="signup-input"
              name="password"
              value={user.password}
              onChange={handleChange}
            />
            <div ref={passwordReq} className="password-reqs">
              <ul className="password-reqs-list">
                <li ref={passwordReqLen} className="password-reqs-item-length">
                  At least 8 characters
                </li>
                <li ref={passwordReqNum} className="password-reqs-item-num">
                  At least one number
                </li>
              </ul>
            </div>
            <input
              type="password"
              placeholder="confirm password"
              className="signup-input"
              name="confirm"
              value={confirm}
              onChange={(e) => {
                setConfirm(e.target.value);
              }}
            />
          </div>
          <div className="signup-button">
            <button type="submit">create account</button>
            <p hidden ref={errorMessage}>
              Please enter an email and password
            </p>
          </div>
        </form>
      </div>
    </Wrapper>
  );
};

export default Signup;

const Wrapper = styled.div`
  .password-reqs {
    display: none;
    margin: 5px;
    margin-left: 35px;
    margin-bottom: 20px;
  }
  .password-reqs-item-length {
    color: red;
  }
  .password-reqs-item-num {
    color: red;
  }
  .body {
    width: 387px;
    height: 417px;
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
  }

  p {
    color: red;
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
    font-size: 44px;
    font-weight: 500;
    margin-top: 0;
    margin-bottom: 0;
  }
  .signup-inputs {
    display: flex;
    flex-direction: column;
    width: 387px;
  }

  input::placeholder {
    font-size: 18px;
  }

  .signup-input {
    font-size: 18px;
    padding-inline-start: 30px;
    padding-inline-end: 30px;
    width: 387px;
    height: 46px;
    border-radius: 50px;
    border-color: black;
    border-width: 1px;
    box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
    margin-bottom: 20px; /* change */
  }
  form {
    display: flex;
    flex-direction: column;
    justify-content: center;
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
  .signup-button {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
  }

  @media only screen and (max-width: 400px) {
    h1 {
      font-family: Roboto, sans-serif;
      font-size: 40px;
      font-weight: 500;
      margin-top: 0;
      margin-bottom: 0;
    }
    .body {
      width: 300px;
      height: 417px;
      display: flex;
      flex-direction: column;
      justify-content: space-evenly;
    }
    .subtitle {
      font-family: Roboto, sans-serif;
      font-size: 25px;
      font-weight: 300;
      color: rgba(0, 0, 0, 0.8);
      margin-bottom: 29px;
      margin-top: 10px;
    }
    .signup-inputs {
      display: flex;
      flex-direction: column;
      width: 300px;
    }
    .signup-input {
      font-size: 18px;
      padding-inline-start: 30px;
      padding-inline-end: 30px;
      width: 300px;
      height: 46px;
      border-radius: 50px;
      border-color: black;
      border-width: 1px;
      box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
      margin-bottom: 20px; /* change */
    }
  }
  @media only screen and (max-width: 300px) {
    h1 {
      font-family: Roboto, sans-serif;
      font-size: 38px;
      font-weight: 500;
      margin-top: 0;
      margin-bottom: 0;
    }
    .body {
      width: 250px;
      height: 417px;
      display: flex;
      flex-direction: column;
      justify-content: space-evenly;
    }
    .subtitle {
      font-family: Roboto, sans-serif;
      font-size: 20px;
      font-weight: 300;
      color: rgba(0, 0, 0, 0.8);
      margin-bottom: 29px;
      margin-top: 10px;
    }
    .signup-inputs {
      display: flex;
      flex-direction: column;
      width: 250px;
    }
    .signup-input {
      font-size: 18px;
      padding-inline-start: 30px;
      padding-inline-end: 30px;
      width: 250px;
      height: 46px;
      border-radius: 50px;
      border-color: black;
      border-width: 1px;
      box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
      margin-bottom: 20px; /* change */
    }
  }
`;
