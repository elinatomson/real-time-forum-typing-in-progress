document.addEventListener('DOMContentLoaded', function() {
    var logOut = document.getElementById('logout-button');
  
    logOut.addEventListener('click', function(event) {
        event.preventDefault();
        loggingOut();
    });
  
    function loggingOut() {
        fetch('http://localhost:8080/logout', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(response => {
            if (response.ok) {
                showLoggedOutMessage();
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
            errorContainer.textContent = 'An error occurred while logging out: ' + error.message;
            formContainer.appendChild(errorContainer);
        });
    }
  
    function showLoggedOutMessage() {
        var formContainer = document.getElementById('formContainer');
        formContainer.innerHTML = '';

        var formContent = `
            <div class="heading">You are logged out. Come visit us again!</div>
            <p class="align">
                <input id="main-page" class="buttons" type="button" value="Back to main page">
            </p>
        `;

        formContainer.innerHTML = formContent;

        var mainPage = document.getElementById('main-page');
  
        mainPage.addEventListener('click', function(event) {
            event.preventDefault();

            window.location.href = '/';
        });
    }
});








  