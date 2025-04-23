import React from 'react';
import { Modal, Button, message } from 'antd';
import { SurveyResponse } from '../types/Survey';
import { base64ToBlob } from '../utils/encoding';
import { deleteSurvey } from '../api/SurveyAPI';

interface SurveyOwnerModalProps {
    survey: SurveyResponse;
    onClose: () => void;
    onDeleteSurvey: (surveyId: string) => void; // New prop for deleting the survey
}

const SurveyOwnerModal: React.FC<SurveyOwnerModalProps> = ({
    survey,
    onClose,
    onDeleteSurvey, // Receive the delete function
}) => {

    const handleDelete = () => {
        Modal.confirm({
            title: 'Are you sure you want to delete this survey?',
            onOk: async () => {
                await deleteSurvey(survey.id);
                onDeleteSurvey(survey.id);
                onClose();
                message.success('Survey deleted successfully');
            },
            onCancel: () => { },
        });
    };

    return (
        <Modal
            title={`Survey Owner View: ${survey.name}`}
            open={true}
            onCancel={onClose}
            footer={[
                <Button key="back" onClick={onClose}>
                    Close
                </Button>,
                <Button key="delete" danger onClick={handleDelete}>
                    Delete Survey
                </Button>,
            ]}
            width={800}
        >
            <p>Total questions: {survey.questions?.length || 0}</p>

            {survey.questions?.map((question, index) => (
                <div key={index} style={{ marginBottom: '1rem' }}>
                    <strong>Q{index + 1}: {question.question}</strong>
                    <ul style={{ paddingLeft: '1.5rem' }}>
                        {(question.answers?.length ?? 0) === 0 ? (
                            <li>No answers yet</li>
                        ) : (
                            question.answers?.map((ans, ansIndex) => (
                                <li key={ansIndex}>
                                    <audio controls>
                                        <source src={URL.createObjectURL(base64ToBlob(ans.answer, "audio/webm"))} type="audio/webm" />
                                        Your browser does not support the audio element.
                                    </audio>
                                </li>
                            ))
                        )}
                    </ul>
                </div>
            ))}
        </Modal>
    );
};

export default SurveyOwnerModal;
