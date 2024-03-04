import React from 'react';
import './styles.css'; // Import your CSS file
import LoginButton from "./TopContainerButtons"
import LogoutButton from "./LogoutButton"
import Profile from "./profile"

function TopContainer() {
  return (
    <div className="top-container">
        <LoginButton/>
        <LogoutButton/>
        <Profile/>
    </div>
  );
}

export default TopContainer;
