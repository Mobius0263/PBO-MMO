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
                    <div className="search-container">
                        <input type="text" placeholder="Search..." className="search