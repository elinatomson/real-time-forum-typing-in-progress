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

            //validations will come here

            submitLogInForm(nickname, password);
        });
    }
  
    function submitLogInForm(nickname, password) {
        //necessary actions will come here
  
            window.location.href = '/';
        }
    });
  