import { Routes, Route } from "react-router-dom";
import HomeScreen from "./HomeScreen";
import CallScreenWithProvider from "./CallScreen";
import AgoraUI from "./AgoraCalls";

function RouteList() {
  return (
    <Routes>
      <Route path="/" element={<HomeScreen />} />
      <Route path="/call/:room" element={<AgoraUI />} /> 
    </Routes>
  );
}

export default RouteList;
