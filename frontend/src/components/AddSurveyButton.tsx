import React, { useState } from 'react';
import { Button } from 'antd';
import AddSurveyModal from './AddSurveyModal';

const AddSurveyButton: React.FC = () => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const handleClose = () => {
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
