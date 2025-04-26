import axios from "axios";

type ApiError = {
    status: number;
    message: string;
}

const handleApiError = (error: unknown): never => {
    if (axios.isAxiosError(error)) {
        const status = error.response?.status ?? 500;
        const message = (error.response?.data.message || error.response?.data) ?? "An unknown error occurred.";
        throw { status, message };
    }
    throw { status: 500, message: "An unknown error occurred." };
};

export { handleApiError };
export type { ApiError };
