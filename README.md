# todo-list-heli-tech
This repository contains the solution for a test task provided by HeliTechnology as part of the application process.


# Todo Service

## Description
This project implements a **Todo Service** that allows users to manage `TodoItem` entities with the following key features:

1. **File Upload**: Upload files to an S3 bucket and store their references.
2. **TodoItem Management**: Create, store, and manage `TodoItem` entities with PostgreSQL.
3. **SQS Queue Integration**: Send `TodoItem` data to an SQS queue.
4. **Hexagonal Architecture**: Clean separation of concerns with ports and adapters.
5. **Dockerized Environment**: Easily deployable using Docker Compose.
6. **Unit Testing & Benchmarking**: Fully tested with mocks for external services.

## Prerequisites

Ensure you have the following installed:

- **Docker**
- **Docker Compose**
- **Make**
- **Go (1.23 or newer)**

## Setup Instructions

### 1. Clone the Repository

```bash
git clone <repository-url>
cd <repository-folder>
```

### 2. Build and Run the Project

Use the Makefile to start the project:

```bash
make run
```

This command will:

- Start the PostgreSQL database.
- Start LocalStack to mock AWS S3 and SQS.
- Apply database migrations automatically.

### 3. Endpoints

#### A. **File Upload**

- **Endpoint**: `POST /upload`
- **Description**: Allows users to upload a file (e.g., text or image) to S3.
- **Response**: Returns a `fileId` representing the file in S3.

#### B. **Create TodoItem**

- **Endpoint**: `POST /todo`
- **Description**: Create a `TodoItem` with a description, due date, and optional `fileId`.
- **Response**: Returns the created `TodoItem`.

### 4. Environment Variables

Configure the following environment variables in a `.env` file:

```env
DATABASE_URL=postgres://user:password@localhost:5432/todo
AWS_ACCESS_KEY_ID=minioadmin
AWS_SECRET_ACCESS_KEY=minioadmin
AWS_REGION=us-east-1
S3_BUCKET_NAME=todo-files
SQS_QUEUE_URL=http://localhost:4566/queue/todo-queue
```

### 5. Running Tests

Run unit tests using the Makefile:

```bash
make test
```

This will test:

- File uploads to S3.
- Creation of `TodoItem` in PostgreSQL.
- Sending messages to the SQS queue.

### 6. Running Benchmarks

Run benchmarks for critical operations:

```bash
make benchmark
```

Benchmarked operations include:

- Inserting a `TodoItem` into PostgreSQL.
- Uploading files to S3.
- Sending messages to SQS.

## Architecture

This project follows the **Hexagonal Architecture**:

- **Domain Logic**: Core business logic independent of infrastructure.
- **Ports**: Interfaces for interacting with the core logic (e.g., repository, SQS, S3).
- **Adapters**: Implementations of the ports for external systems (e.g., PostgreSQL, S3, SQS).

## Tools and Libraries

- **PostgreSQL**: For data persistence.
- **AWS S3**: For file storage.
- **AWS SQS**: For message queuing.
- **LocalStack**: To mock AWS services during development.
- **gomock/mockery**: For generating mocks in unit tests.

## Folder Structure

```plaintext
.
├── cmd/                # Entry points for the application
├── internal/           # Core application logic and domain
│   ├── domain/         # Domain entities and interfaces
│   ├── adapters/       # Infrastructure implementations (PostgreSQL, S3, SQS)
│   ├── ports/          # Interfaces for external dependencies
├── migrations/         # Database migration files
├── tests/              # Unit tests and benchmarks
├── docker-compose.yml  # Docker Compose configuration
├── Makefile            # Automation commands
├── .env                # Environment variables
└── README.md           # Project documentation
```

## Expected Deliverables

- **Dockerized Setup**: Includes PostgreSQL and LocalStack for mocking S3 and SQS.
- **Database Migrations**: Creates the `TodoItem` table with `id`, `description`, `dueDate`, and `fileId` columns.
- **Unit Tests**: Verifies file uploads, TodoItem creation, and SQS messaging.
- **Benchmarks**: Measures performance of key operations.

## License

This project is intended for evaluation purposes only.
