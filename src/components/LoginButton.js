import { useAuth0 } from '@auth0/auth0-react';
import './LoginButton.css'; // Import the CSS file

const LoginButton = () => {
    const { loginWithRedirect, isAuthenticated } = useAuth0();
    return (
        !isAuthenticated && (
            <div className="login-button-container"> {/* Added container */}
                <button className="login-button" onClick={() => loginWithRedirect()}>
                    Sign In
                </button>
            </div>
        )
    )
}

export default LoginButton;