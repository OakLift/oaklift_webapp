import { useAuth0 } from '@auth0/auth0-react';
import { Link } from 'react-router-dom';
import './Profile.css';

const Profile = () => {
  const { user, isAuthenticated } = useAuth0();

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
          <Link to="/join_meeting">
            <button>Join Meeting</button>
          </Link>
          <Link to="/create_meeting">
            <button>Create Meeting</button>
          </Link>
        </div>
      </div>
    )
  );
};

export default Profile;