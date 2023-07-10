import { usersForChat } from './users.js';
import { unreadMessages } from './messages.js';
import { loadPostPage } from './readpost.js';
import { newPost } from './posting.js';
import { logOut } from './logout.js';
import { displayErrorMessage } from './errormessage.js';

export function loadUserPage() {
  fetch('/userpage')
    .then(response => {
      if (response.ok) {
        return response.json(); 
      }else{
        return response.text(); 
      }
    })
    .then(user => {
      var contentContainer = document.createElement('div');
      var modifiedHTML = `
      <header class="header" id="header">
        <div class="name">Welcome ${user.nickname}!</div>
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
      <div id="unread-messages" class="unread"></div>
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

    document.body.innerHTML = modifiedHTML; 
    document.body.appendChild(contentContainer); 

    usersForChat()
    newPost()
    logOut()
    unreadMessages()

    var mainPage = document.getElementById('mainpage');
    mainPage.addEventListener('click', function(event) {
        event.preventDefault();
        loadUserPage()
        window.history.pushState({ page: 'userpage' }, '', '/');
    });

  fetch('/posts')
      .then(response => response.json())
      .then(posts => {
        if (Array.isArray(posts)) { 
          var postContainer = document.getElementById('postContainer');
          postContainer.innerHTML = '';
          posts.forEach(post => {
              var postElement = document.createElement('div');
              postElement.classList.add('post');
              var formattedDate = new Date(post.date).toLocaleString();
              postElement.innerHTML = `
              <div class="post-list-title">${post.title}</div>
              <div class="poster">
                Posted by ${post.nickname}
                at ${formattedDate}
              </div>
              <div class="poster">
                ${post.movies}
                ${post.serials}
                ${post.realityshows}
              </div>
              `;
          
          postElement.addEventListener('click', function() {
            loadPostPage(post.ID);
          });

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
        displayErrorMessage(`An error occurred while displaying posts: ${error.message}`);
    });
  })
  .catch(error => { 
    displayErrorMessage(`An error occurred while displaying the page: ${error.message}`);
  });
}
  
  