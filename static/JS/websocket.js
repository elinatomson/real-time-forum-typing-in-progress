
    // Create a WebSocket connection
    const socket = new WebSocket("ws://localhost:8080/ws");
  
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
      const message = event.data;
      // Add the received message to the message box
      messageBox.value += message + "\n";
    });
  
    // Event listener for the send button
    sendButton.addEventListener("click", () => {
      const message = messageInput.value;
      // Send the message to the server
      socket.send(message);
      // Clear the message input
      messageInput.value = "";
    });
  
    // Event listener for WebSocket connection open
    socket.onopen = function () {
      console.log("WebSocket connection established.");
    };
  
    // Event listener for WebSocket messages received
    socket.onmessage = function (event) {
      // Handle the received message
      const message = event.data;
      // Perform the necessary actions based on the received message
      // For example, update the UI, display notifications, etc.
      console.log("Received message:", message);
    };
  
    // Function to send a message to the server
    function sendMessage(message) {
      socket.send(message);
    }
  
    /* 
    //for reusing in other files
    // newpost.js
  
    // Example: Send a message using the sendMessage function from websocket.js
    function sendPost(content) {
      const message = {
        type: "newPost",
        content: content
      };
      sendMessage(JSON.stringify(message));
    }
  
    help me debug these files so that an echo websocket would appear in the forum once I start the server at localhost
    */
