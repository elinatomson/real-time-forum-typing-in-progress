import { loadUserPage } from './userpage.js';
import { displayErrorMessage } from './errormessage.js';

export function newPost() {
    var newPostButton = document.getElementById('newpost');
    newPostButton.addEventListener('click', showNewPostForm);
}

function showNewPostForm() {
    var formContainer = document.getElementById('formContainer');

    formContainer.innerHTML = `
    <form id="newpostform">
        <p class="align">
            <input id="title" class="input" type="text" name="title" placeholder="Title" required>
        </p>
        <p class="align">
            <textarea id="content" class="input"  cols="43" rows="10" wrap="hard" name="content" placeholder="Content" required></textarea>
        </p>
        <div class="content">Please choose one or more categories</div>
        <p class="radio">
            <input class="radio_input" type="checkbox" name="movies" value="movies" id="movies">
            <label class="radio_label" for="movies">Movies</label>
            <input class="radio_input" type="checkbox" name="serials" value="serials" id="serials">
            <label class="radio_label" for="serials">Serials</label>
            <input class="radio_input" type="checkbox" name="realityshows" value="realityshows" id="realityshows">
            <label class="radio_label" for="realityshows">Reality Shows</label>
        </p>
        <p class="align">
            <button class="buttons" type="submit">Publish Post</button>
            <input id="back" class="buttons" type="button" value="Cancel">
        </p>
    </form>
    `;

    const newPostForm = document.getElementById('newpostform');
    newPostForm.addEventListener('submit', newPostFormSubmit);

    var mainPage = document.getElementById('mainpage');
    mainPage.addEventListener('click', function(event) {
        event.preventDefault();
        loadUserPage()
        window.history.pushState({ page: 'userpage' }, '', '/');
    });

    var backButton = document.getElementById('back'); 
    backButton.addEventListener('click', function(event) {
        event.preventDefault();
        loadUserPage()
        window.history.pushState({ page: 'userpage' }, '', '/');
    });

    window.history.pushState({ page: 'readpost' }, '', `/`);
    window.addEventListener('popstate', function () {
    loadUserPage()
    });
}

function newPostFormSubmit(event) {
    event.preventDefault();

    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;
    
    const postData = {
        title: title,
        content: content,
      };

    //checking if a category is selected and add it to the postData object
    if (document.getElementById('movies').checked) {
       postData.movies = document.getElementById('movies').value;
    }
    if (document.getElementById('serials').checked) {
      postData.serials = document.getElementById('serials').value;
    }
    if (document.getElementById('realityshows').checked) {
      postData.realityshows = document.getElementById('realityshows').value;
    }

    sendPostData(postData);
}


function sendPostData (postData) {
    fetch('/create-post', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(postData)
    })
    .then(response => {
        if (response.ok) {
            loadUserPage() 
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
        displayErrorMessage(`An error occurred while posting: ${error.message}`);
    });
}
