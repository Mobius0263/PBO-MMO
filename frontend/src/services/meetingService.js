import axios from 'axios';

const API_URL = 'http://localhost:8080';

// Create axios instance
const api = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

// Add token to requests if it exists
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

// Create a new meeting
export const createMeeting = async (meetingData) => {
    try {
        const response = await api.post('/api/meetings', meetingData);
        return response.data;
    } catch (error) {
        console.error('Error creating meeting:', error);
        throw error;
    }
};

// Get all meetings for a user
export const getMeetings = async (userId, date) => {
    try {
        let url = '/api/meetings';
        const params = {};
        
        if (userId) {
            params.userId = userId;
        }
        
        if (date) {
            params.date = formatDate(date);
        }
        
        const response = await api.get(url, { params });
        return response.data;
    } catch (error) {
        console.error('Error fetching meetings:', error);
        throw error;
    }
};

// Get a specific meeting by ID
export const getMeetingById = async (id) => {
    try {
        const response = await api.get(`/api/meetings/${id}`);
        return response.data;
    } catch (error) {
        console.error(`Error fetching meeting ${id}:`, error);
        throw error;
    }
};

// Update a meeting
export const updateMeeting = async (id, meetingData) => {
    try {
        const response = await api.put(`/api/meetings/${id}`, meetingData);
        return response.data;
    } catch (error) {
        console.error(`Error updating meeting ${id}:`, error);
        throw error;
    }
};

// Delete a meeting
export const deleteMeeting = async (id) => {
    try {
        const response = await api.delete(`/api/meetings/${id}`);
        return response.data;
    } catch (error) {
        console.error(`Error deleting meeting ${id}:`, error);
        throw error;
    }
};

// Helper function to format date to YYYY-MM-DD
const formatDate = (date) => {
    const d = new Date(date);
    const year = d.getFullYear();
    const month = String(d.getMonth() + 1).padStart(2, '0');
    const day = String(d.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
};