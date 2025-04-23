import api from "./Interceptors"
import { RegisterUserRequest, UserLoginRequest } from "../types/User.ts"

const loginUser = async (details: UserLoginRequest): Promise<boolean> => {
    try {
        const res = await api.post("/user/login", details, { withCredentials: true });

        if (res.status === 200) {
            console.log('Login successful');
            return true;
        } else {
            console.error('Login failed: No access token received');
            return false;
        }
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
    const login: UserLoginRequest = {
        username: req.username,
        password: req.password
    }
    loginUser(login);
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

export { loginUser, logoutUser, registerUser, checkAuth };