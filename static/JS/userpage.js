function loadUserPage() {
    fetch('/userpage')
      .then(response => response.text())
      .then(html => {
        //userPage HTML content
        var contentContainer = document.createElement('div');
        contentContainer.innerHTML = html;
  
        var modifiedHTML = `
        <header class="header">
          <div class="name">Welcome!</div>
          <div class="dropdown">
            <button class="dropbtn">Menu</button>
            <div class="dropdown-content">
              <a id="logout-button">Log out</a>
              <a id="newpost">New post</a>
            </div>
          </div>
          <div class="heading">
            <div id="mainpage">Fun Facts Forum</div>
          </div>
        </header>
        <div id="formContainer" class="form-container">
          <div id="postContainer"></div>
        </div>
      `;
  
      document.body.innerHTML = modifiedHTML; //replacing entire document body with the modified HTML structure
      document.body.appendChild(contentContainer); //adding user-specific content to the document body
  
      //including the .js files, because loadUserPage function is used over there
      var logInScript = document.createElement('script');
      logInScript.src = './static/JS/login.js';
      document.body.appendChild(logInScript);

      var logOutScript = document.createElement('script');
      logOutScript.src = './static/JS/logout.js';
      document.body.appendChild(logOutScript);
  
      var newPostScript = document.createElement('script');
      newPostScript.src = './static/JS/newpost.js';
      document.body.appendChild(newPostScript);

      var readPostScript = document.createElement('script');
      readPostScript.src = './static/JS/readpost.js';
      document.body.appendChild(readPostScript);


    //to see all posts on user mainpage
    fetch('/posts')
        .then(response => response.json())
        .then(posts => {
            var postContainer = document.getElementById('postContainer');
            posts.forEach(post => {
                var postElement = document.createElement('div');
                postElement.classList.add('post');
                var formattedDate = new Date(post.date).toLocaleString();
                postElement.innerHTML = `
                <div class="post-list-title">${post.title}</div>
                <div class="poster">
                <p>${post.content}</p>
                Posted by ${post.nickname}
                at ${formattedDate}
                </div>
                <div class="poster">
                ${post.movies}
                ${post.serials}
                ${post.realityshows}
                </div>
                `;
            
            //to make post clickable, so you can see specific post
              postElement.onclick = function() {
                loadPostPage(post.ID);
              };
                postContainer.appendChild(postElement);
            });
        })
        .catch(error => {
            console.error('An error occurred while fetching the posts:', error);
        });
    })
    .catch(error => {
      console.error('An error occurred while loading the user page:', error);
    });
  }
  
  