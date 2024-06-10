# go-chat-app

This is a simple chat application using WebSocket Gorilla. Users can create and join chat rooms, send and receive messages in real-time.

## Table of Contents

- [User Endpoints](#user-endpoints)
  - [Signup](#signup)
  - [Login](#login)
- [WebSocket Endpoints](#websocket-endpoints)
  - [Create Room](#create-room)
  - [Join Room](#join-room)
  - [Get Rooms](#get-rooms)
  - [Get Clients](#get-clients)
- [Notes](#notes)

## User Endpoints

### Signup

- **Endpoint:** `/signup`
- **Method:** `POST`
- **Description:** Registers a new user.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

### Login

- **Endpoint:** `/login`
- **Method:** `POST`
- **Description:** login user.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

### Create Room

- **Endpoint:** `/ws/createRoom`
- **Method:** `POST`
- **Description:** Create a new room chat.
- **Request Body:**
  ```json
  {
    "id": "string",
    "name": "string"
  }
  ```
- **Request Headers:**
  ```json
  {
    "Authorization": "Bearer <Token>"
  }
  ```

### Get Rooms

- **Endpoint:** `/ws/getRooms`
- **Method:** `GET`
- **Description:** Get list room chat.
- **Request Headers:**
  ```json
  {
    "Authorization": "Bearer <Token>"
  }
  ```

### Join Room Chat

- **Endpoint:** `/ws/joinRoom/:roomId`
- **Method:** `GET`
- **Description:** Join the room chat.
- **Request Headers:**
  ```json
  {
    "Authorization": "Bearer <Token>"
  }
  ```

### Get Clients

- **Endpoint:** `/ws/getClients/:roomId`
- **Method:** `GET`
- **Description:** Get all clients inside the room chat.
- **Request Headers:**
  ```json
  {
    "Authorization": "Bearer <Token>"
  }
  ```

## Notes

To test the WebSocket endpoints, you can use Postman, which now supports WebSocket connections. Here are the steps to use WebSocket features in Postman:

1. **Open Postman**: Ensure you have the latest version of Postman installed.
2. **Navigate to WebSocket Request**: Click on the "New" button and select "WebSocket Request".
3. **Enter URL**: Input the WebSocket endpoint URL, e.g., `ws://yourapi.com/ws/joinRoom/:roomId`.
4. **Connect**: Click on the "Connect" button.
5. **Send Messages**: You can now send and receive messages to test the WebSocket functionality.

---

For further details or support, please refer to the project's documentation or contact the support team.