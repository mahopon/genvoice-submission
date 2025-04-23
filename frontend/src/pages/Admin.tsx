import React, { useEffect } from 'react'
import { useAuth } from '../context/AuthContext'
import { useNavigate } from 'react-router';
import UserTable from '../components/UserTable';

const Admin: React.FC = () => {
  const navigate = useNavigate();
  const { isAuthenticated, role } = useAuth();

  useEffect(() => {
    if (isAuthenticated === null) return;
    if (!isAuthenticated) {
      console.log('User is not authenticated, redirecting...');
      navigate('/login');
    }
    if (role != "ADMIN") {
      console.log('User is not admin role, redirecting...');
      navigate('/');
    }
  }, [isAuthenticated, navigate]);

  return (
    <UserTable />
  )
}

export default Admin
