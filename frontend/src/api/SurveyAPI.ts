import { CreateSurveyRequest, CreateQuestionRequest, CreateAnswerRequest, CollatedSurveyResponse, CreateSurveyResponse } from "../types/Survey";
import api from "./Interceptors";
import { handleApiError } from "../utils/APIerror";

const createSurvey = async (req: CreateSurveyRequest): Promise<CreateSurveyResponse> => {
    try {
        const { data } = await api.post<CreateSurveyResponse>("/survey", req, { withCredentials: true });
        return data;
    } catch (error) {
        return handleApiError(error);
    }
};

const createQuestion = async (questions: CreateQuestionRequest[]): Promise<void> => {
    try {
        await api.post("/survey/question", questions, { withCredentials: true });
    } catch (error) {
        handleApiError(error);
    }
};

const submitAnswer = async (answers: CreateAnswerRequest[]): Promise<void> => {
    try {
        await api.post("/survey/answer", answers, { withCredentials: true });
    } catch (error) {
        handleApiError(error);
    }
};

const getSurveys = async (): Promise<CollatedSurveyResponse> => {
    try {
        const { data } = await api.get<CollatedSurveyResponse>("/survey", { withCredentials: true });
        return data;
    } catch (error) {
        return handleApiError(error);
    }
};

const deleteSurvey = async (surveyId: string): Promise<void> => {
    try {
        await api.delete(`/survey/delete/${surveyId}`, { withCredentials: true });
    } catch (error) {
        handleApiError(error);
    }
};

export { deleteSurvey, getSurveys, submitAnswer, createQuestion, createSurvey };