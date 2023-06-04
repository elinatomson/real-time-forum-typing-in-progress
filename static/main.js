document.addEventListener('DOMContentLoaded', function() {
    var mainPage = document.getElementById('mainpage');
  
    mainPage.addEventListener('click', function(event) {
        event.preventDefault();
        window.location.href = '/static/index.html';//THIS HAS TO BE CHANGED LATER
    });
})