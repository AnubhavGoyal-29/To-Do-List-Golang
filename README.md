# To-Do-List-Golang
A new to do list application in go

# Setup Guide

## **Introduction**
This guide provides step-by-step instructions to set up and run the To-Do-List-Golang project.

---

## **1. Install Golang(requires 1.18+) (Skip if already installed)**
First, check if Golang is already installed:
```sh
go version
```
If it outputs a version (e.g., `go version go1.18 linux/amd64`), you can **skip this step**.

#### **If Golang is not installed, install it using:**
```sh
sudo apt update
sudo apt install -y golang
```

or manually install a specific version (e.g., **Go 1.20.4**):
```sh
wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc
```
Then, verify the installation:
```sh
go version
```

---

## **2. Clone the Repository**
```sh
git clone https://github.com/AnubhavGoyal-29/To-Do-List-Golang.git
cd To-Do-List-Golang
```

---

## **3. Install Dependencies**
```sh
go mod tidy
```

---

## **4. Set Up SQLite Database**
No external setup is required! The project uses SQLite, and the database file (`task_management.db`) will be automatically created.

To check the database:
```sh
sqlite3 task_management.db
```
To list tables:
```sh
.tables
```

---

## **5. Run the Server**
```sh
go run main.go
```

The server should start successfully, and you can now use the API endpoints.

---

## **6. Environment Variables (Optional)**
The project may require environment variables. If needed, create a `.env` file in the project root:
```sh
touch .env
```
Example `.env` file:
```sh
JWT_SECRET=your_secret_key
PORT=8080
```

---


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
    "userToken" : YOUR_JWT_TOKEN
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
