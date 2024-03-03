import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./HomeScreen.css";

function HomeScreen() {
  const [room, setRoom] = useState("");
  const [username, setUsername] = useState("");

  return (
    <form>
      <div>
        <label htmlFor="username">Username</label>
        <input
          value={username}
          title="username"
          onChange={(e) => setUsername(e.target.value)}
        />
        <label htmlFor="room">Room</label>
        <input
          value={room}
          title="room"
          onChange={(e) => setRoom(e.target.value)}
        />
        <Link to={`/call/${room}?userID=${username}`}>
          <input type="submit" name="submit" value="Join Room" />
        </Link>
      </div>
    </form>
  );
}



export default HomeScreen;
