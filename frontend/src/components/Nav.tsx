import { Button } from 'antd';
import { Header } from 'antd/es/layout/layout';
import React, { useEffect, useState } from 'react';
import { useNavigate, useLocation } from 'react-router';
import { logoutUser } from '../api/UserAPI';
import { useAuth } from '../context/AuthContext';

const Nav: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();  // Get current route
  const { isAuthenticated, setAuthStatus, role } = useAuth();
  const [currentPageName, setCurrentPageName] = useState<string>('Home');

  useEffect(() => {
    const routeToName: Record<string, string> = {
      '/': 'Home',
      '/login': 'Login',
      '/register': 'Register',
      '/settings': 'Settings'
    };
    setCurrentPageName(routeToName[location.pathname]);
  }, [location.pathname]);

  const navRegister = () => {
    console.log('Navigate to Register');
    navigate("/register");
  };

  const navLogin = () => {
    console.log('Navigate to Login');
    navigate("/login");
  };

  const handleLogout = async () => {
    console.log('Logout user');
    if (await logoutUser()) {
      console.log("Logged out");
      setAuthStatus(false);
      navigate("/");
    }
  };
  const navSettings = () => {
    console.log('Navigate to Settings');
    navigate("/settings");
  };

  const navHome = () => {
    console.log('Navigate to Home');
    navigate("/");
  };

  // Utility function to disable button if already on that page
  const isActivePage = (path: string) => location.pathname === path;

  return (
    <Header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '0 20px', backgroundColor: "white" }}>
      <div style={{ fontSize: '20px' }}>
        {currentPageName}
      </div>
      <div style={{ display: 'flex', gap: '10px' }}>
        <Button type="primary" onClick={navHome} disabled={isActivePage('/')}>Home</Button>
        {!isAuthenticated && (
          <>
            <Button type="primary" onClick={navRegister} disabled={isActivePage('/register')}>Register</Button>
            <Button type="primary" onClick={navLogin} disabled={isActivePage('/login')}>Login</Button>
          </>
        )}
        {isAuthenticated && (
          <>
            {role == "ADMIN" && (
              <Button type="primary" onClick={navSettings} disabled={isActivePage('/admin')}>Settings</Button>
            )
            }
            <Button type="primary" onClick={navSettings} disabled={isActivePage('/settings')}>Settings</Button>
            <Button type="primary" danger onClick={handleLogout}>Logout</Button>
          </>
        )}
      </div>
    </Header>
  );
};

export default Nav;
