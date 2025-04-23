export type Survey = {
    id: string;
    name: string;
    created_date: Date;
    question: Question[];
}

type Question = {
    question: string;
    created_date: Date;
}

export type Answer = {
    id: number;
    userid: string;
    surveyid: string;
    questionid: string;
    Answer: Blob;
}

export type CreateSurveyRequest = {
    user_id: string;
    name: string;
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
    questions: QuestionResponse[];
}