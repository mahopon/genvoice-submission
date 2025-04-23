import React, { useEffect } from 'react'
import { useAuth } from '../context/AuthContext'
import { useNavigate } from 'react-router';

const Admin: React.FC = () => {
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
    <div>Admin</div>
  )
}

export default Admin
