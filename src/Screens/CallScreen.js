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

function CallScreen(props) {
  const { isLoading: isLoadingCam, localCameraTrack } = useLocalCameraTrack();
  const queryParameters = new URLSearchParams(window.location.search);
  const { isLoading: isLoadingMic, localMicrophoneTrack } =
    useLocalMicrophoneTrack();
  usePublish([localMicrophoneTrack, localCameraTrack]);
  useJoin({
    appid: "cf36a471064d48c38b5b4a05e89d7fc0",
    channel: "testingroom",
    token:
      "007eJxTYHg4bbeXbexdy7PsLmHrA92npDnOvrVM8OWk1X7Cm1iS+wMUGJLTjM0STcwNDcxMUkwsko0tkkyTTBINTFMtLFPM05INklWepDYEMjJM4bjDyMgAgSA+N0NJanFJZl56UX5+LgMDAPMnIYM=",
  });

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

  console.log(queryParameters.get("userID"));
  const deviceLoading = isLoadingMic || isLoadingCam;
  return (
    <div>
      {deviceLoading ? (
        <div>Loading devices...</div>
      ) : (
        <div id="videos">
          {/* Render the local video track */}
          <div className="vid" style={{ height: 300, width: 600 }}>
            <LocalVideoTrack track={localCameraTrack} play={true} />
          </div>
          {/* Render remote users' video and audio tracks */}
          {remoteUsers.map((remoteUser) => (
            <div
              className="vid"
              style={{ height: 300, width: 600 }}
              key={remoteUser.uid}
            >
              <RemoteUser user={remoteUser} playVideo={true} playAudio={true} />
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

function CallScreenWithProvider() {
  const client = useRTCClient(
    AgoraRTC.createClient({ codec: "vp8", mode: "rtc" })
  );
  return (
    <AgoraRTCProvider client={client}>
      <CallScreen client={client}/>
    </AgoraRTCProvider>
  );
}

export default CallScreenWithProvider;
