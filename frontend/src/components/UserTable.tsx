import React, { useState, useEffect } from "react";
import { Table, Button, Input, Select, message } from "antd";
import { ColumnsType } from "antd/es/table";
import { UserDataRequest, UserUpdateRequest } from "../types/User";
import { deleteUser, editWholeUser, getAllUser } from "../api/UserAPI";
import { useAuth } from "../context/AuthContext";
import AddUserButton from "./UserAddButton";
const { Option } = Select;

const UserTable: React.FC = () => {
    const [data, setData] = useState<UserDataRequest[]>([]);
    const [editingKey, setEditingKey] = useState<string>("");
    const [editedRow, setEditedRow] = useState<UserDataRequest | null>(null);
    const [messageApi, contextHolder] = message.useMessage();
    const { userId } = useAuth();

    const fetchData = async () => {
        const users = await getAllUser();
        setData(() => {
            return users;
        });
    };
    useEffect(() => {
        fetchData();
    }, []);

    const isEditing = (record: UserDataRequest) => record.id === editingKey;

    const edit = (record: UserDataRequest) => {
        setEditingKey(record.id);
        setEditedRow({ ...record });
    };

    const onAdd = () => {
        messageApi.success("Added user");
        fetchData();
    }

    const save = async () => {
        if (editedRow) {
            if (!editedRow.name || !editedRow.username) {
                messageApi.error('Name and Username are required');
                return;
            }
            const newData = [...data];
            const index = newData.findIndex((item) => item.id === editingKey);
            if (index > -1) {
                newData[index] = { ...newData[index], ...editedRow };
                setData(newData);
            }

            const update: UserUpdateRequest = {
                name: editedRow.name,
                username: editedRow.username,
                password: editedRow.password,
                role: editedRow.role.toUpperCase()
            }

            const success = await editWholeUser(editingKey, update);

            if (success) {
                messageApi.success('User updated successfully');
            } else {
                messageApi.error('Failed to update user');
            }
        }
        setEditingKey("");
        setEditedRow(null);
    };

    const deleteRow = async (id: string) => {
        const deleted = await deleteUser(id);
        if (deleted) {
            messageApi.success("Deleted user");
            setData(data.filter((item) => item.id !== id));
        }
    };

    const cancel = () => {
        setEditingKey("");
        setEditedRow(null);
    };

    const handleChange = (value: string, field: keyof UserDataRequest) => {
        setEditedRow((prev) => prev ? { ...prev, [field]: value } : null);
    };

    const columns: ColumnsType<UserDataRequest> = [
        {
            title: "ID",
            dataIndex: "id",
            key: "id",
        },
        {
            title: "Created Date",
            dataIndex: "created_date",
            key: "created_date",
            render: (text) => {
                const date = new Date(text);
                return date.toLocaleString();
            },
        },
        {
            title: "Name",
            dataIndex: "name",
            key: "name",
            render: (text, record) =>
                isEditing(record) ? (
                    <Input
                        value={editedRow?.name}
                        onChange={(e) => handleChange(e.target.value, "name")}
                        style={{ width: "100%" }}
                    />
                ) : (
                    text
                ),
        },
        {
            title: "Username",
            dataIndex: "username",
            key: "username",
            render: (text, record) =>
                isEditing(record) ? (
                    <Input
                        value={editedRow?.username}
                        onChange={(e) => handleChange(e.target.value, "username")}
                        style={{ width: "100%" }}
                    />
                ) : (
                    text
                ),
        },
        {
            title: "Password",
            key: "password",
            render: (_, record) =>
                isEditing(record) ? (
                    <Input.Password
                        value={editedRow?.password || ''}
                        onChange={(e) => handleChange(e.target.value, "password" as keyof UserDataRequest)}
                        style={{ width: "100%" }}
                        placeholder="Enter new password"
                    />
                ) : (
                    "••••••••"
                ),
        },
        {
            title: "Role",
            dataIndex: "role",
            key: "role",
            render: (text, record) =>
                isEditing(record) ? (
                    record.id !== userId ? (
                        <Select
                            value={editedRow?.role.toUpperCase()}
                            onChange={(value) => handleChange(value.toUpperCase(), "role")}
                            style={{ width: "100%" }}
                        >
                            <Option value="admin">ADMIN</Option>
                            <Option value="user">USER</Option>
                        </Select>
                    ) : (
                        text
                    )
                ) : (
                    text
                ),
        },
        {
            title: "Actions",
            key: "actions",
            render: (_, record) =>
                isEditing(record) ? (
                    <>
                        <Button type="primary" onClick={save} style={{ marginRight: 8 }}>
                            Save
                        </Button>
                        <Button onClick={cancel}>Cancel</Button>
                    </>
                ) : (
                    <>
                        <Button onClick={() => edit(record)} style={{ marginRight: 8 }}>
                            Edit
                        </Button>
                        {record.id != userId && (
                            <Button danger onClick={() => deleteRow(record.id)}>
                                Delete
                            </Button>
                        )}
                    </>
                ),
        },
    ];

    return (
        <>
            {contextHolder}
            <div style={{ display: "flex", justifyContent: "flex-end", marginBottom: 20 }}>
                <AddUserButton onAdd={onAdd} />
            </div>
            <Table columns={columns} dataSource={data} pagination={false} />
        </>
    );
};

export default UserTable;
