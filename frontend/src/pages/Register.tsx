import React, { useEffect, useState } from 'react';
import RegisterForm from '../components/RegisterForm';
import { useNavigate } from 'react-router';
import { useAuth } from '../context/AuthContext';

const Register: React.FC = () => {

  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();
  const [registered, setRegistered] = useState<boolean>(false);

  useEffect(() => {
    if (isAuthenticated === null) return;
    if (isAuthenticated) {
      console.log("User is authenticated, redirecting...");
      navigate("/");
    }
  }, [isAuthenticated, navigate]);

  useEffect(() => {
    if (registered) {
      navigate("/")
    }
  }, [registered, navigate])

  return (
    <>
      <RegisterForm setRegistered={setRegistered} />
    </>
  );
}

export default Register;
