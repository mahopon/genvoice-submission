import React from "react";
import { Modal, Input, Form, Select, message } from "antd";
import { UserCreateRequest } from "../types/User";
import { createUser } from "../api/UserAPI";
import { ApiError } from "../utils/APIerror";

const { Option } = Select;

interface AddUserModalProps {
    visible: boolean;
    onCancel: () => void;
    onSubmit: () => void;
}

const AddUserModal: React.FC<AddUserModalProps> = ({ visible, onCancel, onSubmit }) => {
    const [form] = Form.useForm();
    const [messageApi, contextHolder] = message.useMessage();

    const handleOk = () => {
        form.validateFields().then(async (values) => {
            const newUser: UserCreateRequest = {
                name: values.name,
                username: values.username,
                password: values.password,
                role: values.role,
            };
            try {
                await createUser(newUser);
                onSubmit();
                form.resetFields();
            } catch (err) {
                const castedErr = err as ApiError;
                messageApi.error(`Unable to register. ${castedErr}`);
            }
        });
    };

    return (
        <Modal
            title="Add New User"
            open={visible}
            onOk={handleOk}
            onCancel={() => {
                onCancel();
                form.resetFields();
            }}
        >
            {contextHolder}
            <Form form={form} layout="vertical" initialValues={{ role: "USER" }} >
                <Form.Item name="name" label="Name" rules={[{ required: true }]}>
                    <Input />
                </Form.Item>
                <Form.Item name="username" label="Username" rules={[{ required: true }]}>
                    <Input />
                </Form.Item>
                <Form.Item name="password" label="Password" rules={[{ required: true }]}>
                    <Input.Password />
                </Form.Item>
                <Form.Item name="role" label="Role" rules={[{ required: true }]}>
                    <Select>
                        <Option value="USER">USER</Option>
                        <Option value="ADMIN">ADMIN</Option>
                    </Select>
                </Form.Item>
            </Form>
        </Modal>
    );
};

export default AddUserModal;