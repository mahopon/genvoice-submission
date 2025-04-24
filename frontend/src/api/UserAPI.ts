import api from "./Interceptors"
import { PasswordChangeRequest, RegisterUserRequest, UserCreateRequest, UserDataRequest, UserLoginRequest, UserLoginResponse, UserUpdateRequest } from "../types/User.ts"


const loginUser = async (details: UserLoginRequest): Promise<UserLoginResponse | false> => {
    try {
        const res = await api.post("/user/login", details, { withCredentials: true });
        console.log('Login successful');
        return {
            name: res.data.name,
            id: res.data.id,
            role: res.data.role
        };
    } catch (error) {
        console.error('Login failed:', error);
        return false;
    }
}

const logoutUser = async (): Promise<boolean> => {
    try {
        const res = await api.post("/user/logout", {}, { withCredentials: true });
        if (res.status === 200) {
            console.log("Logout successful");
            return true;
        }
    } catch (error) {
        console.log(error);
    }
    return false;
}

const registerUser = async (req: RegisterUserRequest) => {
    const res = await api.post("/user/register", req)
    return res.data;
}

const checkAuth = () => {
    return api.get("/user/auth", { withCredentials: true })
        .then((res) => {
            return {
                isAuthenticated: true,
                role: res.data.role,
                userId: res.data.user_id
            };
        })
        .catch((err) => {
            console.error("Error during authentication check:", err);
            return { isAuthenticated: false, role: null, userId: null };
        });
};

const updatePassword = async (user_id: string, req: PasswordChangeRequest): Promise<boolean> => {
    try {
        const res = await api.patch("/user/edit/" + user_id, req, { withCredentials: true });
        return res.status === 200;
    } catch (err) {
        console.error("Error during password change:", err);
        return false;
    }
};

const getAllUser = async (): Promise<UserDataRequest[]> => {
    const res = await api.get<UserDataRequest[]>("/user/all", { withCredentials: true });
    return res.data;
}

const deleteUser = async (userId: string): Promise<boolean> => {
    const res = await api.delete<boolean>("/user/delete/" + userId, { withCredentials: true })
    return res.status == 200;
}

const editWholeUser = async (userId: string, details: UserUpdateRequest): Promise<boolean> => {
    const res = await api.put<boolean>("/user/edit/whole/" + userId, details, { withCredentials: true })
    return res.status == 200;
}

const createUser = async (user: UserCreateRequest): Promise<boolean> => {
    try {
        const res = await api.post<boolean>("/user/create", user, { withCredentials: true });
        return res.status === 200;
    } catch (err) {
        console.log(err);
        return false;
    }
}

export { loginUser, logoutUser, registerUser, checkAuth, updatePassword, getAllUser, deleteUser, editWholeUser, createUser };