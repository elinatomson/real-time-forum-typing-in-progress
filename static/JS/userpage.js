import { usersForChat } from './messages.js';
import { loadPostPage } from './readpost.js';
import { newPost } from './posting.js';
import { logOut } from './logout.js';

export function loadUserPage() {
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
          <div id="forum-users" class="forumusers">
            <div class="post-list-title">Forum users:</div>
            <div id="user-list-container"></div>
          </div>
          <div id="formContainer" class="form-container">
            <div id="postContainer"></div>
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

      usersForChat()
      newPost()
      logOut()

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
    

  
  