import { useContext, useState } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';
import '../styles/Sidebar.css';

function Sidebar() {
    const { user, logoutUser } = useContext(AuthContext);
    const [showLogout, setShowLogout] = useState(false);
    const location = useLocation();
    const navigate = useNavigate();

    const isActive = (path) => {
        return location.pathname === path;
    };

    const handleLogout = () => {
        logoutUser();
        navigate('/login');
    };

    const getInitials = () => {
        if (!user || !user.nama) return '?';
        
        const names = user.nama.split(' ');
        if (names.length === 1) return names[0].charAt(0).toUpperCase();
        return (names[0].charAt(0) + names[names.length - 1].charAt(0)).toUpperCase();
    };

    return (
        <div className="sidebar">
            <div className="sidebar-content">
                <div className="sidebar-header">
                    <div className="app-logo">
                        <i className="fas fa-video-camera"></i>
                        <span>MMO</span>
                    </div>
                </div>

                <ul className="sidebar-menu">
                    <li className={isActive('/dashboard') ? 'active' : ''}>
                        <Link to="/dashboard">
                            <i className="fas fa-tachometer-alt"></i>
                            <span>Dashboard</span>
                        </Link>
                    </li>
                    <li className={isActive('/my-team') ? 'active' : ''}>
                        <Link to="/my-team">
                            <i className="fas fa-users"></i>
                            <span>My Team</span>
                        </Link>
                    </li>
                    <li className={isActive('/meetings') ? 'active' : ''}>
                        <Link to="/meetings">
                            <i className="fas fa-video"></i>
                            <span>Meetings</span>
                        </Link>
                    </li>
                    <li className={isActive('/settings') ? 'active' : ''}>
                        <Link to="/settings">
                            <i className="fas fa-cog"></i>
                            <span>Settings</span>
                        </Link>
                    </li>
                </ul>

                <div className="sidebar-footer">
                    <div className="user-info" onClick={() => setShowLogout(!showLogout)}>
                        <div className="user-avatar">
                            {user?.profileImage ? (
                                <img 
                                    src={user.profileImage} 
                                    alt={user.nama} 
                                    className="profile-image" 
                                    onError={(e) => {
                                        e.target.onerror = null;
                                        e.target.style.display = 'none';
                                        e.target.nextSibling.style.display = 'flex';
                                    }}
                                />
                            ) : (
                                <div className="avatar-placeholder">{getInitials()}</div>
                            )}
                        </div>
                        <div className="user-details">
                            <div className="user-name">{user?.nama || 'User'}</div>
                            <div className="user-role">{user?.role || 'Member'}</div>
                        </div>
                    </div>

                    {showLogout && (
                        <div className="logout-popup">
                            <button onClick={handleLogout} className="logout-button">
                                <i className="fas fa-sign-out-alt"></i>
                                <span>Logout</span>
                            </button>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}

export default Sidebar;