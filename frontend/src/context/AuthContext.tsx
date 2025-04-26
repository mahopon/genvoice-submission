import React, { createContext, useState, useContext, ReactNode, useEffect } from 'react';
import { checkAuth } from '../api/UserAPI';

interface AuthContextType {
    isAuthenticated: boolean | null;
    setAuthStatus: React.Dispatch<React.SetStateAction<boolean | null>>;
    role: string | null;
    setRole: React.Dispatch<React.SetStateAction<string | null>>;
    userId: string | null;
    setUserId: React.Dispatch<React.SetStateAction<string | null>>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface AuthProviderProps {
    children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
    const [isAuthenticated, setAuthStatus] = useState<boolean | null>(null);
    const [role, setRole] = useState<string | null>(null);
    const [userId, setUserId] = useState<string | null>(null);

    useEffect(() => {
        const authenticateUser = async () => {
            try {
                const authData = await checkAuth();
                setAuthStatus(true);
                setRole(authData.role.toUpperCase());
                setUserId(authData.user_id);
                // eslint-disable-next-line @typescript-eslint/no-unused-vars
            } catch (err) {
                setAuthStatus(false);
                setRole(null);
                setUserId(null);
            }

        };
        authenticateUser();
    }, []);

    return (
        <AuthContext.Provider value={{ isAuthenticated, setAuthStatus, role, setRole, userId, setUserId }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = (): AuthContextType => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};
