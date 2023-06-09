document.addEventListener('DOMContentLoaded', function() {
    var newPost1 = document.getElementById('newpost1');
    var newPost2 = document.getElementById('newpost2');
  
    newPost1.addEventListener('click', function(event) {
        event.preventDefault();
        showNewPostForm();
    });
      
    newPost2.addEventListener('click', function(event) {
        event.preventDefault();
        showNewPostForm();
    });
  
    function showNewPostForm() {
        var formContainer = document.getElementById('formContainer');
        formContainer.innerHTML = '';

        var formContent = `
        <form id="newpostform">
            <p class="align">
                <input id="title" class="input" type="text" name="title" placeholder="Title" required>
            </p>
            <p class="align">
                <textarea id="content" class="input"  cols="50" rows="10" wrap="hard" name="content" placeholder="Content" required></textarea>
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
            </p>
        </form>
        `;

        formContainer.innerHTML = formContent;

        var newPostForm = document.getElementById('newpostform');

        newPostForm.addEventListener('submit', function(event) {
            event.preventDefault();

            var title = document.getElementById('title').value;
            var content = document.getElementById('content').value;
            var movies = document.getElementById('movies').value;
            var serials = document.getElementById('serials').value;
            var realityshows = document.getElementById('realityshows').value;

            submitNewPostForm(title, content, movies, serials, realityshows);
        });
    }
  
    function submitNewPostForm(title, content, movies, serials, realityshows) {
        var postData = {
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

        sendPostData(postData)
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
                window.location.href = '/';
            } else {
                return response.text(); //reading response as text
            }
        })
        .then(data => {
            var formContainer = document.getElementById('formContainer');
            if (data) {
            var errorMessage = document.createElement('div');
            errorMessage.className = 'message';
            errorMessage.textContent = 'An error occurred while creating post: ' + data;
            formContainer.appendChild(errorMessage);
            }
        })
        .catch(error => {
            var formContainer = document.getElementById('formContainer');
            var errorMessage = document.createElement('div');
            errorMessage.className = 'message';
            errorMessage.textContent = 'An error occurred while creating post: ' + error.message;
            formContainer.appendChild(errorMessage);
        });
    }
});
