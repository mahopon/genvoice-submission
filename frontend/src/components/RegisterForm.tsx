import React from "react";
import { Form, Input, Button, message } from "antd";
import { loginUser, registerUser } from "../api/UserAPI";
import { RegisterUserRequest, UserLoginRequest, UserLoginResponse, } from "../types/User";
import { useAuth } from "../context/AuthContext";
import { ApiError } from "../utils/APIerror";

interface Props {
    setRegistered: (value: boolean) => void;
}


const RegistrationForm: React.FC<Props> = ({ setRegistered }) => {
    const [form] = Form.useForm();
    const [messageApi, contextHolder] = message.useMessage();

    const { setAuthStatus, setRole, setUserId } = useAuth();
    const onFinish = async (values: RegisterUserRequest) => {
        console.log("Form Submitted:", values);
        const name = values.name;
        const username = values.username
        const password = values.password

        if (!name || !username || !password) {
            console.error("All fields are required!");
            return;
        }
        const newUser: RegisterUserRequest = {
            name: name,
            username: username,
            password: password,
        };

        try {
            await registerUser(newUser);
            try {
                const login: UserLoginRequest = {
                    username: newUser.username,
                    password: newUser.password
                }
                const res: UserLoginResponse = await loginUser(login);
                setAuthStatus(true);
                setRole(res.role);
                setUserId(res.id);
                messageApi.success("Registered and logged in");
                setRegistered(true);
            } catch (err) {
                const castedErr = err as ApiError;
                messageApi.error(`Unable to login. ${castedErr.message}`);
            }
        } catch (err) {
            const castedErr = err as ApiError;
            messageApi.error(`Unable to register. ${castedErr.message}`);
        }
    };


    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100" style={{ width: '300px', margin: '0 auto', alignItems: "center" }}>
            {contextHolder}
            <Form
                form={form}
                layout="vertical"
                onFinish={onFinish}
                autoComplete="off"
            >
                <Form.Item
                    label="Name"
                    name="name"
                    rules={[{ required: true, message: "Please input your name!" }]}
                >
                    <Input placeholder="Enter your name" />
                </Form.Item>

                <Form.Item
                    label="Username"
                    name="username"
                    rules={[{ required: true, message: "Please input your username!" }]}
                >
                    <Input placeholder="Choose a username" />
                </Form.Item>

                <Form.Item
                    label="Password"
                    name="password"
                    rules={[{ required: true, message: "Please input your password!" }]}
                >
                    <Input.Password placeholder="Enter a password" />
                </Form.Item>

                <Form.Item style={{ textAlign: "center" }}>
                    <Button type="primary" htmlType="submit" style={{ width: "100%" }}>
                        Register
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
};

export default RegistrationForm;
