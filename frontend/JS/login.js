import { loadUserPage } from './userpage.js';
import { displayErrorMessage } from './errormessage.js';

document.addEventListener("DOMContentLoaded", () => {
  var logIn = document.getElementById('login-button');
  logIn.addEventListener('click', function(event) {
    event.preventDefault();
    showLogInForm();
  });
});

export function showLogInForm() {
    var formContainer = document.getElementById('formContainer');
    formContainer.innerHTML = `
    <div class="heading">Please log in!</div>
    <form id="login-form">                
      <p class="align">
        <label for="nickname">Nickname</label>
        <input id="nickname" class="input" type="text" placeholder="" name="nickname">
      </p>
      <p class="align">
        <label for="password">Password</label>
        <input id="password" class="input" type="password" placeholder="" name="password">
      </p>
      <p class="align">
        <input class="buttons" type="submit" value="Login">
        <input id="back" class="buttons" type="button" value="Cancel">
      </p>
    </form>
    `;

    var logInForm = document.getElementById('login-form');
    logInForm.addEventListener('submit', function(event) {
      event.preventDefault();

      var nickname = document.getElementById('nickname').value;
      var password = document.getElementById('password').value;
      var userData = {
        nickname: nickname,
        password: password
      };

      submitLogInForm(userData);
    });

    const backButton = document.getElementById('back');
    backButton.addEventListener('click', function(event) {
      event.preventDefault();
      window.location.href = '/';
    });
  }

  function submitLogInForm(userData) {
    fetch('/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(userData)
    })
    .then(response => {
      if (response.ok) {
        loadUserPage(); //load the logged in user page
      } else {
          return response.text(); //returns response.text() to propagate the error message to the next .then() in the chain
      }
    })
    .then(errorMessage => {
      if (errorMessage) {
        displayErrorMessage(errorMessage);  //to handle the error message received from the previous .then() in case the response was not successful
      }
    })
    .catch(error => { //to handle any other errors that might occur during the fetch request or any of the .then() functions
      displayErrorMessage(`An error occurred while logging in: ${error.message}`);
    });
  }