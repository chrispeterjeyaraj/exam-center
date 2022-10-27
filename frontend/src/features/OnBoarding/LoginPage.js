import {useDispatch} from 'react-redux'
import {useSelector} from 'react-redux'
import {Form, Input, Button, Checkbox} from 'antd';
import {UserOutlined, LockOutlined} from '@ant-design/icons';
import { login } from './authenticationSlice'
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from 'react';
import { useTranslation } from "react-i18next";


function LoginPage() {
  const loader = useSelector(state => state.authentication.loader)
  const dispatch = useDispatch();
  const isAuthenticated = useSelector((state) => state.authentication.isAuthenticated)
  const navigate = useNavigate();
  const { t } = useTranslation();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleInputChange = (e) => {
    setEmail(e.target.value);
  };
  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  useEffect(() => {
    if (isAuthenticated) {
      console.log('login', isAuthenticated)
      navigate("/dashboard")
    }
  }, [isAuthenticated])

  return (
    <div className="container">
      <Form
        name="normal_login"
        className="form"
        initialValues={{
          remember: true,
        }}
        onFinish={() => dispatch(login(
          {'email': email, 'password': password},
        ))
        }
      >
        <Form.Item
          name="email"
          rules={[
            {
              required: true,
              message: 'Please input your email!',
            },
          ]}
        >
          <Input size="large"
            prefix={<UserOutlined className="site-form-item-icon"/>}
            placeholder={t("login.email")}
            autoComplete="email"
            onChange={handleInputChange}
          />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[
            {
              required: true,
              message: 'Please input your Password!',
            },
          ]}
        >
          <Input
            prefix={<LockOutlined className="site-form-item-icon"/>}
            type="password"
            placeholder={t("login.password")}
            size="large"
            autoComplete="current-password"
            onChange={handlePasswordChange}
          />
        </Form.Item>
        <Form.Item>
          <Form.Item name="remember" valuePropName="checked" noStyle>
            <Checkbox>{t("login.rememberme")}</Checkbox>
          </Form.Item>
        </Form.Item>

        <Form.Item>
          <Button loading={loader} type="primary" htmlType="submit" className="login-form-button"
            size="large">{t("login.signin")}
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}

export default LoginPage;
