import { useState } from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
  Outlet,
  Navigate,
  useParams,
  useNavigate } from 'react-router-dom';
  import { useAuth0 } from '@auth0/auth0-react';

import MainAuth from "./components/MainAuth";
import JoinMeetingPage from "./components/JoinMeetingPage";
import CreateMeetingPage from "./components/CreateMeetingPage";
import Profile from "./components/Profile";
import Logout from "./components/LogoutButton"

function Home() {
  const { isAuthenticated, user } = useAuth0();

  return isAuthenticated && user ? (
    <div>
      <Logout/>
      <Profile/>
    </div>
  ) : (
    <div>
      <MainAuth/>
    </div>
  );
}

function NoMatch() {
  return (
    <div style={{ padding: 20 }}>
      <h2>404: Page Not Found</h2>
    </div>
  );
}

function AppLayout() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/join_meeting" element={<JoinMeetingPage/>} />
        <Route path="/create_meeting" element={<CreateMeetingPage/>} />
        <Route path="*" element={<NoMatch />} />
      </Routes>
    </>
  );
}

function App() {
  return (
    <Router>
        <AppLayout/>
    </Router>
  );
}

export default App;