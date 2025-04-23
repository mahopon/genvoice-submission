import { CreateSurveyRequest, CreateQuestionRequest, CreateAnswerRequest, CollatedSurveyResponse } from "../types/Survey";
import api from "./Interceptors"

const createSurvey = async (req: CreateSurveyRequest) => {
    const res = await api.post("/survey", req, { withCredentials: true });
    return res.data;
};

const createQuestion = async (req: CreateQuestionRequest[]) => {
    const res = await api.post("/survey/question", req, { withCredentials: true });
    return res.data;
};

const submitAnswer = async (req: CreateAnswerRequest[]) => {
    console.log(req);
    const res = await api.post("/survey/answer", req, { withCredentials: true });
    return res.data;
};

const getSurveys = async (): Promise<CollatedSurveyResponse> => {
    const res = await api.get("/survey", { withCredentials: true });
    return res.data;
}

const deleteSurvey = async (surveyId: string) => {
    const res = await api.delete("/survey/delete/" + surveyId, { withCredentials: true });
    return res.status == 200;
}

export { createSurvey, createQuestion, submitAnswer, getSurveys, deleteSurvey };