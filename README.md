# todo-list-heli-tech
This repository contains the solution for a test task provided by HeliTechnology as part of the application process.




# Todo List Application

## Overview

This project is a **Todo List Application** designed to manage tasks efficiently. It provides a RESTful API for creating, managing, and organizing tasks. The application is built with Go and follows clean architecture principles for maintainability and scalability.

## Features

- **Task Management**: Create, update, delete, and retrieve tasks.
- **File Upload**: Attach files to tasks and store them in S3-compatible storage.
- **Database Integration**: Uses PostgreSQL for task storage.
- **Message Queues**: Integrates with SQS (or LocalStack) for asynchronous task notifications.
- **Dockerized Environment**: Easy deployment using Docker Compose.
- **Automated Migrations**: Handles database schema changes automatically.
- **Testing and Benchmarking**: Includes unit tests and benchmarks.

## Prerequisites

Ensure the following are installed on your system:

- **Docker**
- **Docker Compose**
- **Make**
- **Go (1.23 or newer)**

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Asef264/todo-list-heli-tech
cd todo-list-heli-tech
```

### 2. set up the system with project requirements
```bash
make setup
```

### 3. Run the Application

Use the Makefile to build and run the application:

```bash
make run
```

This will:

- Start the PostgreSQL database.
- Start LocalStack to mock AWS services (S3, SQS).
- Apply database migrations automatically.

### 4. API Endpoints

#### A. **Task Management**

- **`POST /todo_item`**: Create a new task.

#### B. **File Upload**

- **`POST /files`**: Upload a file.
- **`POST /files/:file_name`**: Download a file .

### 5. Configuration

Environment variables can be configured in a `.env` file. Example:

```env
DATABASE_URL=postgres://user:password@localhost:5432/todo
AWS_ACCESS_KEY_ID=minioadmin
AWS_SECRET_ACCESS_KEY=minioadmin
AWS_REGION=ir-west
S3_BUCKET_NAME=todo-files
SQS_QUEUE_URL=http://localhost:4566/queue/todo-queue
```

### 6. Running Tests

Run unit tests with:

```bash
make test
```

### 7. Running Benchmarks

Run performance benchmarks with:

```bash
make benchmark
```

## Project Structure

```plaintext
.
├── cmd/                # Entry points for the application
├── internal/           # Core application logic and domain
├── api/                # API handlers and routing
├── config/             # Configuration files
├── migrations/         # Database migration scripts
├── test/               # Unit tests and integration tests
├── pkg/                # Helper libraries and utilities
├── Makefile            # Automation commands
├── go.mod, go.sum      # Go module files
└── README.md           # Project documentation
```

## Tools and Technologies

- **Go**: The main programming language.
- **PostgreSQL**: Relational database for storing tasks.
- **S3-Compatible Storage**: For file uploads.
- **SQS-Compatible Queue**: For task notifications.
- **Docker Compose**: For containerized development and deployment.
- **gomock/mockery**: For mocking external services during tests.

## License

This project is intended for evaluation purposes and is not meant for production use.
