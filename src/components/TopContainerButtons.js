import React from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import './TopContainerButtons.css'; // Import the CSS file

const LoginButton = () => {
    const { loginWithRedirect, isAuthenticated } = useAuth0();
    return (
        !isAuthenticated && (
            <div className="login-button-container">


                <div className="secondary-buttons-container">
                    <button className="secondary-button">About</button>
                    <button className="secondary-button">Blog</button>
                    <button className="secondary-button">Contact</button>
                </div>

                <button className="login-button" onClick={() => loginWithRedirect()}>
                    Sign In
                </button>
            </div>
        )
    )
}

export default LoginButton;
