import { useState, useEffect, useContext } from 'react';
import { Link } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';
import Sidebar from './Sidebar';
import '../styles/Dashboard.css';

function Dashboard() {
    const { user } = useContext(AuthContext);
    const [upcomingMeetings, setUpcomingMeetings] = useState([]);
    const [teamMembers, setTeamMembers] = useState([]);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        // Fungsi untuk mengambil data meeting dan anggota tim
        const fetchDashboardData = async () => {
            setIsLoading(true);
            try {
                // Di implementasi sebenarnya, kita akan memanggil API
                // Untuk sekarang, kita gunakan data dummy

                // Simulasi delay network
                setTimeout(() => {
                    // Data dummy untuk upcoming meetings
                    setUpcomingMeetings([
                        {
                            id: 1,
                            title: "Project Kickoff",
                            time: "10:00 AM",
                            day: "Today",
                            participants: 5
                        },
                        {
                            id: 2,
                            title: "Weekly Standup",
                            time: "2:30 PM",
                            day: "Today",
                            participants: 8
                        },
                        {
                            id: 3,
                            title: "Design Review",
                            time: "11:00 AM",
                            day: "Tomorrow",
                            participants: 4
                        }
                    ]);
                    
                    // Data dummy untuk team members
                    setTeamMembers([
                        {
                            id: 1,
                            name: "John Doe",
                            role: "Developer",
                            status: "Online"
                        },
                        {
                            id: 2,
                            name: "Jane Smith",
                            role: "Designer",
                            status: "Away"
                        },
                        {
                            id: 3,
                            name: "Bob Johnson",
                            role: "Product Manager",
                            status: "Offline"
                        }
                    ]);
                    
                    setIsLoading(false);
                }, 800);
            } catch (error) {
                console.error("Error fetching dashboard data:", error);
                setIsLoading(false);
            }
        };

        fetchDashboardData();
    }, []);

    // Tampilkan loading spinner jika user belum tersedia atau data sedang loading
    if (!user || isLoading) {
        return (
            <div className="loading-container">
                <div className="loading-spinner"></div>
                <p>Loading...</p>
            </div>
        );
    }

    return (
        <div className="dashboard-main">
            <Sidebar />
            
            <div className="dashboard-content">
                <div className="dashboard-header">
                    <h1>Welcome back, {user.nama}!</h1>
                    <p>Here's what's happening with your team today.</p>
                </div>

                <div className="dashboard-grid">
                    {/* Upcoming Meetings Card */}
                    <div className="dashboard-card meetings">
                        <div className="card-header">
                            <h3>Upcoming Meetings</h3>
                        </div>
                        <div className="card-content">
                            {upcomingMeetings.length > 0 ? (
                                <div className="meetings-list">
                                    {upcomingMeetings.map(meeting => (
                                        <div key={meeting.id} className="meeting-item">
                                            <div className="meeting-title">{meeting.title}</div>
                                            <div className="meeting-time">{meeting.day} at {meeting.time}</div>
                                            <div className="meeting-participants">{meeting.participants} participants</div>
                                        </div>
                                    ))}
                                </div>
                            ) : (
                                <div className="empty-state">
                                    <p>No upcoming meetings</p>
                                </div>
                            )}
                        </div>
                    </div>

                    {/* Team Members Card */}
                    <div className="dashboard-card users">
                        <div className="card-header">
                            <h3>Team Members</h3>
                        </div>
                        <div className="card-content">
                            <div className="users-grid">
                                {teamMembers.map(member => (
                                    <div key={member.id} className="user-card">
                                        <div className="user-avatar">
                                            {member.name.split(' ').map(n => n[0]).join('')}
                                        </div>
                                        <div className="user-name">{member.name}</div>
                                        <div className="user-status">{member.status}</div>
                                    </div>
                                ))}
                            </div>
                        </div>
                    </div>

                    {/* Quick Stats Card */}
                    <div className="dashboard-card stats">
                        <div className="card-header">
                            <h3>Quick Stats</h3>
                        </div>
                        <div className="card-content">
                            <div className="stat-item">
                                <div className="stat-number">{upcomingMeetings.length}</div>
                                <div className="stat-label">Meetings Today</div>
                            </div>
                            <div className="stat-item">
                                <div className="stat-number">{teamMembers.filter(m => m.status === 'Online').length}</div>
                                <div className="stat-label">Online Members</div>
                            </div>
                        </div>
                    </div>
                </div>

                {/* Quick Actions */}
                <div className="quick-actions">
                    <Link to="/create-meeting" className="action-button">
                        Create New Meeting
                    </Link>
                    <Link to="/schedule" className="action-button secondary">
                        View Schedule
                    </Link>
                </div>
            </div>
        </div>
    );
}

export default Dashboard;