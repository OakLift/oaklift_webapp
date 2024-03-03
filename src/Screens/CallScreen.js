import AgoraRTC, {
  useJoin,
  LocalVideoTrack,
  RemoteUser,
  useLocalCameraTrack,
  useLocalMicrophoneTrack,
  usePublish,
  useRemoteUsers,
  AgoraRTCProvider,
  useRTCClient,
  useClientEvent,
} from "agora-rtc-react";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./CallScreen.css"

function CallScreen(props) {
  const navigate = useNavigate();
  const { isLoading: isLoadingCam, localCameraTrack } = useLocalCameraTrack();
  const { isLoading: isLoadingMic, localMicrophoneTrack } =
    useLocalMicrophoneTrack();
  usePublish([localMicrophoneTrack, localCameraTrack]);
  useJoin({
    appid: "cf36a471064d48c38b5b4a05e89d7fc0",
    channel: "testingroom",
    token:
      "007eJxTYHg4bbeXbexdy7PsLmHrA92npDnOvrVM8OWk1X7Cm1iS+wMUGJLTjM0STcwNDcxMUkwsko0tkkyTTBINTFMtLFPM05INklWepDYEMjJM4bjDyMgAgSA+N0NJanFJZl56UX5+LgMDAPMnIYM=",
  });

  const [endCall, setEndCall] = useState(false);

  const remoteUsers = useRemoteUsers();
  const agoraEngine = useRTCClient();
  useClientEvent(agoraEngine, "user-joined", (user) => {
    console.log("The user", user.uid, " has joined the channel");
  });

  useClientEvent(agoraEngine, "user-left", (user) => {
    console.log("The user", user.uid, " has left the channel");
  });

  useClientEvent(agoraEngine, "user-published", (user, mediaType) => {
    console.log("The user", user.uid, " has published media in the channel");
  });

  useEffect(() => {
    if (endCall) {
      navigate("/");
      return () => {
        localCameraTrack?.close();
        localMicrophoneTrack?.close();
      };
    }
  }, [localCameraTrack, localMicrophoneTrack, endCall, navigate]);
  const deviceLoading = isLoadingMic || isLoadingCam;
  return (
    <div
      style={{ display: "flex", flexDirection: "column", alignItems: "center" }}
    >
      {deviceLoading ? (
        <div>Loading devices...</div>
      ) : (
        <div id="videos" className="video_screen">
          {/* Render the local video track */}
          <div
            className="vid"
            style={{ height: 300, width: 400, margin: "10px" }}
          >
            <LocalVideoTrack track={localCameraTrack} play={true} />
          </div>
          {/* Render remote users' video and audio tracks */}
          {remoteUsers.map((remoteUser) => (
            <div
              className="vid"
              style={{ height: 300, width: 400, margin: "10px" }}
              key={remoteUser.uid}
            >
              <RemoteUser user={remoteUser} playVideo={true} playAudio={true} />
            </div>
          ))}
        </div>
      )}
      <button onClick={() => setEndCall(true)} className="end_call_button">
        End Call
      </button>
    </div>
  );
}

function CallScreenWithProvider() {
  const client = useRTCClient(
    AgoraRTC.createClient({ codec: "vp8", mode: "rtc" })
  );
  return (
    <AgoraRTCProvider client={client}>
      <CallScreen client={client} />
    </AgoraRTCProvider>
  );
}

export default CallScreenWithProvider;
