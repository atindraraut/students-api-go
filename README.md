# Students API

## Overview
The Students API is a RESTful web service built using Go. It provides functionality to manage student data, including creating, retrieving, updating, and deleting student records. The project is designed with a modular structure, making it easy to extend and maintain.

## Features
- CRUD operations for student records.
- Configuration management using YAML files.
- SQLite database integration for persistent storage.
- Modular and clean code structure.
- Utility functions for standardized API responses.

## Project Structure
```
go.mod
cmd/
    students-api/
        main.go
config/
    local.yaml
internal/
    config/
        config.go
    http/
        handlers/
            student/
                student.go
    types/
        types.go
    utils/
        response/
            response.go
storage/
    storage.db
    storage.go
    sqlite/
        sqlite.go
```

### Key Directories and Files
- **cmd/students-api/main.go**: Entry point of the application.
- **config/local.yaml**: Configuration file for the application.
- **internal/config/config.go**: Handles application configuration.
- **internal/http/handlers/student/student.go**: Contains HTTP handlers for student-related operations.
- **internal/utils/response/response.go**: Utility functions for API responses.
- **storage/sqlite/sqlite.go**: SQLite database integration.

## Prerequisites
- Go 1.20 or later
- SQLite

## Getting Started

### Clone the Repository
```bash
git clone https://github.com/your-username/students-api-go.git
cd students-api-go
```

### Install Dependencies
```bash
go mod tidy
```

### Run the Application
```bash
make run
```
Alternatively, you can run the application directly using Go:
```bash
go run cmd/students-api/main.go -config config/local.yaml
```

### Configuration
The application uses a YAML configuration file located at `config/local.yaml`. Update this file to configure the application (e.g., database connection settings).

### Database
The project uses SQLite for data storage. The database file is located at `storage/storage.db`. If the file does not exist, it will be created automatically when the application starts.

## API Endpoints

### Base URL
`http://localhost:8080`

### Endpoints
- **GET /students**: Retrieve a list of all students.
- **GET /students/{id}**: Retrieve a specific student by ID.
- **POST /students**: Create a new student.
- **PUT /students/{id}**: Update an existing student.
- **DELETE /students/{id}**: Delete a student by ID.

## Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push the branch.
4. Open a pull request.


## Contact
For any questions or feedback, please contact [atindraraut80@gmail.com].