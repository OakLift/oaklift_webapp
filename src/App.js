import React from 'react';
import TopContainer  from "./components/TopContainer"
import './App.css'; 

function App() {
  return (
    <div className="app-container gradient-background"> 
      <header className="header">
      </header>
      <main className="content">
        <TopContainer/>
      </main>
    </div>
  );
}

export default App;
