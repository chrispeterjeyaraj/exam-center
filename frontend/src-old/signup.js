import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { fetchData } from "../../../helpers/api";
import "./signup.css";

const SignUp = () => {
  const [firstname, setFirstname] = useState("");
  const [lastname, setLastname] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const { t } = useTranslation();

  const handleFirstNameChange = (e) => {
    setFirstname(e.target.value);
  };
  const handleLastNameChange = (e) => {
    setLastname(e.target.value);
  };
  const handleEmailChange = (e) => {
    setEmail(e.target.value);
  };
  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };
  const navigate = useNavigate();
  const loginUser = () => {
    if (!email || !password) {
      alert(t("login.required"));
    } else {
      fetchData("signup", "POST", {
        "Firstname": firstname,
        "Lastname": lastname,
        "Email": email,
        "Password": password,
      })
        .then((data) => data.json())
        .then((data) => {
          if (data.Status !== "Success") {
            alert(t("login.invalid.credentials"));
          } else {
            alert(t("login.signup.success"));
            navigate(`/`);
          }
        })
        .catch((err) => alert(t("login.nonetwork"), err));
    }
  };

  return (
    <>
      <div className="cont">
        <div className="form sign-in">
          <h2>{t("login.signup.welcome")}</h2>
          <label>
            <span>{t("login.signup.firstname")}</span>
            <input type="text" onChange={handleFirstNameChange} />
          </label>
          <label>
            <span>{t("login.signup.lastname")}</span>
            <input type="text" onChange={handleLastNameChange} />
          </label>
          <label>
            <span>{t("login.email")}</span>
            <input type="text" onChange={handleEmailChange} />
          </label>
          <label>
            <span>{t("login.password")}</span>
            <input type="password" onChange={handlePasswordChange} />
          </label>
          <button type="button" className="submit" onClick={loginUser}>{t("login.signup")}</button>
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

export default SignUp;
