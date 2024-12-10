# todo-list-heli-tech
This repository contains the solution for a test task provided by HeliTechnology as part of the application process.


## folder structure 
todo-service/
├── cmd/                            # Main application entry point
│   └── todo-service/               # Main application entry point folder
│       └── main.go                 # Start the application
├── internal/                       # Core business logic (Hexagonal Architecture)
│   ├── app/                        # Application layer: business logic (use cases)
│   ├── domain/                     # Core domain entities and models
│   ├── ports/                      # Interfaces (ports) for external systems
│   └── adapters/                   # Adapters that implement the ports (infrastructure)
├── api/                            # HTTP handlers, routers, and validation
│   ├── handler/                    # Handlers for API routes
│   ├── router/                     # API routing setup
│   └── validator/                  # Request validation logic
├── config/                         # Configuration settings (env vars, constants)
├── migrations/                     # Database migrations
├── test/                           # Unit and integration tests (new section)
│   ├── app/                        # Tests for the application layer
│   │   ├── create_todo_test.go     # Test for creating TodoItems
│   │   └── todo_service_test.go    # Test for service layer logic (TodoItem management)
│   ├── domain/                     # Tests for domain models
│   │   └── todo_item_test.go       # Unit tests for the TodoItem entity
│   ├── adapters/                   # Tests for adapters (mocking S3, SQS)
│   │   ├── s3_adapter_test.go      # Mock S3 test
│   │   └── sqs_adapter_test.go     # Mock SQS test
│   ├── api/                        # Tests for API endpoints
│   │   ├── todo_handler_test.go    # Tests for /todo and other API routes
│   │   └── file_handler_test.go    # Tests for file upload handling
│   └── integration/                # Integration tests (end-to-end)
│       ├── todo_service_integration_test.go  # Test full end-to-end functionality
│       └── s3_sqs_integration_test.go        # Test integration with S3 and SQS
├── docs/                           # API docs (Swagger/OpenAPI, etc.)
├── scripts/                        # Helper scripts (e.g., for testing or migration)
├── go.mod                          # Go module definition
├── go.sum                          # Go checksum
├── Makefile                        # Makefile for building, testing, running
├── README.md                       # Project documentation
└── docker-compose.yml              # Docker Compose configuration
