import { displayErrorMessage } from './errormessage.js';

let socket;
let isTyping = false; //if the user is typing or not
let typingTimeout = null; //how quickly the typing message disappears

export function webSoc(nicknameTo, nicknameFrom) {
  socket = new WebSocket("ws://localhost:8080/ws");

  const messageBox = document.getElementById("message-box");
  const messageInput = document.getElementById("message-input");
  const sendButton = document.getElementById("send-button");
  const typingStatus = document.getElementById("typing-status");

  socket.addEventListener("open", () => {
    console.log("WebSocket connection established.");
  });

  const logOut = document.getElementById("logout-button");
  logOut.addEventListener("click", () => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.close();
      console.log("WebSocket connection closed.");
    }
  });

  socket.addEventListener("error", (error) => {
    console.error("WebSocket error:", error);
  });

  socket.addEventListener("message", (event) => {
    const message = JSON.parse(event.data);
    handleMessage(message);
  });

  sendButton.addEventListener("click", (event) => {
    event.preventDefault();
    const message = messageInput.value;
    sendMessage(message, nicknameTo, nicknameFrom);
    messageInput.value = "";
  });

  messageInput.addEventListener("keydown", (event) => {
    if (event.keyCode === 13) {
      event.preventDefault();
      const message = messageInput.value;
      sendMessage(message, nicknameTo, nicknameFrom);
      messageInput.value = "";
    }
  });

  //eventlistener for the typing message 
  messageInput.addEventListener("input", () => {
    const message = messageInput.value.trim();
    if (message !== "") {
        isTyping = true;
        sendTyping(true, nicknameTo);
    } else {
        isTyping = false;
        sendTyping(false, nicknameTo);
    }
  });

  function handleMessage(message) {
    let nickname = nicknameTo;
    if (message.nicknameto === nicknameTo) {
      nickname = nicknameFrom;
    }
    const messageText = message.message;
    const formattedTime = new Date(message.date).toLocaleString();
  
    //checks if the received message is not a "typing" message (message.typing is false).
    if (!message.typing) {
      if (messageText !== "") {
        messageBox.value += `${formattedTime} - ${nickname}: ${messageText}\n`;
      }
      messageBox.scrollTop = messageBox.scrollHeight;
    //if the received message is a "typing" message (message.typing is true), then displaying the typing text for the recipient
    } else {
      if (nickname === nicknameTo) {
        typingStatus.textContent = `${nickname} is typing...`;
        clearTimeout(typingTimeout); 
        typingTimeout = setTimeout(() => {
          isTyping = false;
          sendTyping(false, nicknameTo);
          typingStatus.textContent = "";
        }, 1000); 
      } else {
        clearTimeout(typingTimeout);
        if (isTyping) {
          isTyping = false;
          sendTyping(false, nicknameTo);
          typingStatus.textContent = "";
        }
      }
    }
  }

  //sending a "typing" message to the server over the WebSocket connection
  function sendTyping(typing, nicknameTo) {
    if (!socket || socket.readyState !== WebSocket.OPEN) {
      console.error("WebSocket connection not open.");
      return;
    }

    const data = {
      typing: typing,
      nicknamefrom: nicknameFrom,
      nicknameto: nicknameTo
    };
    socket.send(JSON.stringify(data));
  }

  function sendMessage(message, nicknameTo, nicknameFrom) {
    if (!socket || socket.readyState !== WebSocket.OPEN) {
      console.error("WebSocket connection not open.");
      return;
    }

    const date = new Date();
    const data = {
      message: message,
      nicknamefrom: nicknameFrom,
      nicknameto: nicknameTo,
      date: date
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
          displayErrorMessage(errorMessage);
        }
      })
      .catch(error => {
        displayErrorMessage(`An error occurred while sending message: ${error.message}`);
      });
  }
}
