import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from "react-router"
import './index.css'
import Home from "./pages/Home.tsx"
import Register from './pages/Register.tsx'
import Login from './pages/Login.tsx'
import Settings from './pages/Settings.tsx'
import Admin from './pages/Admin.tsx'
import '@ant-design/v5-patch-for-react-19'
import Layout from './pages/Layout.tsx'
import { AuthProvider } from './context/AuthContext.tsx'

createRoot(document.getElementById('root')!).render(
  <BrowserRouter basename={"/genvoice-submission/"}>
    <AuthProvider>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="register" element={<Register />} />
          <Route path="login" element={<Login />} />
          <Route path="settings" element={<Settings />} />
          <Route path="admin" element={<Admin />} />
        </Route>
      </Routes>
    </AuthProvider>
  </BrowserRouter>
  ,
)
