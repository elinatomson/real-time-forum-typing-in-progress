import { webSoc } from './websocket.js';
import { loadUserPage } from './userpage.js';
 
export function handleUserClick(user) {
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
          //sort chronologically and reverse the order of messages so the newest ones are at the bottom
          const messagesToDisplay = newMessages.sort((a, b) => a.date - b.date).reverse().join("\n") + "\n"; 
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

function messagesAsRead(nicknameFrom) {
  fetch(`/messages/markAsRead?nicknameFrom=${nicknameFrom}`)
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
  
  //to get sender names and make them clickable
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

    //if there is more than 1 sender, then separate them with commas
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
          //adding sender name to the set
          senders.add(sender); 
        });

        const count = messages.length;
        //converting the set to an array of sender names
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
