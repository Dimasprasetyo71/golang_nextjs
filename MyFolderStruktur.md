# Project Folder Structure

This document provides an overview of the project folder structure golang .

## Root Folders and Files

- `.github/`: Contains GitHub configuration files such as workflows for CI/CD pipelines.
- `.vercel/`: Configuration files for Vercel deployment.
- `.vscode/`: Settings and configurations for VS Code workspace.
- `cmd/`: Contains the entry point for the application, typically a main file for launching services.
- `config/`: Configuration files and settings (e.g., database configuration, application settings).
- `controllers/`: Controller files that handle requests and responses.
- `deployment/`: Docker, Kubernetes, and other deployment-related files.
- `docs/`: Documentation files, possibly including API documentation.
- `dto/`: Data Transfer Objects, used for structuring data that’s sent or received.
- `go/`: A general folder for Go-related dependencies or tools.
- `helpers/`: Helper functions or utilities used across the application.
- `middleware/`: Middleware functions for request processing (e.g., authentication, logging).
- `models/`: Models representing the structure of database tables or domain objects.
- `repositories/`: Repository layer for data access logic.
- `routes/`: Routing files that define the API endpoints and link them to controllers.
- `services/`: Service layer containing business logic.
- `testing/`: Contains testing files and related configurations.
- `utils/`: Utility functions or classes used throughout the project.
- `validations/`: Validation rules and schemas for input data validation.

## Root Files

- `.env`: Environment variables for sensitive configurations.
- `.gitignore`: Specifies files and folders to be ignored by Git.
- `code_of_conduct.md`: Project code of conduct.
- `CONTRIBUTING.md`: Guidelines for contributing to the project.
- `Dockerfile`: Docker configuration for containerizing the application.
- `go.mod`: Module file listing dependencies.
- `go.sum`: Sum file for dependency verification.
- `LICENSE`: Project license file.
- `LICENSE.txt`: Text-based license file.
- `main.go`: Main entry point for the Go application.
- `makefile`: A makefile for build automation.
- `README.md`: Project overview, setup instructions, and documentation.
- `SECURITY.md`: Security policy or guidelines for reporting vulnerabilities.
- `swagger.yaml`: Swagger configuration file for API documentation.

## Additional Files

<!-- BY DIMAS PRASETYO -->