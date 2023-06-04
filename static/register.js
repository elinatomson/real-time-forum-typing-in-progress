document.addEventListener('DOMContentLoaded', function() {
  var registerButton = document.getElementById('register-button');

  registerButton.addEventListener('click', function(event) {
    event.preventDefault();
    showRegistrationForm();
  });

  function showRegistrationForm() {
    var formContainer = document.getElementById('formContainer');
    formContainer.innerHTML = '';

    var formContent = `
      <div class="heading">Please fill out all the fields below</div>
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
          <option class="input" value="">-- Select --</option>
          <option class="input" value="male">Male</option>
          <option class="input" value="female">Female</option>
          <option class="input" value="other">Other</option>
          </select>
        </p>
        <p class="align">
          <label for="firstName">First Name:</label>
          <input class="input" type="text" id="firstName" name="firstName" required><br><br>
        </p>
        <p class="align">
          <label for="lastName">Last Name:</label>
          <input class="input" type="text" id="lastName" name="lastName" required><br><br>
        </p>
        <p class="align">
          <label for="email">E-mail:</label>
          <input class="input" type="email" id="email" name="email" required><br><br>
        </p>
        <p class="align">
          <label for="password">Password:</label>
          <input class="input" type="password" id="password" name="password" required><br><br>
        </p>
        <p class="align">
          <input class="buttons" type="submit" value="Register">
        </p>
      </form>
    `;

    formContainer.innerHTML = formContent;

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

      //validations will be here

      submitRegistrationForm(nickname, age, gender, firstName, lastName, email, password);
    });
  }

  function submitRegistrationForm(nickname, age, gender, firstName, lastName, email, password) {
    //necessary actions will be here

    window.location.href = '/static/index.html'; //THIS HAS TO BE CHANGED LATER
  }
});
