//forum name clickable and directing to mainpage
document.addEventListener('DOMContentLoaded', function() {
    var mainPage = document.getElementById('mainpage');
  
    mainPage.addEventListener('click', function(event) {
        event.preventDefault();
        window.location.href = '/';
    });
})