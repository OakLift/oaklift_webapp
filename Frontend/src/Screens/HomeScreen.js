import React, { useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import "./HomeScreen.css";
import generateToken from "../TokenGenerator";

function HomeScreen() {
  const [room, setRoom] = useState("");
  const navigate = useNavigate();
  const location = useLocation();
  const handleRoom = () => {
    const username = localStorage.getItem("userID");
    if (username == null || username == undefined || username === "") {
      navigate("/");
    }
    generateToken(username, room, 36000).then((local_token) => {
      localStorage.setItem("token_" + username, local_token);
      localStorage.setItem("channel_" + username, room);
      navigate(`/call?userID=${username}`);
    });
  };
  console.log(location.pathname)
  return (
    <form>
      <div>
        <label htmlFor="room">Room ID</label>
        <input
          value={room}
          title="room"
          onChange={(e) => setRoom(e.target.value)}
        />
        {location.pathname.includes("join") ? (
          <input
            type="button"
            name="submit"
            value="Join Room"
            onClick={handleRoom}
          />
        ) : (
          <input
            type="button"
            name="submit"
            value="Create Room"
            onClick={handleRoom}
          />
        )}
      </div>
    </form>
  );
}

export default HomeScreen;
