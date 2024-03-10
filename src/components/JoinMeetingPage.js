import React from 'react';

const JoinMeetingPage = () => {
    return (
        <div>
            <label>
                Meeting Link:
                <input type="text"></input>
            </label>

            <label>
                Token:
                <input type="text"></input>
            </label>

            <label>
                Channel ID:
                <input type="text"></input>
            </label>
        </div>
    );
};

export default JoinMeetingPage;