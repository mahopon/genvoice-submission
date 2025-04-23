import { CreateSurveyRequest, CreateQuestionRequest, CreateAnswerRequest, SurveyResponse } from "../types/Survey";
import api from "./Interceptors"

const createSurvey = async (req: CreateSurveyRequest) => {
    const res = await api.post("/survey", req);
    return res.data;
};

const createQuestion = async (req: CreateQuestionRequest) => {
    const res = await api.post("/survey/question", req);
    return res.data;
};

const submitAnswer = async (req: CreateAnswerRequest[]) => {
    console.log(req);
    const res = await api.post("/survey/answer", req, { withCredentials: true });
    return res.data;
};

const getSurveys = async (): Promise<SurveyResponse[]> => {
    const res = await api.get("/survey", { withCredentials: true });
    return res.data;
}

export { createSurvey, createQuestion, submitAnswer, getSurveys };