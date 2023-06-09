document.addEventListener('DOMContentLoaded', function() {
    var logIn = document.getElementById('login-button');
  
    logIn.addEventListener('click', function(event) {
      event.preventDefault();
      showLogInForm();
    });
  
    function showLogInForm() {
      var formContainer = document.getElementById('formContainer');
      formContainer.innerHTML = '';
  
      var formContent = `
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
  
        formContainer.innerHTML = formContent;
  
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
            window.location.href = '/';
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
});
  