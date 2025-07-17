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

// Get all users
export const getUsers = async () => {
    try {
        console.log('Fetching users from:', `${API_URL}/users`);
        const response = await api.get('/users');
        
        // Add baseURL to profileImage if exists
        const users = response.data.map(user => {
            if (user.profileImage && !user.profileImage.startsWith('http')) {
                user.profileImage = API_URL + user.profileImage;
            }
            return user;
        });
        
        return users;
    } catch (error) {
        console.error('Error fetching users:', error);
        throw error;
    }
};

// Get user by ID
export const getUserById = async (id) => {
    try {
        const response = await api.get(`/api/users/${id}`);
        
        // Add baseURL to profileImage if exists
        if (response.data.profileImage && !response.data.profileImage.startsWith('http')) {
            response.data.profileImage = API_URL + response.data.profileImage;
        }
        
        return response.data;
    } catch (error) {
        console.error(`Error fetching user ${id}:`, error);
        throw error;
    }
};

// Update user
export const updateUser = async (id, userData) => {
    try {
        const response = await api.put(`/api/users/${id}`, userData);
        return response.data;
    } catch (error) {
        console.error(`Error updating user ${id}:`, error);
        throw error;
    }
};

// Upload profile image
export const uploadProfileImage = async (id, formData) => {
    try {
        // Use FormData for file uploads
        const config = {
            headers: {
                'Content-Type': 'multipart/form-data',
            },
        };
        
        const token = localStorage.getItem('token');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        
        const response = await axios.post(
            `${API_URL}/api/users/${id}/profile-image`, 
            formData,
            config
        );
        
        return response.data;
    } catch (error) {
        console.error(`Error uploading profile image for user ${id}:`, error);
        throw error;
    }
};