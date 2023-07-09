# real-time-forum-typing-in-progress

## Project Description
This is a private forum that provides a platform for registered users to engage in discussions by creating forum posts, associating categories with their posts, and participating in reading and commenting on all the posts. Additionally, the forum offers a chat feature that enables users to send private messages to each other. his version of the forum uses a typing in progress engine, allowing users to see when the person they are chating with is typing. 

The chat functionality allows real-time communication between users who are online simultaneously. Even if a user is offline, they will still receive notifications about unread messages when they log in.

The backend functionality is implemented in Go handlers, which handle data processing and interactions with the database. WebSockets are utilized for real-time communication between clients and the server. The frontend of the forum is built using JavaScript, which handles all the frontend events and interactions. WebSockets are also used on the client-side for real-time chat functionality. HTML is used to organize the elements of the page, providing the structure and layout of the forum. CSS is implemented to style and customize the appearance of the elements, ensuring a visually user-friendly interface.

## How to use
- Type in your terminal: go run main.go
- In the Windows Security alert cklick "Allow access"
- Open http://localhost:8080
- To stop the server, click Ctrl + C in your terminal

OR with Docker:
- You should have Docker installed. 
- To build the image and run the container type in your terminal: sh build_docker.sh 
- Then open localhost:8080 in your browser to visit the forum.
- When you are finished and you want to delete docker container and images press Enter in your terminal. It will run delete_docker script. 

## For auditing
Login to the forum using two different users on different browsers(for example Firefox and Chrome).
The audit questions are [here](https://github.com/01-edu/public/blob/master/subjects/real-time-forum/typing-in-progress/audit.md)
You can register totally new forum users or you can use excisting ones:
- Nickname: Leonardo
- Password: Tere1
- Nickname: Elina
- Password: Tere1


## Authors
- [@elinat](https://01.kood.tech/git/elinat)
- [@Anni.M](https://01.kood.tech/git/Anni.M)

