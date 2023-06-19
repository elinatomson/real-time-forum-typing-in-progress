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
        userItem.dataset.user = user.nickname;

        //CSS class to indicate the online/offline status
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
  fetch('/messages?nicknameTo=' + nicknameTo)
  .then(response => response.json())
  .then(messages => {  
      const messageBox = document.getElementById("message-box");

      if (messages && messages.length > 0) {
          messages.forEach(message => {
              messageBox.value += `${message.nicknamefrom}: ${message.message}\n`;
          });
      } else {
          messageBox.value = "";
      }
  })
  .catch(error => {
      var formContainer = document.getElementById('formContainer');
      var errorContainer = document.createElement('div');
      errorContainer.className = 'message';
      errorContainer.textContent = 'An error occurred while displaying messages: ' + error.message;
      formContainer.appendChild(errorContainer);
  });
}