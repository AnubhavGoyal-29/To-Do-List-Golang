# To-Do-List-Golang
A new to do list application in go

# API Documentation for To-Do List Microservices

## **User APIs**

### 1. **Create User**
- **Endpoint**: `POST /users/register`
- **Description**: Creates a new user account.
- **Request Body (JSON):**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword"
  }
  ```
- **cURL Request:**
  ```sh
  curl -X POST "http://localhost:8080/users/register" \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "password": "securepassword"}'
  ```
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```

### 2. **User Login**
- **Endpoint**: `POST /users/login`
- **Description**: Authenticates user and returns a JWT token.
- **Request Body (JSON):**
  ```json
  {
    "email": "john@example.com",
    "password": "securepassword"
  }
  ```
- **cURL Request:**
  ```sh
  curl -X POST "http://localhost:8080/users/login" \
  -H "Content-Type: application/json" \
  -d '{"email": "john@example.com", "password": "securepassword"}'
  ```
- **Response:**
  ```json
  {
    "token": "your_jwt_token_here"
  }
  ```

### 3. **Get User Profile**
- **Endpoint**: `GET /users/me`
- **Description**: Retrieves the authenticated user's details.
- **Headers:**
  ```
  Authorization: Bearer your_jwt_token_here
  ```
- **cURL Request:**
  ```sh
  curl -X GET "http://localhost:8080/users/me" \
  -H "Authorization: Bearer your_jwt_token_here"
  ```
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```

### 4. **Logout User**
- **Endpoint**: `POST /users/logout`
- **Description**: Logs out the user by invalidating the token.
- **Headers:**
  ```
  Authorization: Bearer your_jwt_token_here
  ```
- **cURL Request:**
  ```sh
  curl -X POST "http://localhost:8080/users/logout" \
  -H "Authorization: Bearer your_jwt_token_here"
  ```
- **Response:**
  ```json
  {
    "message": "Logout successful"
  }
  ```

---

## **Task APIs**

### 1. **Create Task**
- **Endpoint**: `POST /tasks/create`
- **Description**: Creates a new task for the authenticated user.
- **Headers:**
  ```
  Authorization: Bearer your_jwt_token_here
  ```
- **Request Body (JSON):**
  ```json
  {
    "title": "Complete API Documentation",
    "status": 1,
    "priority": 2,
    "type": 1
  }
  ```
- **cURL Request:**
  ```sh
  curl -X POST "http://localhost:8080/tasks/create" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_jwt_token_here" \
  -d '{"title": "Complete API Documentation", "status": 1, "priority": 2, "type": 1}'
  ```
- **Response:**
  ```json
  {
    "id": 101,
    "title": "Complete API Documentation",
    "status": "Pending",
    "priority": "Medium",
    "type": "Feature"
  }
  ```

### 2. **Get All Tasks**
- **Endpoint**: `GET /tasks`
- **Description**: Retrieves all tasks for the authenticated user.
- **Headers:**
  ```
  Authorization: Bearer your_jwt_token_here
  ```
- **cURL Request:**
  ```sh
  curl -X GET "http://localhost:8080/tasks" \
  -H "Authorization: Bearer your_jwt_token_here"
  ```
- **Response:**
  ```json
  [
    {
      "id": 101,
      "title": "Complete API Documentation",
      "status": "Pending",
      "priority": "Medium",
      "type": "Feature"
    }
  ]
  ```

### 3. **Update Task**
- **Endpoint**: `PUT /tasks/update/:id`
- **Description**: Updates an existing task.
- **Headers:**
  ```
  Authorization: Bearer your_jwt_token_here
  ```
- **Request Body (JSON):**
  ```json
  {
    "status": 2,
    "priority": 1
  }
  ```
- **cURL Request:**
  ```sh
  curl -X PUT "http://localhost:8080/tasks/update/101" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_jwt_token_here" \
  -d '{"status": 2, "priority": 1}'
  ```
- **Response:**
  ```json
  {
    "message": "Task updated successfully"
  }
  ```

### 4. **Delete Task**
- **Endpoint**: `DELETE /tasks/delete/:id`
- **Description**: Deletes a task.
- **Headers:**
  ```
  Authorization: Bearer your_jwt_token_here
  ```
- **cURL Request:**
  ```sh
  curl -X DELETE "http://localhost:8080/tasks/delete/101" \
  -H "Authorization: Bearer your_jwt_token_here"
  ```
- **Response:**
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```

---

## **Notes**
1. **Authorization**:
   - All **user-related** endpoints (except **create** and **login**) require a valid **JWT token** in the request header.
   - All **task-related** endpoints require a valid **JWT token** in the request header.
2. **Status Codes**:
   - `1 = Pending`, `2 = Completed`, `3 = Working`
3. **Priority Codes**
   - `1 = Low`, `2 = Medium`, `3 = High`
4. **Task Type**
   - `1 = Bug`, `2 = Feature`, `3 = Improvement`
