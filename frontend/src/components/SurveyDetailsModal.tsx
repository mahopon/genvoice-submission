import React, { useEffect, useState } from 'react';
import { Modal } from 'antd';
import { AnswerResponse, CreateAnswerRequest, QuestionResponse, SurveyResponse } from '../types/Survey';
import VoiceRecorder from './VoiceRecorder';
import { MessageInstance } from 'antd/es/message/interface';
import { submitAnswer } from '../api/SurveyAPI';
import { base64ToBlob, blobToBase64 } from '../utils/encoding';

interface SurveyDetailsModalProps {
    selectedSurvey: SurveyResponse | null;
    setSelectedSurvey: React.Dispatch<React.SetStateAction<SurveyResponse | undefined>>;
    onSurveyComplete: () => void;
    message: MessageInstance;
}

const SurveyDetailsModal: React.FC<SurveyDetailsModalProps> = ({ selectedSurvey, setSelectedSurvey, message, onSurveyComplete }) => {
    const [recordings, setRecordings] = useState<Record<string, Blob | undefined>>({});
    const [recordingQuestionId, setRecordingQuestionId] = useState<string | null>(null);  // Track the recording question

    useEffect(() => {
        if (!selectedSurvey?.questions?.length) return;

        const updatedRecordings: Record<string, Blob | undefined> = {};

        selectedSurvey.questions.forEach((question: QuestionResponse) => {
            if (question.answers?.length) {
                question.answers.forEach((ans: AnswerResponse) => {
                    updatedRecordings[question.id] = base64ToBlob(ans.answer, "audio/webm");
                });
            }
        });

        setRecordings(updatedRecordings);
    }, [selectedSurvey]);

    const handleModalClose = () => {
        setSelectedSurvey(undefined);
    };

    const handleModalOk = async () => {
        let mapRecordings: CreateAnswerRequest[] = [];

        if (Object.keys(recordings).length > 0) {
            mapRecordings = await Promise.all(
                Object.entries(recordings).map(async ([questionId, blob]) => {
                    const question = selectedSurvey?.questions?.find(q => q.id.toString() === questionId);
                    const base64Answer = blob ? await blobToBase64(blob) : "";

                    return {
                        survey_id: selectedSurvey!.id,
                        question_id: Number(question!.id),
                        user_id: "",
                        answer: base64Answer,
                    };
                })
            );
        }

        if (mapRecordings && mapRecordings.length > 0) {
            message.success("Submitted answers");
            submitAnswer(mapRecordings).then(() => {
                onSurveyComplete();
                setSelectedSurvey(undefined);
            });
        } else {
            message.warning("No answers to submit");
            setSelectedSurvey(undefined);
        }
    };

    const handleDeleteClick = (questionId: string) => {
        setRecordings((prev) => {
            const updatedRecordings = { ...prev };
            updatedRecordings[questionId] = undefined;
            return updatedRecordings;
        });
    };

    return (
        <>
            <Modal
                open={!!selectedSurvey}
                onCancel={handleModalClose}
                onOk={handleModalOk}
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
                                                <div style={{ marginTop: '0.5rem' }}>
                                                    <VoiceRecorder
                                                        questionId={question.id}
                                                        existingAudio={recordings[question.id]}
                                                        isRecording={recordingQuestionId === question.id}
                                                        isDisabled={recordingQuestionId !== null && recordingQuestionId !== question.id}
                                                        onDelete={handleDeleteClick}
                                                        onRecordingComplete={(blob: Blob) => {
                                                            setRecordings((prev) => ({ ...prev, [question.id]: blob }));
                                                            setRecordingQuestionId(null);
                                                        }}
                                                        onStartRecording={() => {
                                                            setRecordingQuestionId(question.id);
                                                        }}
                                                    />
                                                </div>
                                            </li>
                                        ))
                                    ) : (
                                        <li>
                                            No answers saved.
                                            <div style={{ marginTop: '0.5rem' }}>
                                                <VoiceRecorder
                                                    questionId={question.id}
                                                    isRecording={recordingQuestionId === question.id}
                                                    isDisabled={recordingQuestionId !== null && recordingQuestionId !== question.id}
                                                    onDelete={handleDeleteClick}
                                                    onRecordingComplete={(blob: Blob) => {
                                                        setRecordings((prev) => ({ ...prev, [question.id]: blob }));
                                                        setRecordingQuestionId(null);
                                                    }}
                                                    onStartRecording={() => {
                                                        setRecordingQuestionId(question.id);
                                                    }}
                                                />
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
        </>
    );
};

export default SurveyDetailsModal;
