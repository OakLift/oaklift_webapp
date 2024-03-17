import React from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import './LogoutButton.css';

const LogoutButton = () => {
  const { logout, isAuthenticated } = useAuth0();

  return (
    isAuthenticated && (
      <div className="logout-button-container">
        <button
          type="button"
          className="logout-button"
          onClick={() => logout({ returnTo: window.location.origin })}
        >
          Sign Out
        </button>
      </div>
    )
  );
};

export default LogoutButton;