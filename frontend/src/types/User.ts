export type UserLoginRequest = {
    username: string;
    password: string;
}

export type RegisterUserRequest = {
    name: string;
    username: string;
    password: string;
}

export type PasswordChangeRequest = {
    current_password: string;
    new_password: string;
}

export type UserLoginResponse = {
    name: string;
    id: string;
    role: string;
}

export type UserDataRequest = {
    key: string;
    id: string;
    name: string;
    username: string;
    password?: string;
    role: string;
    createdDate: string;
}

export type UserUpdateRequest = {
    name: string;
    username: string;
    password?: string;
    role: string;
}

export type UserCreateRequest = {
    name: string;
    username: string;
    password?: string;
    role: string;
}

export type AuthResponse = {
    role: string;
    user_id: string;
}