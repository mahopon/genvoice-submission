import { Outlet } from 'react-router';
import Nav from '../components/Nav';

const Layout: React.FC = () => {
    return (
        <div>
            <Nav />
            <main className="p-4">
                <Outlet />
            </main>
        </div>
    );
};

export default Layout;
