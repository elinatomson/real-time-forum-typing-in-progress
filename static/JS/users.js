import { handleUserClick } from './messages.js';
import { webSoc } from './websocket.js';
import { displayErrorMessage } from './errormessage.js';

export function usersForChat() {
    fetch('/users')
        .then(response => response.json())
        .then(users => {
            const currentUser = getCurrentUser(users);
            sortUsers(users);
            renderUserList(users, currentUser);
            attachUserClickListeners(users, currentUser);
        })
        .catch(error => {
            displayErrorMessage(`An error occurred while displaying users: ${error.message}`);
        });
}

//iterates over the users array and checks for the currentUser property in each user object 
//if it finds a user with currentUser set to true, it returns the value of the nickname property
function getCurrentUser(users) {
    for (const user of users) {
        if (user.currentUser) {
            return user.nickname;
        }
    }
    return null;
}

//sorting the users array based on the last_message_date property in descending order 
//if the dates are equal, it sorts them alphabetically 
function sortUsers(users) {
    users.sort((a, b) => {
        const aDate = new Date(a.last_message_date);
        const bDate = new Date(b.last_message_date);

        if (aDate > bDate) return -1;
        if (aDate < bDate) return 1;
        return a.nickname.localeCompare(b.nickname);
    });
}

//populating the user list in the HTML document
function renderUserList(users, currentUser) {
	const userListContainer = document.getElementById('user-list-container');
	userListContainer.className = 'users-container';
  
	users.forEach(user => {
		//exclude the logged in user from the list
		if (user.nickname !== currentUser) {
			const userItem = document.createElement('div');
			userItem.className = 'user';
			userItem.textContent = user.nickname;
			userItem.dataset.user = user.nickname;
			userItem.classList.add(user.online ? 'online' : 'offline');
			userListContainer.appendChild(userItem);
		}
	});
  }

//attaching click event listeners to the user elements in the user list
function attachUserClickListeners(users, currentUser) {
    const userElements = document.querySelectorAll('.user');

    userElements.forEach(userElement => {
        userElement.addEventListener('click', () => {
            const nicknameTo = userElement.dataset.user;
            const nicknameFrom = currentUser;
            handleUserClick(nicknameTo);
            webSoc(nicknameTo, nicknameFrom);
        });
    });
}
