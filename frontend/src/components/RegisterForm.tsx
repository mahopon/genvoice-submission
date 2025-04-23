import React from "react";
import { Form, Input, Button, Card, message } from "antd";
import { loginUser, registerUser } from "../api/UserAPI";
import { RegisterUserRequest, UserLoginRequest, UserLoginResponse, } from "../types/User";
import { useAuth } from "../context/AuthContext";

interface Props {
    setRegistered: (value: boolean) => void;
}


const RegistrationForm: React.FC<Props> = ({ setRegistered }) => {
    const [form] = Form.useForm();
    const [messageApi, contextHolder] = message.useMessage();

    const { setAuthStatus, setRole, setUserId } = useAuth();
    const onFinish = (values: RegisterUserRequest) => {
        console.log("Form Submitted:", values);
        const name = values.name;
        const username = values.username
        const password = values.password

        if (!name || !username || !password) {
            console.error("All fields are required!");
            return;
        }
        const newUser: RegisterUserRequest = {
            name: name as string,
            username: username as string,
            password: password as string,
        };
        registerUser(newUser).then(() => {
            const login: UserLoginRequest = {
                username: newUser.username,
                password: newUser.password
            }
            loginUser(login).then((res: UserLoginResponse | false) => {
                setAuthStatus(true);
                setRole((res as UserLoginResponse).role);
                setUserId((res as UserLoginResponse).id);
                messageApi.success("Registered and logged in");
            })
            setRegistered(true);
        }).catch((err) => {
            messageApi.error(`Unable to register. ${err.response.data.error}`);
        });
    };


    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            {contextHolder}
            <Card className="rounded-2xl shadow-lg" style={{ maxWidth: "300px", width: "100%", margin: "0 auto" }}>
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
            </Card>
        </div>
    );
};

export default RegistrationForm;
