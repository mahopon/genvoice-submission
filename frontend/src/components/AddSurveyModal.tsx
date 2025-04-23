import React, { useState } from 'react';
import { Button, Modal, Form, Input, Space, Row, Col, Tooltip, message } from 'antd';
import { CloseOutlined } from '@ant-design/icons';
import { CreateQuestionRequest, CreateSurveyRequest, CreateSurveyResponse } from '../types/Survey';
import { createQuestion, createSurvey } from '../api/SurveyAPI';

interface Props {
    open: boolean;
    onClose: () => void;
}

interface FormValues {
    surveyName: string;
    questions: { question: string }[];
}

const CreateSurveyModal: React.FC<Props> = ({ open, onClose }) => {
    const [form] = Form.useForm();
    const [questions, setQuestions] = useState([{}]);
    const [messageApi, contextHolder] = message.useMessage();

    const addQuestion = () => {
        setQuestions([...questions, {}]);
    };

    const handleClose = () => {
        form.resetFields();
        setQuestions([{}]);
        onClose();
    };

    const handleOk = async (values: FormValues) => {
        const { surveyName, questions } = values;

        const newSurvey: CreateSurveyRequest = {
            name: surveyName
        };

        if (questions.length === 0) {
            alert("At least one question is required.");
            return;
        }

        try {
            const survey: CreateSurveyResponse = await createSurvey(newSurvey);
            console.log('Survey created with ID:', survey.survey_id);
            const questionArr: CreateQuestionRequest[] = questions.map((qn) => ({
                survey_id: survey.survey_id,
                question: qn.question
            }));
            await createQuestion(questionArr);
            messageApi.success("Created survey");
            handleClose();
        } catch (error) {
            console.error("Error creating survey:", error);
        }
    };

    return (
        <Modal
            title="Create Survey"
            open={open}
            onCancel={handleClose}
            onOk={() => form.submit()}
            destroyOnClose
            style={{ maxWidth: '750px' }}
            styles={{ body: { maxHeight: '60vh', overflowY: 'auto', overflowX: 'hidden' } }}
        >
            {contextHolder}
            <Form
                form={form}
                layout="vertical"
                name="createSurveyForm"
                onFinish={handleOk}
            >
                <Form.Item
                    label="Survey Name"
                    name="surveyName"
                    rules={[{ required: true, message: 'Please enter a survey name' }]}
                    style={{ marginBottom: 16 }}
                >
                    <Input style={{ width: '100%' }} />
                </Form.Item>

                <Form.List
                    name="questions"
                    initialValue={questions}
                    rules={[
                        {
                            validator: async (_, questions) => {
                                if (!questions || questions.length < 1) {
                                    return Promise.reject(new Error('At least one question is required'));
                                }
                            },
                        },
                    ]}
                >
                    {(fields, { add, remove }) => (
                        <>
                            {fields.map(({ key, name }) => (
                                <Row key={key} gutter={16} align="middle" style={{ display: 'flex', alignItems: 'center', marginBottom: 16 }}>
                                    <Col span={20}>
                                        <div style={{ display: 'flex', alignItems: 'center' }}>
                                            <span style={{ marginRight: 8 }}>Question {name + 1}</span>
                                            {fields.length > 1 && (
                                                <Tooltip title="Remove Question">
                                                    <Button
                                                        type="link"
                                                        icon={<CloseOutlined />}
                                                        onClick={() => {
                                                            remove(name);
                                                            setQuestions(questions.filter((_, index) => index !== name));
                                                        }}
                                                        style={{ padding: 0 }}
                                                    />
                                                </Tooltip>
                                            )}
                                        </div>
                                        <Form.Item
                                            name={[name, 'question']}
                                            rules={[{ required: true, message: 'Please enter a question' }]}
                                            style={{ marginBottom: 8 }}
                                        >
                                            <Input />
                                        </Form.Item>
                                    </Col>
                                </Row>
                            ))}

                            <Space style={{ width: '100%', marginTop: 10 }}>
                                <Button
                                    type="dashed"
                                    onClick={() => {
                                        add();
                                        addQuestion();
                                    }}
                                    style={{ width: '100%' }}
                                >
                                    Add Question
                                </Button>
                            </Space>
                        </>
                    )}
                </Form.List>
            </Form>
        </Modal>
    );
};

export default CreateSurveyModal;
