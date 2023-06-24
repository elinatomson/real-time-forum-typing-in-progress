import { webSoc } from './websocket.js';
import { loadUserPage } from './userpage.js';
 
function handleUserClick(user) {
  const formContainer = document.getElementById('formContainer');
  formContainer.innerHTML = 
  `Chat with ${user}
  <div>
    <textarea id="message-box" rows="10" cols="50" readonly></textarea>
    <div>
      <input type="text" id="message-input">
      <button id="send-button">Send</button>
    </div>
  </div>
  <p class="align">
    <input id="back" class="buttons" type="button" value="Back to main page">
  </p>`
;

  //to mark all previously received messages as read when the user starts typing a new message.
  const messageInput = document.getElementById('message-input');
  messageInput.addEventListener('click', () => {
    messagesAsRead(user);
  });

  displayMessages(user);
  //if you are logged in, then clicking on the Forum name, you will see the userPage as a mainpage
  var mainPage = document.getElementById('mainpage');
  mainPage.addEventListener('click', function(event) {
      event.preventDefault();
      loadUserPage();
  });
  //by clicking on the "Back to main page" button, you will see the userPage as a mainpage
  var backButton = document.getElementById('back'); 
  backButton.addEventListener('click', function(event) {
      event.preventDefault();
      loadUserPage(); //TODO! if you click it for the first time, then everything in the userpage is somehow dublicated. 
  });
}
  
function attachUserClickListeners() {
  const users = document.querySelectorAll('.user');
  users.forEach(user => {
    user.addEventListener('click', () => {
      handleUserClick(user.dataset.user);
      webSoc(user.dataset.user)
    });
  });
}

export  function usersForChat() {
  //to get online and offline users list
    fetch('/users')
    .then(response => response.json())
    .then(users => {

      const userListContainer = document.getElementById('user-list-container');
      userListContainer.className = 'users-container';

      users.forEach(user => {
        const userItem = document.createElement('div');
        userItem.className = 'user';
        userItem.textContent = user.nickname;
        // Check if the user has unread messages
        if (user.unreadMessages) {
          const unreadIndicator = document.createElement('span');
          unreadIndicator.className = 'unread-indicator';
          unreadIndicator.textContent = '(Unread Messages)';
          userItem.appendChild(unreadIndicator);
        }

        userItem.dataset.user = user.nickname;
        userItem.classList.add(user.online ? 'online' : 'offline');
        userListContainer.appendChild(userItem);
      });
      attachUserClickListeners();
    })
    .catch(error => {
      var formContainer = document.getElementById('formContainer');
      var errorContainer = document.createElement('div');
      errorContainer.className = 'message';
      errorContainer.textContent = error.message;
      formContainer.appendChild(errorContainer);
    });
}

export function displayMessages(nicknameTo) {
  const messageBox = document.getElementById("message-box");
  const pageSize = 10;
  let currentPage = 1;

  function loadMessages(page) {
    fetch(`/messages?nicknameTo=${nicknameTo}&page=${page}&pageSize=${pageSize}`)
      .then(response => response.json())
      .then(messages => {
        if (messages && messages.length > 0) {
          const newMessages = messages.map(message => {
            var formattedDate = new Date(message.date).toLocaleString();
            //the way to display every message in the conversation history
            return `${formattedDate} - ${message.nicknamefrom}: ${message.message}`;
          });
          const messagesToDisplay = newMessages.sort((a, b) => a.date - b.date).reverse().join("\n") + "\n"; //sord chronologically and reverse the order of messages so the newest ones are at the bottom
          messageBox.value = messagesToDisplay + messageBox.value; // prepend messages

          //if you are opening the chat, meaning it is the first page
          if (page === 1) {
            //set the scroll position to the bottom
            messageBox.scrollTop = messageBox.scrollHeight;
            messagesAsRead(nicknameTo);
          }
        }
      })
      .catch(error => {
        console.error('An error occurred while loading messages:', error);
      });
  }

  loadMessages(currentPage);

  //reload the last 10 messages and when scrolled up to see more messages you will see 10 more and so on
  messageBox.addEventListener('scroll', () => {
    if (messageBox.scrollTop === 0) {
      currentPage++;
      loadMessages(currentPage);
    }
  });
}

function messagesAsRead(nicknameTo) {
  fetch(`/messages/markAsRead?nicknameTo=${nicknameTo}`)
    .then(response => {
      if (response.ok) {
        console.log('All messages sent for the logged in user marked as read.');
      } else {
        console.error('Failed to mark messages as read.');
      }
    })
    .catch(error => {
      console.error('An error occurred while marking messages as read:', error);
    });
}

function unreadMessages() {
  fetch('/messages/unread')
    .then(response => response.json())
    .then(messages => {
      if (messages && messages.length > 0) {
        messages.forEach(message => {
          const sender = message.nicknamefrom;
          const date = new Date(message.date).toLocaleString();
          const content = message.message;

          // Display the unread message or handle it as required
          console.log(`Unread Message - Sender: ${sender}, Date: ${date}, Message: ${content}`);
        });
      } else {
        // No unread messages
        console.log('No unread messages');
      }
    })
    .catch(error => {
      console.error('An error occurred while fetching unread messages:', error);
    });
}

unreadMessages();
