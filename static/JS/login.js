var logIn = document.getElementById('login-button');
logIn.addEventListener('click', function(event) {
  event.preventDefault();
  showLogInForm();
});

function showLogInForm() {
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
}

function submitLogInForm(userData) {
  fetch('http://localhost:8080/login', {
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
        return response.text(); //reading response as text
    }
  })
  .then(errorMessage => {
    if (errorMessage) {
      var formContainer = document.getElementById('formContainer');
      var errorContainer = document.createElement('div');
      errorContainer.className = 'message';
      errorContainer.textContent = errorMessage;
      formContainer.appendChild(errorContainer);
    }
  })
  .catch(error => {
    var formContainer = document.getElementById('formContainer');
    var errorContainer = document.createElement('div');
    errorContainer.className = 'message';
    errorContainer.textContent = 'An error occurred while logging in: ' + error.message;
    formContainer.appendChild(errorContainer);
  });
}


function loadUserPage() {
  fetch('/userpage')
    .then(response => response.text())
    .then(html => {
      //userPage HTML content
      var contentContainer = document.createElement('div');
      contentContainer.innerHTML = html;

      var modifiedHTML = `
      <header class="header">
        <div class="heading">Welcome!</div>
        <div class="dropdown">
          <button class="dropbtn">Menu</button>
          <div class="dropdown-content">
            <a id="logout-button">Log out</a>
            <a id="newpost">New post</a>
          </div>
        </div>
        <div class="heading">
          <div id="mainpage">Fun Facts Forum</div>
        </div>
      </header>
      <div id="formContainer" class="form-container">
        <div class="heading2">POSTS!</div>
      </div>
    `;

    document.body.innerHTML = modifiedHTML; //replacing entire document body with the modified HTML structure
    document.body.appendChild(contentContainer); //adding user-specific content to the document body

    //including the .js files, because loadUserPage function is used there
    var script = document.createElement('script');
    script.src = './static/JS/logout.js';
    document.body.appendChild(script);

    var newPostScript = document.createElement('script');
    newPostScript.src = './static/JS/newpost.js';
    document.body.appendChild(newPostScript);

    var registerScript = document.createElement('script');
    registerScript.src = './static/JS/login.js';
    document.body.appendChild(registerScript);
  })
  .catch(error => {
    console.error('An error occurred while loading the user page:', error);
  });
}

