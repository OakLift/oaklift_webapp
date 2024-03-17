import React from 'react';
import './JoinMeetingPage.css'; // Import the CSS file

const JoinMeetingPage = () => {
  return (
    <div className="join-meeting-page">
      <div>
        <label htmlFor="meetingLink">
          Meeting Link:
          <input type="text" id="meetingLink" />
        </label>
      </div>
      <div>
        <label htmlFor="token">
          Token:
          <input type="text" id="token" />
        </label>
      </div>
      <div>
        <label htmlFor="channelId">
          Channel ID:
          <input type="text" id="channelId" />
        </label>
      </div>
    </div>
  );
};

export default JoinMeetingPage;