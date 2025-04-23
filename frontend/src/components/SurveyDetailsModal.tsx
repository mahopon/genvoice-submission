import React from 'react';
import { Modal } from 'antd';
import { SurveyResponse } from '../types/Survey';
import VoiceRecorder from './VoiceRecorder';

interface SurveyDetailsModalProps {
    selectedSurvey: SurveyResponse | null;
    onOk: () => void;
    onClose: () => void;
}

const SurveyDetailsModal: React.FC<SurveyDetailsModalProps> = ({ selectedSurvey, onOk, onClose }) => {
    return (
        <Modal
            open={!!selectedSurvey}
            onCancel={onOk}
            onOk={onClose}
            title={selectedSurvey?.name || 'Survey Details'}
        >
            {selectedSurvey?.questions?.length ? (
                <div>
                    {selectedSurvey.questions.map((question) => (
                        <div key={question.id} style={{ marginBottom: '1em' }}>
                            <strong>{question.question}</strong>
                            <ul>
                                {question.answers && question.answers.length > 0 ? (
                                    question.answers.map((answer) => (
                                        <li key={answer.id}>
                                            {answer.answer}
                                            <div style={{ marginTop: '0.5rem' }}>
                                                <VoiceRecorder audioFile={answer.answer} />
                                            </div>
                                        </li>
                                    ))
                                ) : (
                                    <li>
                                        No answers saved.
                                        <div style={{ marginTop: '0.5rem' }}>
                                            <VoiceRecorder />
                                        </div>
                                    </li>
                                )}
                            </ul>

                        </div>
                    ))}
                </div>
            ) : (
                <p>No questions found.</p>
            )}
        </Modal>
    );
};

export default SurveyDetailsModal;
