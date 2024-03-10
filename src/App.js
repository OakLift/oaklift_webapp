import { useState } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";

import MainAuth from "./components/MainAuth";
import Profile from "./components/Profile";
import Logout from "./components/LogoutButton";
import HomeScreen from "./Screens/HomeScreen";
import CallScreenWithProvider from "./Screens/CallScreen";

function Home() {
  const { isAuthenticated, user } = useAuth0();

  return isAuthenticated && user ? (
    <div>
      <Logout />
      <Profile />
    </div>
  ) : (
    <div>
      <MainAuth />
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
        <Route path="/" element={<Home />} />
        <Route path="/home/:mode" element={<HomeScreen />} />
        <Route path="/call" element={<CallScreenWithProvider />} />
        <Route path="*" element={<NoMatch />} />
      </Routes>
    </>
  );
}

function App() {
  return (
    <Router>
      <AppLayout />
    </Router>
  );
}

export default App;
