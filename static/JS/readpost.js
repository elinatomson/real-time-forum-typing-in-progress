function loadPostPage(postID) {
    fetch(`/readpost?id=${postID}`)
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

        var commentForm = document.getElementById('commentform');
        commentForm.addEventListener('submit', function(event) {
            event.preventDefault();
    
            var comment = document.getElementById('comment').value;
            var postID = data.post.ID
    
            submitComment(comment, postID);
        });
        //if you are logged in, then clicking on the Forum name, you will see the userPage as a mainpage
        var mainPage = document.getElementById('mainpage');
        mainPage.addEventListener('click', function(event) {
            event.preventDefault();
            loadUserPage();
        });
    
        //by clicking on the "Back to main page" button, you will see the userPage as a mainpage
        var backButton = document.getElementById('back');
        backButton.addEventListener('click', function(event) {
            event.preventDefault();
            loadUserPage();
        });
    })
    .catch(error => {
        console.error('An error occurred while loading the post:', error);
    });
}

function submitComment(comment, postID) {
    var commentData = {
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
            loadUserPage() //sending user to the userPage which is in this case as a mainpage
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
        errorContainer.textContent = error.message;
        formContainer.appendChild(errorContainer);
    });
}