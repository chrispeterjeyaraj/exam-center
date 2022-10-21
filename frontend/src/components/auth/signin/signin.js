import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { fetchData } from "../../../helpers/api";
import "./signin.css";

const SignIn = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { t } = useTranslation();

  const handleInputChange = (e) => {
    setEmail(e.target.value);
  };
  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };
  // const handleSignUpClick = (e) => {
  //   navigate(`/signup`);
  // }

  const navigate = useNavigate();
  const loginUser = () => {
    if (!email || !password) {
      alert(t("login.required"));
    } else {
      fetchData("signin", "POST", {
        "Email": email,
        "Password": password,
      })
        .then((data) => data.json())
        .then((data) => {
          console.log("user received is", data);
          if (data.status === "invalid") {
            alert(t("login.invalid.credentials"));
          } else {
            localStorage.setItem("sessionId", data.sessionId);
            navigate(`/exams`);
          }
        })
        .catch((err) => alert(t("login.nonetwork"), err));
    }
  };

  return (
    <>
      <div className="cont">
        <div className="form sign-in">
          <h2>{t("login.welcome")}</h2>
          <label>
            <span>{t("login.email")}</span>
            <input type="text" onChange={handleInputChange} />
          </label>
          <label>
            <span>{t("login.password")}</span>
            <input type="password" onChange={handlePasswordChange} />
          </label>
          <button type="button" className="submit" onClick={loginUser}>{t("login.signin")}</button>
          <a href="/signup">{ t("login.signin.signup") }</a>
        </div>
        <div className="sub-cont">
          <div className="img">
            <div className="img__text m--up">
              <p>{t("login.quote")}</p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default SignIn;
