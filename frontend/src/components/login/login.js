import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { fetchData } from "../../helpers/api";
import "./login.css";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const { translation } = useTranslation();
  const handleInputChange = (e) => {
    setUsername(e.target.value);
  };
  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };
  const navigate = useNavigate();
  const loginUser = () => {
    if (!username || !password) {
      alert(translation("login.required"));
    } else {
      fetchData("fetchUser", "POST", {
        username,
        password,
      })
        .then((data) => data.json())
        .then((data) => {
          console.log("user received is", data);
          if (data.status == "invalid") {
            alert(translation("login.invalid.credentials"));
          } else {
            localStorage.setItem("sessionId", data.sessionId);
            navigate(`/exams`);
          }
        })
        .catch((err) => alert(translation("login.nonetwork"), err));
    }
  };

  return (
    <>
      <div className="cont">
        <div className="form sign-in">
          <h2>{translation("login.welcome")}</h2>
          <label>
            <span>{translation("login.username")}</span>
            <input type="text" onChange={handleInputChange} />
          </label>
          <label>
            <span>{translation("login.password")}</span>
            <input type="password" onChange={handlePasswordChange} />
          </label>
          <button type="button" className="submit" onClick={loginUser}></button>
        </div>
        <div className="sub-cont">
          <div className="img">
            <div className="img__text m--up">
              <p>{translation("login.quote")}</p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Login;
