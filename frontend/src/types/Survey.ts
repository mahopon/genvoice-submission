export type CreateSurveyRequest = {
    name: string;
};

export type CreateSurveyResponse = {
    survey_id: string;
};

export type CreateQuestionRequest = {
    survey_id: string;
    question: string;
};

export type CreateAnswerRequest = {
    survey_id: string;
    question_id: number;
    user_id: string;
    answer: string;
};

export type AnswerResponse = {
    id: number;
    user_id: string;
    survey_id: string;
    question_id: number;
    answer: string;
}

export type QuestionResponse = {
    id: string;
    question: string;
    created_date: Date;
    survey_id: string;
    answers?: AnswerResponse[];
}

export type SurveyResponse = {
    id: string;
    name: string;
    created_date: string;
    created_by: string;
    created_by_name: string;
    questions: QuestionResponse[];
}

export type CollatedSurveyResponse = {
    user_made: SurveyResponse[];
    others_made: SurveyResponse[];
}