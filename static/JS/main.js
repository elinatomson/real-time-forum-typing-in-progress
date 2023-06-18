import { loadUserPage } from './userpage.js';

//forum name clickable and directing to mainpage
loadPage()
//when there is a matching cookie in the database, then ise userpage is displayed as a mainpage
compareCookieWithDatabase()

function loadPage() {
    var mainPage = document.getElementById('mainpage');
    mainPage.addEventListener('click', function(event) {
        event.preventDefault();
        window.location.href = '/';
    });
}

//https://www.w3schools.com/js/js_cookies.asp
function getCookie(cname) { 
    let name = cname + "="; 
    let decodedCookie = decodeURIComponent(document.cookie); 
    let ca = decodedCookie.split(';'); 
    for(let i = 0; i <ca.length; i++) { 
    let c = ca[i]; 
        while (c.charAt(0) == ' ') { 
            c = c.substring(1); 
        } 
        if (c.indexOf(name) == 0) { 
            return c.substring(name.length, c.length); 
        } 
    } 
    return ""; 
} 


function compareCookieWithDatabase() {
    var clientCookie = getCookie("sessionId");

    fetch("/session", {
    method: "POST",
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify({ cookie: clientCookie })
    })
    .then((response) => response.text())
    .then((result) => {
        if (result === "Cookie matches!") {
            loadUserPage(); //if the cookie matches with the one which is in the database, then clicking on the forum name the userpage is loaded
        } 
    })
    .catch((error) => {
        console.error("An error occurred:", error);
    });
}