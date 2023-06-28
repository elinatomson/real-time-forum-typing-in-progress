import { handleUserClick } from './messages.js';
import { webSoc } from './websocket.js';
import { displayErrorMessage } from './errormessage.js';

export function usersForChat() {
	fetch('/users')
		.then(response => response.json())
		.then(users => {
			//sort users by last message date or nickname in alphabetical order
			users.sort((a, b) => {
				const aDate = new Date(a.last_message_date);
				const bDate = new Date(b.last_message_date);

				if (aDate > bDate) return -1;
				if (aDate < bDate) return 1;
				return a.nickname.localeCompare(b.nickname);
			});

			const userListContainer = document.getElementById('user-list-container');
			userListContainer.className = 'users-container';

			users.forEach(user => {
				const userItem = document.createElement('div');
				userItem.className = 'user';
				userItem.textContent = user.nickname;
				userItem.dataset.user = user.nickname;
				//CSS class to indicate the online/offline status
				userItem.classList.add(user.online ? 'online' : 'offline');
				userListContainer.appendChild(userItem);
			});

			attachUserClickListeners();
		})
		.catch(error => {
			displayErrorMessage(`An error occurred while displaying users: ${error.message}`);
	});
}

function attachUserClickListeners() {
    const users = document.querySelectorAll('.user');
    users.forEach(user => {
      user.addEventListener('click', () => {
        handleUserClick(user.dataset.user);
        webSoc(user.dataset.user)
      	});
    });
}