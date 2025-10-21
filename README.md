# Go Web API

This is a simple Go web application that provides a RESTful API for managing users. The project is built with a clean architecture approach to ensure separation of concerns, maintainability, and testability.

## Project Structure

The project follows a clean architecture pattern, separating concerns into the following layers:

-   `cmd/web`: Main application entry point. This is where the application is initialized and started.
-   `delivery/http`: This layer is responsible for handling HTTP requests and responses. It includes:
    -   `routers`: Defines the API routes.
    -   `handlers`: Contains the HTTP handlers that receive requests, call the appropriate use cases, and send responses.
    -   `dto`: Data Transfer Objects for request and response bodies.
    -   `middleware`: For cross-cutting concerns like logging, authentication, etc.
-   `usecase`: This layer contains the core business logic of the application. It orchestrates the flow of data between the delivery layer and the repository layer.
-   `repository`: This is the data access layer. It provides an abstraction over the data storage, so the rest of the application is not concerned with the specific implementation of the database. In this project, we are using an in-memory repository.
-   `entity`: This layer contains the core business entities of the application, such as the `User` model.

## Layer Architecture

The layers are organized to enforce the dependency rule, where inner layers are independent of outer layers. The data flow is as follows:

```
+--------------------+      +--------------------+      +------------------+      +----------------+
|   Delivery (HTTP)  | ---> |      Usecase       | ---> |    Repository    | ---> |     Entity     |
+--------------------+      +--------------------+      +------------------+      +----------------+
        (Request)                (Business Logic)          (Data Access)           (Domain Model)
```

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