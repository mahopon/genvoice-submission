import React, { useEffect } from 'react'
import LoginForm from '../components/LoginForm'
import "../css/Login.css"
import { useNavigate } from 'react-router';
import { useAuth } from "../context/AuthContext";

const Login: React.FC = () => {
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (isAuthenticated) {
      console.log('User is authenticated, redirecting...');
      navigate("/");
    }
  }, [isAuthenticated, navigate]);

  return (
    <>
      <LoginForm />
    </>
  )
}

export default Login