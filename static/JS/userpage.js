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
              <a id="newpost">New post</a>
              <a id="logout-button">Log out</a>
            </div>
          </div>
          <div class="heading">
            <div id="mainpage">Fun Facts Forum</div>
          </div>
        </header>
        <div class="container">
          <div class="forumusers">
            <div class="post-list-title">Forum users:</div>
            <div id="user-list-container"></div>
          </div>
          <div id="formContainer" class="form-container">
            <div id="postContainer"></div>
          </div>
        </div>
        <div>
          <textarea id="message-box" rows="10" cols="50" readonly></textarea>
          <div>
            <input type="text" id="message-input">
            <button id="send-button">Send</button>
          </div>
        </div>
        <footer class="footer">
          <div>Authors:</div>
          <a class="authors" href="https://01.kood.tech/git/elinat">elinat</a> <br>
          <a class="authors" href="https://01.kood.tech/git/Anni.M">Anni.M</a>
        </footer>      
      `;
  
      document.body.innerHTML = modifiedHTML; //replacing entire document body with the modified HTML structure
      document.body.appendChild(contentContainer); //adding user-specific content to the document body

      //to get online and offline users list
        fetch('/users')
        .then(response => response.json())
        .then(users => {
    
          // Update the user page HTML with the user list
          const userListContainer = document.getElementById('user-list-container');
          userListContainer.className = 'users-container';
    
          users.forEach(user => {
            const userItem = document.createElement('div');
            userItem.className = 'user';
            userItem.textContent = user.nickname;
    
            // Add a CSS class to indicate the online/offline status
            userItem.classList.add(user.online ? 'online' : 'offline');
    
            userListContainer.appendChild(userItem);
          });
        })
        .catch(error => {
          console.error('Error loading user list:', error);
      });

      //including the .js files, because loadUserPage function is used in those
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

      var websocketScript = document.createElement('script');
      websocketScript.src = './static/JS/websocket.js';
      document.body.appendChild(websocketScript);

    //to see all posts on user mainpage
    fetch('/posts')
        .then(response => response.json())
        .then(posts => {
          if (Array.isArray(posts)) { // Check if posts is an array
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
            
            //to make post clickable, so you can see this post
              postElement.onclick = function() {
                loadPostPage(post.ID);
              };
                postContainer.appendChild(postElement);
            });
          } else {
            var formContainer = document.getElementById('formContainer');
            var errorContainer = document.createElement('div');
            errorContainer.className = 'message';
            errorContainer.textContent = 'No posts available. Please click on Menu to create a post.';
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
    })
    .catch(error => { 
      var formContainer = document.getElementById('formContainer');
      var errorContainer = document.createElement('div');
      errorContainer.className = 'message';
      errorContainer.textContent = error.message;
      formContainer.appendChild(errorContainer);
    });
    }
    

  
  
  
  