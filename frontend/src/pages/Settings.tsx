import React, { useEffect } from 'react'
import { useNavigate } from 'react-router';
import { useAuth } from '../context/AuthContext';
import SettingsForm from '../components/SettingsForm';

const Settings: React.FC = () => {
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (isAuthenticated === null) return;
    if (!isAuthenticated) {
      console.log('User is not authenticated, redirecting...');
      navigate('/login');
    }
  }, [isAuthenticated, navigate]);
  return (
    <SettingsForm />
  )
}

export default Settings