import { displayErrorMessage } from './errormessage.js';

export function logOut() {
    var logout = document.getElementById('logout-button');
    logout.addEventListener('click', function(event) {
        event.preventDefault();
        loggingOut();
    });
}

function loggingOut() {
    fetch('/logout', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (response.ok) {
            showLoggedOutMessage();
        } else {
            return response.text(); 
        }
    })
    .then(errorMessage => { 
        if (errorMessage) {
            displayErrorMessage(errorMessage);
        }
    })
    .catch(error => { 
        displayErrorMessage(`An error occurred while logging out: ${error.message}`);
    });
}

function showLoggedOutMessage() {
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

    document.body.innerHTML = modifiedHTML; 
    
    var mainPage2 = document.getElementById('main-page2');
    mainPage2.addEventListener('click', function(event) {
        event.preventDefault();
        window.location.href = '/';
        window.history.pushState({ page: 'mainpage' }, '', '/');
    });

    window.history.pushState({ page: 'mainpage', }, '', `/`);
    window.addEventListener('popstate', function () {
        window.location.href = '/';
    });
}

