import api from "./Interceptors"
import { AuthResponse, PasswordChangeRequest, RegisterUserRequest, UserCreateRequest, UserDataRequest, UserLoginRequest, UserLoginResponse, UserUpdateRequest } from "../types/User.ts"
import { handleApiError } from "../utils/APIerror.ts";

const loginUser = async (details: UserLoginRequest): Promise<UserLoginResponse> => {
    try {
        const { data } = await api.post<UserLoginResponse>("/user/login", details, { withCredentials: true });
        return data
    } catch (err) {
        return handleApiError(err)
    }
}

const logoutUser = async (): Promise<boolean> => {
    try {
        const res = await api.post<void>("/user/logout", {}, { withCredentials: true });
        return res.status === 200;
    } catch (err) {
        return handleApiError(err)
    }
}

const registerUser = async (req: RegisterUserRequest): Promise<boolean> => {
    try {
        const res = await api.post<void>("/user/register", req);
        return res.status === 200;
    } catch (err) {
        return handleApiError(err)
    }
}

const checkAuth = async (): Promise<AuthResponse> => {
    try {
        const res = await api.get<AuthResponse>("/user/auth", { withCredentials: true });
        return {
            role: res.data.role,
            user_id: res.data.user_id
        };
    } catch (err) {
        return handleApiError(err)
    }
};

const updatePassword = async (user_id: string, req: PasswordChangeRequest): Promise<boolean> => {
    try {
        const res = await api.patch("/user/edit/" + user_id, req, { withCredentials: true });
        return res.status === 200;
    } catch (err) {
        return handleApiError(err)
    }
};

const getAllUser = async (): Promise<UserDataRequest[]> => {
    try {
        const res = await api.get<UserDataRequest[]>("/user/all", { withCredentials: true });
        return res.data;
    } catch (err) {
        return handleApiError(err)
    }
}

const deleteUser = async (userId: string): Promise<boolean> => {
    try {
        const res = await api.delete<void>("/user/delete/" + userId, { withCredentials: true });
        return res.status === 200;
    } catch (err) {
        return handleApiError(err)
    }
}

const editWholeUser = async (userId: string, details: UserUpdateRequest): Promise<boolean> => {
    try {
        const res = await api.put<boolean>("/user/edit/whole/" + userId, details, { withCredentials: true });
        return res.status === 200;
    } catch (err) {
        return handleApiError(err)
    }
}

const createUser = async (user: UserCreateRequest): Promise<boolean> => {
    try {
        const res = await api.post<boolean>("/user/create", user, { withCredentials: true });
        return res.data;
    } catch (err) {
        return handleApiError(err)
    }
}

export { loginUser, logoutUser, registerUser, checkAuth, updatePassword, getAllUser, deleteUser, editWholeUser, createUser };
