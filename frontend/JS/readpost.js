import { loadUserPage } from './userpage.js';
import { displayErrorMessage } from './errormessage.js';

export function loadPostPage(postID) {
    fetch(`/read-post?id=${postID}`)
    .then(response => response.json())
    .then(data => {
        var formContainer = document.getElementById('formContainer');
        var formattedDate = new Date(data.post.date).toLocaleString();

        var postHTML = `
                    <div class="readpost">
                        <div class="heading">
                            ${data.post.title}
                        </div>
                        <p class="content">
                            ${data.post.content}
                        </p>
                        <div class="poster">
                            Posted by ${data.post.nickname}
                            at ${formattedDate}
                        </div>
                        <div class="poster">
                            ${data.post.movies}
                            ${data.post.serials}
                            ${data.post.realityshows}
                        </div>
                    </div>
                    `;
        var commentsHTML = "";
        if (data.comments) {
            data.comments.forEach(comment => {
                var commentFormattedDate = new Date(comment.commentdate).toLocaleString();
                commentsHTML += `
                    <div class="readcomment">
                        <div class="content">
                            ${comment.comment}
                        </div>
                        <div class="poster">
                            Commented by ${comment.commentnickname} 
                            at ${commentFormattedDate}
                        </div>
                    </div>
                `;
            });
        }

        formContainer.innerHTML = postHTML + commentsHTML + `
            <form id="commentform">
                <p class="align">
                    <textarea id="comment" class="input" cols="40" rows="5" wrap="hard" name="comment" placeholder="Comment" required></textarea>
                </p>
                <div class="align">
                    <button class="buttons" type="submit">Add comment</button>
                </div>
            </form>
            <p class="align">
                <input id="back" class="buttons" type="button" value="Back to main page">
            </p>
        `;

        const commentForm = document.getElementById('commentform');
        commentForm.addEventListener('submit', event => submitComment(event, data));

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

        window.history.pushState({ page: 'readpost', }, '', `/`);
        window.addEventListener('popstate', function () {
        loadUserPage()
        });
    })
    .catch(error => {
        displayErrorMessage(`An error occurred while loading a post: ${error.message}`);
    });
}

function submitComment(event, data) {
    event.preventDefault();

    const comment = document.getElementById('comment').value;
    const postID = data.post.ID

    const commentData = {
        comment: comment,
        postID: postID,
    };

    sendCommentData(commentData)
}

function sendCommentData(commentData) {
    fetch(`/commenting?id=${commentData.postID}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(commentData)
    })
    .then(response => {
        if (response.ok) {
            loadPostPage(commentData.postID); 
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
        displayErrorMessage(`An error occurred while posting a comment: ${error.message}`);
    });
}