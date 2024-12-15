# Todo List Project - HeliTechnology

## Overview
This is a test project created for HeliTechnology, showcasing my skills in developing a scalable and maintainable backend service for managing a Todo List. The project adheres to Hexagonal Architecture principles and integrates with essential services like PostgreSQL and AWS SQS (via LocalStack). It is containerized using Docker for easy setup and deployment.

## Features
- Clean and maintainable Hexagonal Architecture.
- RESTful API endpoints for managing Todo items.
- Integration with PostgreSQL for data persistence.
- AWS SQS integration using LocalStack for local development.
- Docker and Docker Compose support for streamlined deployment.

## Requirements
- [Go](https://golang.org/) (Version 1.20 or higher)
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- PostgreSQL client
- [LocalStack](https://github.com/localstack/localstack) for SQS simulation

## Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/Asef264/todo-list-heli-tech
cd todo-list-heli-tech
```

### 2. Build the Vendor Directory
Ensure all dependencies are vendored for consistency.
```bash
go mod vendor
```

### 3. Setup Environment Variables
Create a `.env` file in the project root with the following content:
```env
STORAGE_TYPE=s3
AWS_ACCESS_KEY_ID=minioadmin
AWS_SECRET_ACCESS_KEY=minioadmin
AWS_REGION=us-east-1
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=heli_tech_backend
POSTGRES_HOST=database
POSTGRES_PORT=5432
SQS_QUEUE_URL=http://localstack:4566/000000000000/todo-queue
```

### 4. Start Docker Services
Build and run the application using Docker Compose:
```bash
docker compose up --build
```
This will start:
- A PostgreSQL database
- LocalStack for SQS
- The Go application

### 5. Create the SQS Queue
After LocalStack is running, create the required SQS queue:
```bash
docker exec -it <container-id> aws --endpoint-url=http://localstack:4566 sqs create-queue --queue-name todo-queue
```
Replace `<container-id>` with the LocalStack container ID.

## Usage
### Endpoints
#### Todo Item Management
1. **Create a Todo Item**
   - **POST** `/todo_items`
   - Body:
     ```json
     {
       "description": "some description",
       "due_data":"someDueDate in time",
       "file_id":"some file id in uuid"
     }
     ```

2. **Upload a File**
   - **POST** `/files`
   - Multipart form-data with `file` field.

3. **Download a File**
   - **GET** `/files/:file_name`

### Running Tests
Run the tests and benchmarks:
```bash
go test ./... -v
```

## Project Structure
```plaintext
.
├── config              # Configuration files and loaders
├── cmd
├── internal            # Application core
│   ├── adapters        # Frameworks, drivers, and gateways
│   ├── ports           # Interfaces for repository and external services
│   ├── service         # Business logic and domain services
├── migrations          # Database migration scripts
├── pkg                 # Shared utilities and helper packages
├── Dockerfile          # Docker image definition
├── docker-compose.yml  # Multi-container Docker configuration
├── makeFile
├── .gitignore
└── README.md           # Project documentation
```

## Notes
- Ensure LocalStack and Docker are properly installed and running.
- If any issues arise, check the logs for troubleshooting:
```bash
docker logs <container-id>
```

