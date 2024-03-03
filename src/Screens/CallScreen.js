import AgoraRTC from "agora-rtc-sdk-ng";
import {
  useJoin,
  LocalVideoTrack,
  RemoteUser,
  useLocalCameraTrack,
  useLocalMicrophoneTrack,
  usePublish,
  useRemoteUsers,
  AgoraRTCProvider,
  useRTCClient
} from "agora-rtc-react";
import { useState } from "react";

function CallScreen() {
  
  const remoteUsers = useRemoteUsers();
  const { isLoading: isLoadingCam, localCameraTrack } = useLocalCameraTrack();
  const queryParameters = new URLSearchParams(window.location.search);
  const { isLoading: isLoadingMic, localMicrophoneTrack } =
    useLocalMicrophoneTrack();
  const [joined, setJoined] = useState(false);

  usePublish([localMicrophoneTrack, localCameraTrack]);
  console.log(queryParameters.get("userID"))
  const deviceLoading = isLoadingMic || isLoadingCam;
  useJoin({
    appid: "cf36a471064d48c38b5b4a05e89d7fc0",
    channel: "testingroom",
    token:
      "007eJxTYHBuYkxcuYn323Epz+15NV3R38+cZWKSrlC4sNOy1qpBlUuBITnN2CzRxNzQwMwkxcQi2dgiyTTJJNHANNXCMsU8LdlAifFJakMgI8OjSDZGRgYIBPG5GUpSi0sy89KL8vNzGRgAemYfxQ==",
  });
  return (
    <div>
      {deviceLoading ? (
        <div>Loading devices...</div>
      ) : (
        <div id="videos">
          <div className="vid" style={{ height: 300, width: 600 }}>
            <LocalVideoTrack track={localCameraTrack} play={true} />
          </div>
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
      <CallScreen/>
    </AgoraRTCProvider>
  )
}

export default CallScreenWithProvider
