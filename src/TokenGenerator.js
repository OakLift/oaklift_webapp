function generateToken(userID, channelName, expiration) {
  return fetch("http://ec2-18-225-54-2.us-east-2.compute.amazonaws.com/getToken", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      tokenType: "rtc",
      channel: channelName,
      role: "publisher",
      uid: userID,
      expire: expiration,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log(data)
      return data.token
    });
}

export default generateToken;
