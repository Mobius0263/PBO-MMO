import { useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import Sidebar from './Sidebar';
import '../styles/MyTeam.css';

function MyTeam() {
    const { user } = useContext(AuthContext);

    return (
        <div className="page-container">
            <Sidebar />
            <div className="main-content">
                <div className="page-header">
                    <h1>My Team</h1>
                    <p>Manage your team members and collaboration</p>
                </div>
                
                <div className="team-content">
                    <div className="team-stats">
                        <div className="stat-card">
                            <div className="stat-icon">
                                <i className="fas fa-users"></i>
                            </div>
                            <div className="stat-info">
                                <h3>5</h3>
                                <p>Team Members</p>
                            </div>
                        </div>
                        
                        <div className="stat-card">
                            <div className="stat-icon">
                                <i className="fas fa-project-diagram"></i>
                            </div>
                            <div className="stat-info">
                                <h3>3</h3>
                                <p>Active Projects</p>
                            </div>
                        </div>
                        
                        <div className="stat-card">
                            <div className="stat-icon">
                                <i className="fas fa-calendar-check"></i>
                            </div>
                            <div className="stat-info">
                                <h3>12</h3>
                                <p>Completed Tasks</p>
                            </div>
                        </div>
                    </div>

                    <div className="team-members">
                        <div className="section-header">
                            <h2>Team Members</h2>
                            <button className="btn-primary">
                                <i className="fas fa-plus"></i>
                                Add Member
                            </button>
                        </div>
                        
                        <div className="members-grid">
                            <div className="member-card">
                                <div className="member-avatar">
                                    <div className="avatar-placeholder">JD</div>
                                </div>
                                <div className="member-info">
                                    <h3>John Doe</h3>
                                    <p>Project Manager</p>
                                    <span className="status online">Online</span>
                                </div>
                                <div className="member-actions">
                                    <button className="btn-icon">
                                        <i className="fas fa-envelope"></i>
                                    </button>
                                    <button className="btn-icon">
                                        <i className="fas fa-phone"></i>
                                    </button>
                                </div>
                            </div>
                            
                            <div className="member-card">
                                <div className="member-avatar">
                                    <div className="avatar-placeholder">AS</div>
                                </div>
                                <div className="member-info">
                                    <h3>Alice Smith</h3>
                                    <p>Developer</p>
                                    <span className="status away">Away</span>
                                </div>
                                <div className="member-actions">
                                    <button className="btn-icon">
                                        <i className="fas fa-envelope"></i>
                                    </button>
                                    <button className="btn-icon">
                                        <i className="fas fa-phone"></i>
                                    </button>
                                </div>
                            </div>
                            
                            <div className="member-card">
                                <div className="member-avatar">
                                    <div className="avatar-placeholder">BJ</div>
                                </div>
                                <div className="member-info">
                                    <h3>Bob Johnson</h3>
                                    <p>Designer</p>
                                    <span className="status offline">Offline</span>
                                </div>
                                <div className="member-actions">
                                    <button className="btn-icon">
                                        <i className="fas fa-envelope"></i>
                                    </button>
                                    <button className="btn-icon">
                                        <i className="fas fa-phone"></i>
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default MyTeam;
