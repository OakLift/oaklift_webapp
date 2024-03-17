function generateToken(userID, channelName, expiration) {
  return fetch("https://api.ak-coffee.info/getToken", {
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
