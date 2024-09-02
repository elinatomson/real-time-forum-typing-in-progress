import { webSoc } from './websocket.js';
import { loadUserPage } from './userpage.js';
 
export function handleUserClick(user) {
  const formContainer = document.getElementById('formContainer');
  formContainer.innerHTML = `
  <div class="unread">Chat with ${user}</div>
  <div class="align">
    <textarea class="messagebox" id="message-box" rows="10" cols="45" readonly></textarea>
    <div>
      <input type="text" id="message-input" class="input">
      <div id="typing-status"></div>
      <button class="buttons" id="send-button">Send</button>
    </div>
  </div>
  <p class="align">
    <input id="back" class="buttons" type="button" value="Back to main page">
  </p>`
;

  const messageInput = document.getElementById('message-input');
  messageInput.addEventListener('click', () => {
    messagesAsRead(user);
  });

  displayMessages(user);

  var mainPage = document.getElementById('mainpage');
  mainPage.addEventListener('click', function(event) {
      event.preventDefault();
      loadUserPage()
      messagesAsRead(user);
      window.history.pushState({ page: 'userpage' }, '', '/');
  });

  var backButton = document.getElementById('back'); 
  backButton.addEventListener('click', function(event) {
      event.preventDefault();
      loadUserPage()
      messagesAsRead(user);
      window.history.pushState({ page: 'userpage' }, '', '/');
  });

  window.history.pushState({ page: 'chat', }, '', `/`);
  window.addEventListener('popstate', function () {
  loadUserPage()
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
            return `${formattedDate} - ${message.nicknamefrom}: ${message.message}`;
          });

          const messagesToDisplay = newMessages.sort((a, b) => a.date - b.date).reverse().join("\n") + "\n"; 
          messageBox.value = messagesToDisplay + messageBox.value; 

          if (page === 1) {
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

function messagesAsRead(nicknameFrom) {
  fetch(`/messages/mark-as-read?nicknameFrom=${nicknameFrom}`)
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

function unreadMessageCount(count, senders) {
  const unreadMessagesElement = document.getElementById("unread-messages");
  const senderNamesElement = document.createElement("span");
  senderNamesElement.id = "sender-names";
  
  //to get messages sender names and make them clickable
  senders.forEach((sender, index) => {
    const senderSpan = document.createElement("span");
    senderSpan.textContent = sender;
    senderSpan.classList.add("user");
    senderSpan.dataset.user = sender;
    senderSpan.addEventListener("click", () => {
      handleUserClick(sender);
      webSoc(sender);
    });
    senderNamesElement.appendChild(senderSpan);

    if (index < senders.length - 1) {
      const commaSpan = document.createElement("span");
      commaSpan.textContent = ", ";
      senderNamesElement.appendChild(commaSpan);
    }
  });
  
  unreadMessagesElement.innerHTML = `Unread messages: ${count} from `;
  unreadMessagesElement.appendChild(senderNamesElement);
}

export function unreadMessages() {
  fetch('/messages/unread')
    .then(response => response.json())
    .then(messages => {
      if (messages && messages.length > 0) {
        //using a Set to store unique sender names
        const senders = new Set(); 

        messages.forEach(message => {
          const sender = message.nicknamefrom;
          senders.add(sender); 
        });

        const count = messages.length;
        const senderNames = Array.from(senders);

        unreadMessageCount(count, senderNames); 
      } else {
        console.log('No unread messages');
        const unreadMessagesElement = document.getElementById('unread-messages');
        unreadMessagesElement.textContent = 'No unread messages';
      }
    })
    .catch(error => {
      console.error('An error occurred while fetching unread messages:', error);
    });
}
