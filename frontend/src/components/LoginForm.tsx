import React, { useState } from 'react';
import { Form, Input, Button, message } from 'antd';
import { loginUser } from '../api/UserAPI'
import { UserLoginRequest } from '../types/User';
import { useNavigate } from 'react-router';
import { useAuth } from '../context/AuthContext';

interface FormValues {
  username: string;
  password: string;
}

const LoginForm: React.FC = () => {
  const { setAuthStatus } = useAuth();
  const [messageApi, contextHolder] = message.useMessage();
  const [submitted, setSubmitted] = useState<boolean>(false);
  const navigate = useNavigate();
  const onFinish = async (values: FormValues) => {
    setSubmitted(true);
    console.log('Received values:', values);
    const details: UserLoginRequest = {
      username: values.username,
      password: values.password
    };
    if (await loginUser(details)) {
      messageApi.success('Successfully logged in!');
      setTimeout(() => {
        setAuthStatus(true);
        navigate("/");
      }, 1000);
    } else {
      setSubmitted(false);
      messageApi.error('Login failed. Please try again.');
    }
  };

  return (
    <div style={{ width: '300px', margin: '0 auto' }}>
      {contextHolder}
      <Form
        name="login_form"
        onFinish={onFinish}
        initialValues={{ remember: true }}
        layout="vertical"
      >
        <Form.Item
          name="username"
          label="Username"
          rules={[{ required: true, message: 'Please input your username!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="password"
          label="Password"
          rules={[{ required: true, message: 'Please input your password!' }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit" style={{ width: '100%' }} disabled={submitted}>
            {!submitted ? ("Log in") : ("Logging in...")}
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default LoginForm;
