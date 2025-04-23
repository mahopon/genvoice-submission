import React, { useState } from 'react';
import { Button } from 'antd';
import AddSurveyModal from './AddSurveyModal';

interface Props {
    refresh: () => void;
}

const AddSurveyButton: React.FC<Props> = ({ refresh }) => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const handleClose = () => {
        refresh();
        setIsModalOpen(false);
    };

    return (
        <>
            <Button type="primary" onClick={() => setIsModalOpen(true)}>
                Add Survey
            </Button>
            <AddSurveyModal open={isModalOpen} onClose={handleClose} />
        </>
    );
};

export default AddSurveyButton;
