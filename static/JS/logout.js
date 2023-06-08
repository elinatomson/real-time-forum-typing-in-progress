document.addEventListener('DOMContentLoaded', function() {
    var logOut = document.getElementById('logout-button');
  
    logOut.addEventListener('click', function(event) {
        event.preventDefault();
        loggingOut();
    });
  
    function loggingOut() {
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

  