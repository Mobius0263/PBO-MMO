import { useState, useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import Sidebar from './Sidebar';
import '../styles/Settings.css';

function Settings() {
    const { user, updateUser } = useContext(AuthContext);
    const [activeTab, setActiveTab] = useState('profile');
    const [formData, setFormData] = useState({
        nama: user?.nama || '',
        email: user?.email || '',
        bio: user?.bio || '',
        currentPassword: '',
        newPassword: '',
        confirmPassword: ''
    });
    const [notifications, setNotifications] = useState({
        emailNotifications: true,
        pushNotifications: true,
        meetingReminders: true,
        weeklyDigest: false
    });

    const handleProfileUpdate = (e) => {
        e.preventDefault();
        // TODO: Call API to update profile
        console.log('Updating profile:', formData);
    };

    const handlePasswordChange = (e) => {
        e.preventDefault();
        if (formData.newPassword !== formData.confirmPassword) {
            alert('New passwords do not match');
            return;
        }
        // TODO: Call API to change password
        console.log('Changing password');
    };

    const handleNotificationChange = (setting) => {
        setNotifications(prev => ({
            ...prev,
            [setting]: !prev[setting]
        }));
    };

    return (
        <div className="page-container">
            <Sidebar />
            <div className="main-content">
                <div className="page-header">
                    <h1>Settings</h1>
                    <p>Manage your account and preferences</p>
                </div>
                
                <div className="settings-content">
                    <div className="settings-tabs">
                        <button 
                            className={`tab-button ${activeTab === 'profile' ? 'active' : ''}`}
                            onClick={() => setActiveTab('profile')}
                        >
                            <i className="fas fa-user"></i>
                            Profile
                        </button>
                        <button 
                            className={`tab-button ${activeTab === 'security' ? 'active' : ''}`}
                            onClick={() => setActiveTab('security')}
                        >
                            <i className="fas fa-shield-alt"></i>
                            Security
                        </button>
                        <button 
                            className={`tab-button ${activeTab === 'notifications' ? 'active' : ''}`}
                            onClick={() => setActiveTab('notifications')}
                        >
                            <i className="fas fa-bell"></i>
                            Notifications
                        </button>
                        <button 
                            className={`tab-button ${activeTab === 'preferences' ? 'active' : ''}`}
                            onClick={() => setActiveTab('preferences')}
                        >
                            <i className="fas fa-cog"></i>
                            Preferences
                        </button>
                    </div>

                    <div className="settings-panel">
                        {activeTab === 'profile' && (
                            <div className="tab-content">
                                <h2>Profile Information</h2>
                                <form onSubmit={handleProfileUpdate} className="settings-form">
                                    <div className="profile-picture-section">
                                        <div className="current-picture">
                                            {user?.profileImage ? (
                                                <img src={user.profileImage} alt="Profile" />
                                            ) : (
                                                <div className="avatar-placeholder">
                                                    {user?.nama ? user.nama.charAt(0).toUpperCase() : 'U'}
                                                </div>
                                            )}
                                        </div>
                                        <div className="picture-actions">
                                            <button type="button" className="btn-secondary">
                                                <i className="fas fa-camera"></i>
                                                Change Picture
                                            </button>
                                            <button type="button" className="btn-danger">
                                                <i className="fas fa-trash"></i>
                                                Remove
                                            </button>
                                        </div>
                                    </div>

                                    <div className="form-group">
                                        <label>Full Name</label>
                                        <input
                                            type="text"
                                            value={formData.nama}
                                            onChange={(e) => setFormData({...formData, nama: e.target.value})}
                                        />
                                    </div>
                                    
                                    <div className="form-group">
                                        <label>Email</label>
                                        <input
                                            type="email"
                                            value={formData.email}
                                            onChange={(e) => setFormData({...formData, email: e.target.value})}
                                        />
                                    </div>
                                    
                                    <div className="form-group">
                                        <label>Bio</label>
                                        <textarea
                                            value={formData.bio}
                                            onChange={(e) => setFormData({...formData, bio: e.target.value})}
                                            rows="4"
                                            placeholder="Tell us about yourself..."
                                        />
                                    </div>
                                    
                                    <button type="submit" className="btn-primary">
                                        Save Changes
                                    </button>
                                </form>
                            </div>
                        )}

                        {activeTab === 'security' && (
                            <div className="tab-content">
                                <h2>Change Password</h2>
                                <form onSubmit={handlePasswordChange} className="settings-form">
                                    <div className="form-group">
                                        <label>Current Password</label>
                                        <input
                                            type="password"
                                            value={formData.currentPassword}
                                            onChange={(e) => setFormData({...formData, currentPassword: e.target.value})}
                                        />
                                    </div>
                                    
                                    <div className="form-group">
                                        <label>New Password</label>
                                        <input
                                            type="password"
                                            value={formData.newPassword}
                                            onChange={(e) => setFormData({...formData, newPassword: e.target.value})}
                                        />
                                    </div>
                                    
                                    <div className="form-group">
                                        <label>Confirm New Password</label>
                                        <input
                                            type="password"
                                            value={formData.confirmPassword}
                                            onChange={(e) => setFormData({...formData, confirmPassword: e.target.value})}
                                        />
                                    </div>
                                    
                                    <button type="submit" className="btn-primary">
                                        Update Password
                                    </button>
                                </form>

                                <div className="security-section">
                                    <h3>Two-Factor Authentication</h3>
                                    <p>Add an extra layer of security to your account</p>
                                    <button className="btn-secondary">
                                        <i className="fas fa-shield-alt"></i>
                                        Enable 2FA
                                    </button>
                                </div>
                            </div>
                        )}

                        {activeTab === 'notifications' && (
                            <div className="tab-content">
                                <h2>Notification Preferences</h2>
                                <div className="notification-settings">
                                    <div className="notification-item">
                                        <div className="notification-info">
                                            <h3>Email Notifications</h3>
                                            <p>Receive important updates via email</p>
                                        </div>
                                        <label className="toggle-switch">
                                            <input
                                                type="checkbox"
                                                checked={notifications.emailNotifications}
                                                onChange={() => handleNotificationChange('emailNotifications')}
                                            />
                                            <span className="toggle-slider"></span>
                                        </label>
                                    </div>

                                    <div className="notification-item">
                                        <div className="notification-info">
                                            <h3>Push Notifications</h3>
                                            <p>Get instant notifications on your device</p>
                                        </div>
                                        <label className="toggle-switch">
                                            <input
                                                type="checkbox"
                                                checked={notifications.pushNotifications}
                                                onChange={() => handleNotificationChange('pushNotifications')}
                                            />
                                            <span className="toggle-slider"></span>
                                        </label>
                                    </div>

                                    <div className="notification-item">
                                        <div className="notification-info">
                                            <h3>Meeting Reminders</h3>
                                            <p>Get reminded before your meetings start</p>
                                        </div>
                                        <label className="toggle-switch">
                                            <input
                                                type="checkbox"
                                                checked={notifications.meetingReminders}
                                                onChange={() => handleNotificationChange('meetingReminders')}
                                            />
                                            <span className="toggle-slider"></span>
                                        </label>
                                    </div>

                                    <div className="notification-item">
                                        <div className="notification-info">
                                            <h3>Weekly Digest</h3>
                                            <p>Receive a weekly summary of your activities</p>
                                        </div>
                                        <label className="toggle-switch">
                                            <input
                                                type="checkbox"
                                                checked={notifications.weeklyDigest}
                                                onChange={() => handleNotificationChange('weeklyDigest')}
                                            />
                                            <span className="toggle-slider"></span>
                                        </label>
                                    </div>
                                </div>
                            </div>
                        )}

                        {activeTab === 'preferences' && (
                            <div className="tab-content">
                                <h2>Application Preferences</h2>
                                <div className="preferences-section">
                                    <div className="preference-item">
                                        <label>Theme</label>
                                        <select className="preference-select">
                                            <option value="light">Light</option>
                                            <option value="dark">Dark</option>
                                            <option value="auto">Auto</option>
                                        </select>
                                    </div>

                                    <div className="preference-item">
                                        <label>Language</label>
                                        <select className="preference-select">
                                            <option value="en">English</option>
                                            <option value="id">Bahasa Indonesia</option>
                                        </select>
                                    </div>

                                    <div className="preference-item">
                                        <label>Timezone</label>
                                        <select className="preference-select">
                                            <option value="UTC+7">UTC+7 (Jakarta)</option>
                                            <option value="UTC+0">UTC+0 (London)</option>
                                            <option value="UTC-5">UTC-5 (New York)</option>
                                        </select>
                                    </div>
                                </div>

                                <button className="btn-primary">
                                    Save Preferences
                                </button>
                            </div>
                        )}
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Settings;
