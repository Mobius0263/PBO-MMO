import { useState, useContext, useEffect } from 'react';
import { AuthContext } from '../context/AuthContext';
import Sidebar from './Sidebar';
import '../styles/Meetings.css';

function Meetings() {
    const { user } = useContext(AuthContext);
    const [meetings, setMeetings] = useState([]);
    const [showCreateModal, setShowCreateModal] = useState(false);
    const [newMeeting, setNewMeeting] = useState({
        title: '',
        description: '',
        date: '',
        time: '',
        participants: []
    });

    useEffect(() => {
        // TODO: Fetch meetings from API
        // For now, using dummy data
        setMeetings([
            {
                id: 1,
                title: 'Weekly Standup',
                description: 'Weekly team synchronization meeting',
                date: '2025-07-22',
                time: '09:00',
                participants: ['John Doe', 'Alice Smith', 'Bob Johnson'],
                status: 'upcoming'
            },
            {
                id: 2,
                title: 'Project Review',
                description: 'Review current project progress',
                date: '2025-07-21',
                time: '14:00',
                participants: ['John Doe', 'Alice Smith'],
                status: 'completed'
            }
        ]);
    }, []);

    const handleCreateMeeting = (e) => {
        e.preventDefault();
        // TODO: Submit to API
        const meeting = {
            id: meetings.length + 1,
            ...newMeeting,
            participants: newMeeting.participants.split(',').map(p => p.trim()),
            status: 'upcoming'
        };
        setMeetings([...meetings, meeting]);
        setNewMeeting({ title: '', description: '', date: '', time: '', participants: [] });
        setShowCreateModal(false);
    };

    const formatDate = (dateStr) => {
        const date = new Date(dateStr);
        return date.toLocaleDateString('en-US', { 
            weekday: 'long', 
            year: 'numeric', 
            month: 'long', 
            day: 'numeric' 
        });
    };

    return (
        <div className="page-container">
            <Sidebar />
            <div className="main-content">
                <div className="page-header">
                    <h1>Meetings</h1>
                    <button 
                        className="btn-primary"
                        onClick={() => setShowCreateModal(true)}
                    >
                        <i className="fas fa-plus"></i>
                        Schedule Meeting
                    </button>
                </div>
                
                <div className="meetings-content">
                    <div className="meetings-stats">
                        <div className="stat-card">
                            <div className="stat-icon">
                                <i className="fas fa-calendar-alt"></i>
                            </div>
                            <div className="stat-info">
                                <h3>{meetings.filter(m => m.status === 'upcoming').length}</h3>
                                <p>Upcoming Meetings</p>
                            </div>
                        </div>
                        
                        <div className="stat-card">
                            <div className="stat-icon">
                                <i className="fas fa-check-circle"></i>
                            </div>
                            <div className="stat-info">
                                <h3>{meetings.filter(m => m.status === 'completed').length}</h3>
                                <p>Completed</p>
                            </div>
                        </div>
                        
                        <div className="stat-card">
                            <div className="stat-icon">
                                <i className="fas fa-clock"></i>
                            </div>
                            <div className="stat-info">
                                <h3>2.5</h3>
                                <p>Avg Duration (hrs)</p>
                            </div>
                        </div>
                    </div>

                    <div className="meetings-list">
                        <h2>All Meetings</h2>
                        <div className="meetings-grid">
                            {meetings.map(meeting => (
                                <div key={meeting.id} className={`meeting-card ${meeting.status}`}>
                                    <div className="meeting-header">
                                        <h3>{meeting.title}</h3>
                                        <span className={`status-badge ${meeting.status}`}>
                                            {meeting.status}
                                        </span>
                                    </div>
                                    <p className="meeting-description">{meeting.description}</p>
                                    <div className="meeting-details">
                                        <div className="meeting-time">
                                            <i className="fas fa-calendar"></i>
                                            <span>{formatDate(meeting.date)}</span>
                                        </div>
                                        <div className="meeting-time">
                                            <i className="fas fa-clock"></i>
                                            <span>{meeting.time}</span>
                                        </div>
                                    </div>
                                    <div className="meeting-participants">
                                        <span>Participants: {meeting.participants.join(', ')}</span>
                                    </div>
                                    <div className="meeting-actions">
                                        <button className="btn-secondary">
                                            <i className="fas fa-edit"></i>
                                            Edit
                                        </button>
                                        <button className="btn-primary">
                                            <i className="fas fa-video"></i>
                                            Join
                                        </button>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                </div>

                {/* Create Meeting Modal */}
                {showCreateModal && (
                    <div className="modal-overlay" onClick={() => setShowCreateModal(false)}>
                        <div className="modal-content" onClick={e => e.stopPropagation()}>
                            <div className="modal-header">
                                <h2>Schedule New Meeting</h2>
                                <button 
                                    className="modal-close"
                                    onClick={() => setShowCreateModal(false)}
                                >
                                    <i className="fas fa-times"></i>
                                </button>
                            </div>
                            <form onSubmit={handleCreateMeeting} className="meeting-form">
                                <div className="form-group">
                                    <label>Meeting Title</label>
                                    <input
                                        type="text"
                                        value={newMeeting.title}
                                        onChange={(e) => setNewMeeting({...newMeeting, title: e.target.value})}
                                        required
                                    />
                                </div>
                                <div className="form-group">
                                    <label>Description</label>
                                    <textarea
                                        value={newMeeting.description}
                                        onChange={(e) => setNewMeeting({...newMeeting, description: e.target.value})}
                                        rows="3"
                                    />
                                </div>
                                <div className="form-row">
                                    <div className="form-group">
                                        <label>Date</label>
                                        <input
                                            type="date"
                                            value={newMeeting.date}
                                            onChange={(e) => setNewMeeting({...newMeeting, date: e.target.value})}
                                            required
                                        />
                                    </div>
                                    <div className="form-group">
                                        <label>Time</label>
                                        <input
                                            type="time"
                                            value={newMeeting.time}
                                            onChange={(e) => setNewMeeting({...newMeeting, time: e.target.value})}
                                            required
                                        />
                                    </div>
                                </div>
                                <div className="form-group">
                                    <label>Participants (comma-separated emails)</label>
                                    <input
                                        type="text"
                                        value={newMeeting.participants}
                                        onChange={(e) => setNewMeeting({...newMeeting, participants: e.target.value})}
                                        placeholder="john@example.com, alice@example.com"
                                    />
                                </div>
                                <div className="modal-actions">
                                    <button 
                                        type="button" 
                                        className="btn-secondary"
                                        onClick={() => setShowCreateModal(false)}
                                    >
                                        Cancel
                                    </button>
                                    <button type="submit" className="btn-primary">
                                        Schedule Meeting
                                    </button>
                                </div>
                            </form>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}

export default Meetings;
