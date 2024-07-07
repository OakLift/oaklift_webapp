import React from 'react';
import { useAuth0 } from "@auth0/auth0-react";
import LoginButton from "./LoginButton";
import LogoutButton from "./LogoutButton";
import './MainAuth.css'; // Import the CSS file

const MainAuth = () => {
    const { isLoading, error } = useAuth0();
    return (
        <main className="column">

          <div className="container">
            <h1> Login </h1>

            <form>
              <label htmlFor="email">Email</label>
              <input 
                type="email"
                id="email"
              />

              <label htmlFor="password">Password</label>
              <input 
                type="password"
                id="password"
              />

            </form>

     


            <button>Sign In</button>


            {error && <p>AuthenticationError</p>}
            {!error && isLoading && <p>Loading...</p>}
            {!error && !isLoading && (
              <>
                <LoginButton/>
              </>
            )}
          </div>
      
        </main>
    );
};

export default MainAuth;