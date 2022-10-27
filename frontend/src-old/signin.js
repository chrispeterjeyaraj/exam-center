import React, { useState } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { useAuth } from "../../../hooks/useAuth";
import "../../../shared-styles.css";

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

  const auth = useAuth();
  const loginUser = () => {
    if (!email || !password) {
      alert(t("login.required"));
    } else {
      auth.login({ email: email, password: password });
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
