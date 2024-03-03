import { Routes, Route } from "react-router-dom";
import HomeScreen from "./HomeScreen";
import CallScreenWithProvider from "./CallScreen";

function RouteList() {
  return (
    <Routes>
      <Route path="/" element={<HomeScreen />} />
      <Route path="/call/:room" element={<CallScreenWithProvider />} /> 
    </Routes>
  );
}

export default RouteList;
