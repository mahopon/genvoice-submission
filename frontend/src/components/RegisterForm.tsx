import React from "react";
import { Form, Input, Button, Card } from "antd";
import { registerUser } from "../api/UserAPI";
import { RegisterUserRequest, } from "../types/User";

interface Props {
    setRegistered: (value: boolean) => void;
}


const RegistrationForm: React.FC<Props> = ({ setRegistered }) => {
    const [form] = Form.useForm();

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
        registerUser(newUser);
        setRegistered(true);
    };


    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <Card title="Register" className="rounded-2xl shadow-lg" style={{ maxWidth: 300, width: "100%" }}>
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

                    <Form.Item>
                        <Button type="primary" htmlType="submit">
                            Register
                        </Button>
                    </Form.Item>
                </Form>
            </Card>
        </div>
    );
};

export default RegistrationForm;
