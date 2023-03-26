import React, { useState } from "react";
import styled from "styled-components";
import { ReactMediaRecorder } from "react-media-recorder";
import { BiMicrophone, BiMicrophoneOff } from "react-icons/bi";
import { useNavigate } from "react-router-dom";

const Home = () => {
  const [recording, setRecording] = useState(false);
  const navigate = useNavigate();
  const handleSubmit = async (media) => {
    setTimeout(function () {
      navigate("/Rizzpository", { replace: true });
    }, 5000);

    // console.log(blob);
    // const formData = new FormData();
    // formData.append("file", blob, blob.name);
    // const url = "";
    // const options = {
    //   method: "POST",
    //   mode: "cors",
    //   headers: {
    //     Accept: "audio/wav",
    //     "Content-Type": "audio/wav",
    //   },
    //   body: formData,
    // };

    // fetch(url, options)
    //   .then((response) => {
    //     // Handle the response from the server
    //     if (response.ok) {
    //       console.log("AUDIO SENT");
    //     }
    //   })
    //   .catch((error) => {
    //     // Handle any errors that occur during the request
    //     if (error) {
    //       console.log("ERRRRRRORRRRRR");
    //     }
    //   });
  };
  return (
    <Wrapper>
      <div className="main">
        <div>
          <h1 className="title">Rizz it up</h1>
          <p className="footer-logo-content">and get a rizz-report</p>
        </div>
        <ReactMediaRecorder
          audio
          render={({
            status,
            startRecording,
            stopRecording,
            mediaBlobUrl,
            clearBlobUrl,
          }) => (
            <div className="recorder">
              <button
                onClick={() => {
                  if (recording) {
                    stopRecording();
                    setRecording(false);
                  } else {
                    startRecording();
                    setRecording(true);
                  }
                }}
                className="button"
              >
                {recording ? (
                  <BiMicrophoneOff size={30}></BiMicrophoneOff>
                ) : (
                  <BiMicrophone size={30}></BiMicrophone>
                )}
              </button>
              {mediaBlobUrl ? (
                <>
                  <audio className="mediablob" src={mediaBlobUrl} controls />
                  <div className="options">
                    <button
                      className="butt"
                      onClick={() => {
                        handleSubmit(mediaBlobUrl);
                      }}
                    >
                      Rizzerate
                    </button>
                    <button className="butt" onClick={clearBlobUrl}>
                      Redo
                    </button>
                  </div>
                </>
              ) : (
                <div></div>
              )}
            </div>
          )}
        />
      </div>
    </Wrapper>
  );
};

export default Home;
const Wrapper = styled.div`
  .main {
    display: flex;
    flex-direction: column;
    margin-top: 100px;
    margin-bottom: 100px;
    align-items: center;
  }
  .recorder {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .button {
    width: 100px;
    height: 100px;
    background: none;
    border-radius: 50px;
    cursor: pointer;
  }
  .button:hover {
    background-color: #dedede;
  }
  .mediablob {
    margin: 20px;
  }
  .options {
    display: flex;
    justify-content: space-evenly;
    width: 50%;
  }
  .title {
    font-size: 60px;
    font-family: Roboto, sans-serif;
    margin-bottom: 0px;
  }

  .footer-logo-content {
    font-family: Roboto, sans-serif;

    font-size: 25px;
    margin: 0;
    font-weight: 300;
    opacity: 70%;
    margin-bottom: 30px;
  }
`;
