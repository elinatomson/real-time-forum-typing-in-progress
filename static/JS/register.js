var registerButton = document.getElementById('register-button')
registerButton.addEventListener('click', function(event) {
  event.preventDefault();
  showRegistrationForm();
});

function showRegistrationForm() {
  var formContainer = document.getElementById('formContainer');
  formContainer.innerHTML = `
  <div class="heading">Please fill out all the fields below</div>
  <div class="content">Username and password have to be 5 letters long!</div>
  <form id="registration-form">
    <p class="align">
      <label for="nickname">Nickname:</label>
      <input class="input" type="text" id="nickname" name="nickname" required>
    </p>
    <p class="align">
      <label for="age">Age:</label>
      <input class="input" type="number" id="age" name="age" required>
    </p>
    <p class="align">
      <label for="gender">Gender:</label>
      <select class="input" id="gender" name="gender" required>
        <option class="input" value="">Select</option>
        <option class="input" value="male">Male</option>
        <option class="input" value="female">Female</option>
        <option class="input" value="other">Other</option>
      </select>
    </p>
    <p class="align">
      <label for="firstName">First Name:</label>
      <input class="input" type="text" id="firstName" name="firstName" required>
    </p>
    <p class="align">
      <label for="lastName">Last Name:</label>
      <input class="input" type="text" id="lastName" name="lastName" required>
    </p>
    <p class="align">
      <label for="email">E-mail:</label>
      <input class="input" type="email" id="email" name="email" required>
    </p>
    <p class="align">
      <label for="password">Password:</label>
      <input class="input" type="password" id="password" name="password" required>
    </p>
    <p class="align">
      <label for="password-repeat">Repeat Password:</label>
      <input class="input" type="password" id="password-repeat" name="password-repeat" required>
    </p>
    <p class="align">
      <input class="buttons" type="submit" value="Register">
    </p>
  </form>
`;

  var registrationForm = document.getElementById('registration-form');

  registrationForm.addEventListener('submit', function(event) {
    event.preventDefault();

    var nickname = document.getElementById('nickname').value;
    var age = document.getElementById('age').value;
    var gender = document.getElementById('gender').value;
    var firstName = document.getElementById('firstName').value;
    var lastName = document.getElementById('lastName').value;
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;
    var password_repeat = document.getElementById('password-repeat').value;

    submitRegistrationForm(nickname, age, gender, firstName, lastName, email, password, password_repeat);
  });
}

function submitRegistrationForm(nickname, age, gender, firstName, lastName, email, password, password_repeat) {
  //validations
  var nameLength = nickname.length >= 5 && nickname.length <= 50;
  var passwordLength = password.length >= 5 && password.length <= 50;
  var passwordMatch = password === password_repeat;

  if (!nameLength || !passwordLength) {
    var formContainer = document.getElementById('formContainer');
    var errorMessage = document.createElement('div');
    errorMessage.className = 'message';
    errorMessage.textContent = 'Please check username and password criteria!';
    formContainer.appendChild(errorMessage);
    return;
  }

  if (!passwordMatch) {
    var formContainer = document.getElementById('formContainer');
    var errorMessage = document.createElement('div');
    errorMessage.className = 'message';
    errorMessage.textContent = 'Inserted passwords are not the same!';
    formContainer.appendChild(errorMessage);
    return;
  }

  //validation passed, send user data to the server
  var userData = {
    nickname: nickname,
    age: age,
    gender: gender,
    firstName: firstName,
    lastName: lastName,
    email: email,
    password: password
  };
  sendUserData(userData);
}

function sendUserData(userData) {
  fetch('http://localhost:8080/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(userData)
  })
  .then(response => {
    if (response.ok) {
      showLogInForm(); //direct the user to log in
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
    errorContainer.textContent = error.message;
    formContainer.appendChild(errorContainer);
  });
}

