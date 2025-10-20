# Go Web API

This is a simple Go web application that provides a RESTful API for managing users.

## Project Structure

The project follows a clean architecture pattern, separating concerns into the following layers:

-   `cmd/web`: Main application entry point.
-   `delivery/http`: HTTP handlers and routing.
-   `entity`: Core business entities.
-   `repository`: Data storage and retrieval.
-   `usecase`: Business logic and use cases.

## Getting Started

To build and run the project, follow these steps:

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/go-web.git
    ```

2.  **Navigate to the project directory:**

    ```bash
    cd go-web
    ```

3.  **Run the application:**

    ```bash
    go run cmd/web/main.go
    ```

The server will start on port 8087.

## API Endpoints

### Get User by ID

-   **GET** `/users/{id}`

Retrieves a user by their ID.

**Example request:**

```bash
curl http://localhost:8087/users/1
```
