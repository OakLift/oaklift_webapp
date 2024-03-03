import { useState } from "react";
import AgoraUIKit from "agora-react-uikit";

function AgoraUI() {
  const [videoCall, setVideoCall] = useState(true);
  const rtcProps = {
    appId: "cf36a471064d48c38b5b4a05e89d7fc0",
    channel: "testingroom",
    token: "007eJxTYHg4bbeXbexdy7PsLmHrA92npDnOvrVM8OWk1X7Cm1iS+wMUGJLTjM0STcwNDcxMUkwsko0tkkyTTBINTFMtLFPM05INklWepDYEMjJM4bjDyMgAgSA+N0NJanFJZl56UX5+LgMDAPMnIYM=",
  };
  const callbacks = {
    EndCall: () => setVideoCall(false),
  };
  const rtmProps = {};
  const styleProps = {};
  return videoCall ? (
    <div style={{ display: "flex", width: "100vw", height: "100vh" }}>
      <AgoraUIKit
        rtcProps={rtcProps}
        callbacks={callbacks}
        rtmProps={rtmProps}
        styleProps={styleProps}
      />
    </div>
  ) : (
    <h3 onClick={() => setVideoCall(true)}>Start Call</h3>
  );
}
export default AgoraUI;
