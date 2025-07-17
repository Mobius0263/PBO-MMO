import { useState, useContext, useEffect } from 'react';
import { useNavigate, Link, useLocation } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';
import { login } from '../services/authService';
import '../styles/Auth.css';

function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    
    const { loginUser } = useContext(AuthContext);
    const navigate = useNavigate();
    const location = useLocation();
    
    useEffect(() => {
        // Check if user was redirected from register page
        if (location.state?.registrationSuccess) {
            setSuccessMessage('Registrasi berhasil! Silakan login.');
        }
    }, [location]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setIsLoading(true);
        setError('');

        try {
            // Validate inputs
            if (!email || !password) {
                throw new Error('Email dan password harus diisi');
            }

            // Call login API
            const response = await login({ email, password });
            
            // Update auth context with user data
            loginUser(response.user);
            
            // Redirect to dashboard
            navigate('/dashboard');
        } catch (err) {
            console.error('Login failed:', err);
            setError(err.message || 'Login gagal. Silakan periksa email dan password Anda.');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="auth-container">
            <div className="auth-card">
                <h2>Login</h2>
                
                {error && (
                    <div className="error-message">
                        {error}
                    </div>
                )}
                
                {successMessage && (
                    <div className="success-message">
                        {successMessage}
                    </div>
                )}
                
                <form onSubmit={handleSubmit} className="auth-form">
                    <div className="form-group">
                        <label htmlFor="email">Email</label>
                        <input
                            type="email"
                            id="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            placeholder="Email"
                            required
                        />
                    </div>
                    
                    <div className="form-group">
                        <label htmlFor="password">Password</label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            placeholder="Password"
                            required
                        />
                    </div>
                    
                    <button 
                        type="submit" 
                        className="auth-button"
                        disabled={isLoading}
                    >
                        {isLoading ? 'Loading...' : 'Login'}
                    </button>
                </form>
                
                <p className="auth-redirect">
                    Belum punya akun? <Link to="/register">Register</Link>
                </p>
            </div>
        </div>
    );
}

export default Login;