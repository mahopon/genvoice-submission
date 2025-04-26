import React from 'react';
import { Form, Input, Button, message } from 'antd';
import { PasswordChangeRequest } from '../types/User';
import { updatePassword } from '../api/UserAPI';
import { useAuth } from '../context/AuthContext';
import { ApiError } from '../utils/APIerror';

const SettingsForm: React.FC = () => {
    const { userId } = useAuth();
    const [form] = Form.useForm();

    interface FormValues {
        currentPassword: string;
        newPassword: string;
    }

    const onFinish = async (values: FormValues) => {
        const req: PasswordChangeRequest = {
            current_password: values.currentPassword,
            new_password: values.newPassword
        };
        try {
            await updatePassword(userId!, req);
            message.success("Password updated successfully");
            form.resetFields();
        } catch (err) {
            const castedErr = err as ApiError;
            message.error(`Couldn't update password. ${castedErr.message}`);
        }
    };

    return (
        <div style={{ display: 'flex', justifyContent: 'center', minHeight: '80vh' }}>
            <Form
                form={form}
                layout="vertical"
                onFinish={onFinish}
                style={{ width: '100%', maxWidth: 400 }}
            >
                <Form.Item
                    label="Current Password"
                    name="currentPassword"
                    rules={[{ required: true, message: 'Please enter your current password' }]}
                >
                    <Input.Password />
                </Form.Item>

                <Form.Item
                    label="New Password"
                    name="newPassword"
                    rules={[{ required: true, message: 'Please enter your new password' }]}
                >
                    <Input.Password />
                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
                        Change Password
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
};

export default SettingsForm;
