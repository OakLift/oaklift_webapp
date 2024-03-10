import { useAuth0 } from '@auth0/auth0-react';
import { Link } from 'react-router-dom';
import './Profile.css';
import { useEffect } from 'react';

const Profile = () => {
  const { user, isAuthenticated } = useAuth0();

  useEffect(() => {
    if(isAuthenticated) {
      localStorage.setItem("userID", user.name)
    }
  },[isAuthenticated, user])
  return (
    isAuthenticated && (
      <div> 

        <div className="user-profile">
          <div className="profile-picture">
            <img src={user.picture} alt="Profile Picture" />
          </div>

          <div className="profile-info">
            <h2>{user.name}</h2>
            <p>Software Engineer</p>
            <p>Location: San Jose, CA</p>
            <p>Email: {user.email}</p>
          </div>
        </div>
        
        <div className="button-container">
          <Link to="/home/join">
            <button>Join Meeting</button>
          </Link>
          <Link to="/home/create">
            <button>Create Meeting</button>
          </Link>
        </div>
      </div>
    )
  );
};

export default Profile;