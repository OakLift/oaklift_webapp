import LoginButton from "./components/LoginButton"
import LogoutButton from "./components/LogoutButton"
import Profile from "./components/profile";
import './App.css'; 

function App() {
  return (
    <div className="app-container">
      <header className="header">
        <h1>Auth0 Login</h1>
      </header>
      <main className="content">
        <LoginButton/>
        <LogoutButton/>
        <Profile/>
      </main>
    </div>
  );
}

export default App;
