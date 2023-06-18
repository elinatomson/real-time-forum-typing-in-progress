export function logOut() {
    var logout = document.getElementById('logout-button');
    logout.addEventListener('click', function(event) {
        event.preventDefault();
        loggingOut();
    });
}

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
            return response.text(); //returns response.text() to propagate the error message to the next .then() in the chain
        }
    })
    .then(errorMessage => { //to handle the error message received from the previous .then() in case the response was not successful
        if (errorMessage) {
            var formContainer = document.getElementById('formContainer');
            var errorContainer = document.createElement('div'); //creating an error container element and appending error message to the form container.
            errorContainer.className = 'message';
            errorContainer.textContent = errorMessage;
            formContainer.appendChild(errorContainer);
        }
    })
    .catch(error => { //to handle any other errors that might occur during the fetch request or any of the .then() functions
        var formContainer = document.getElementById('formContainer');
        var errorContainer = document.createElement('div');
        errorContainer.className = 'message';
        errorContainer.textContent = error.message;
        formContainer.appendChild(errorContainer);
    });
}

function showLoggedOutMessage(html) {
   //logout HTML content
   var contentContainer = document.createElement('div');
   contentContainer.innerHTML = html;

   var modifiedHTML = `
   <header class="header">
     <div class="heading">
       <div id="mainpage">Fun Facts Forum</div>
     </div>
   </header>
   <div class="heading">You are logged out. Come visit us again!</div>
   <p class="align">
       <input id="main-page2" class="buttons" type="button" value="Back to main page">
   </p>
   <footer class="footer">
     <div>Authors:</div>
     <a class="authors" href="https://01.kood.tech/git/elinat">elinat</a> <br>
     <a class="authors" href="https://01.kood.tech/git/Anni.M">Anni.M</a>
   </footer>
 `;

    document.body.innerHTML = modifiedHTML; //replacing entire document body with the modified HTML structure
    document.body.appendChild(contentContainer); //adding user-specific content to the document body

    var mainPage2 = document.getElementById('main-page2');
    mainPage2.addEventListener('click', function(event) {
        event.preventDefault();
        window.location.href = '/';
    });
}

