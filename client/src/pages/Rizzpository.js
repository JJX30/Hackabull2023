import React from "react";
import styled from "styled-components";
import { RiArrowDropDownLine } from "react-icons/ri";

const Rizzpository = () => {
  return (
    <Wrapper>
      <div className="doc">
        <div className="header">
          <h1 className="header-text">Rizzpository</h1>
          <div className="divider"></div>
        </div>
        <div className="report">
          <div className="report-heading">
            <h1 className="report-text">Report 1</h1>
            <RiArrowDropDownLine size={60}></RiArrowDropDownLine>
          </div>
          <p className="report-sub">Rizz confidence level: 95.677777%</p>
          <div className="textarea">
            <p className="report-sub">Captured text:</p>
            <p className="Text">
              {" "}
              <mark clasName="highlight">You need to</mark> hire me because{" "}
              <mark clasName="highlight">I'm the best candidate </mark>
              for the job
            </p>
          </div>
        </div>
      </div>
    </Wrapper>
  );
};

export default Rizzpository;

const Wrapper = styled.div`
  font-family: Roboto, sans-serif;
  .doc {
    display: flex;
    flex-direction: column;
    align-items: left;
    width: 100%;
    margin-left: 200px;
    margin-right: 200px;
  }
  .report-heading {
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  .header-text {
    font-size: 65px;
  }
  .report-text {
    font-size: 40px;
  }
  .report-sub {
    font-size: 23px;
  }
  .divider {
    width: 70%;
    height: 1px;
    background-color: black;
  }
  .header {
    margin-bottom: 30px;
  }

  .Text {
    font-size: 20px;
    font-weight: 300;
  }
  .report {
    margin-bottom: 50px;
  }
  .highlight {
    background-color: green;
    /* background-color: #7cff3b; */
  }
`;
