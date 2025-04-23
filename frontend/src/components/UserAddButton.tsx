import React, { useState } from "react";
import { Button } from "antd";
import AddUserModal from "./UserAddModal";

interface UserAddButtonProps {
    onAdd: () => void;
}

const UserAddButton: React.FC<UserAddButtonProps> = ({ onAdd }) => {
    const [modalVisible, setModalVisible] = useState(false);

    const handleSubmit = async () => {
        onAdd();
        setModalVisible(false);
    };

    return (
        <>
            <Button type="primary" onClick={() => setModalVisible(true)} style={{ float: "right", marginBottom: 16 }}>
                Add User
            </Button>
            <AddUserModal
                visible={modalVisible}
                onCancel={() => setModalVisible(false)}
                onSubmit={handleSubmit}
            />
        </>
    );
};

export default UserAddButton;
