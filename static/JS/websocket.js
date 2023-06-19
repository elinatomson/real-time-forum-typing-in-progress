let socket = null;

export function webSoc(nicknameTo, nicknameFrom) {
  if (socket && socket.readyState === WebSocket.OPEN) {
    // WebSocket connection already open
    return;
  }

  // Create a new WebSocket connection
  socket = new WebSocket("ws://localhost:8080/ws");

  // Get the necessary DOM elements
  const messageBox = document.getElementById("message-box");
  const messageInput = document.getElementById("message-input");
  const sendButton = document.getElementById("send-button");

  // Event listener for WebSocket connection open
  socket.addEventListener("open", () => {
    console.log("WebSocket connection established.");
  });

  // Event listener for WebSocket connection close
  socket.addEventListener("close", () => {
    console.log("WebSocket connection closed.");
  });

  // Event listener for WebSocket errors
  socket.addEventListener("error", (error) => {
    console.error("WebSocket error:", error);
  });

  // Event listener for receiving messages from the server
  socket.addEventListener("message", (event) => {
    const message = JSON.parse(event.data);
    // Handle the received message
    handleMessage(message);
  });

  // Event listener for the send button
  sendButton.addEventListener("click", (event) => {
    event.preventDefault(); // Prevent page reload
    const message = messageInput.value;
    // Send the message to the server
    sendMessage(message, nicknameTo, nicknameFrom);
    // Clear the message input
    messageInput.value = "";
  });

  // Function to handle the received message
  function handleMessage(message) {
    const messageText = message.message;
    // Display the message in the textbox or wherever you want to show it
    messageBox.value += messageText + "\n";
  }
}

export function sendMessage(message, nicknameTo, nicknameFrom) {
  if (!socket || socket.readyState !== WebSocket.OPEN) {
    console.error("WebSocket connection not open.");
    return;
  }

  const data = {
    message: message,
    nicknamefrom: nicknameFrom,
    nicknameto: nicknameTo
  };

  socket.send(JSON.stringify(data));

  fetch("/message", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  })
    .then(response => {
      if (response.ok) {
        console.log("Message sent successfully");
      } else {
        return response.text();
      }
    })
    .then(errorMessage => {
      if (errorMessage) {
        console.error("Error sending message:", errorMessage);
      }
    })
    .catch(error => {
      console.error("An error occurred while sending messages:", error);
    });
}
