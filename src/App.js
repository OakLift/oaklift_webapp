import LoginButton from "./components/LoginButton"
import LogoutButton from "./components/LogoutButton"
import Profile from "./components/profile";

function App() {
  return (
    <main classname="column">
      <h1>Auth0 login</h1>
      <LoginButton/>
      <LogoutButton/>
      <Profile/>
    </main>
  );
}

export default App;
