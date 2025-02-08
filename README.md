# Go Starter Kit

A simple Go project template to help you get started quickly. This starter kit provides a basic setup with essential tools and configurations for building Go web applications, APIs, or microservices.

## Table of Contents

- [Folder Structure](#folder-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [License](#license)

## Folder Structure

```
go-starter-kit/
├── cmd/                  # Contains the entry point and command initialization for the application
│   └── cmd.go
├── configs/
│   └── config.go         # Application configuration
├── internal/
│   ├── handlers/         # Handles HTTP requests (Presentation Layer)
│   │   └── users_handler.go
│   ├── middlewares/      # Handles HTTP middlewares (e.g., authentication)
│   │   └── bearer.go
│   │── models/           # Data models (Domain Layer)
│   │   └── users.go
│   ├── routers/          # Defines HTTP route handlers
│   │   └── router.go
│   ├── repositories/     # Interacts with the database (Persistence Layer)
│   │   └── users_repository.go
│   ├── services/         # Core business logic (Application Layer)
│   │   └── users_service.go
│   └── utils/            # Utility functions
│       └── logger.go
├── migrations/           # Manages database schema changes through migration scripts
├── .dockerignore
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── example-config.json
├── go.mod
├── go.sum
├── main.go               # Entry point of the application
├── Makefile
└── README.md
```

## Installation

1. Make sure you have Golang installed. If not, you can download it from [Golang Official Website](https://go.dev/doc/install).

2. Install 'make' if not already installed. 

    * On Debian/Ubuntu, you can use

    ```bash
    sudo apt-get update
    sudo apt-get install make
    ```

   * On macOS, you can use [Homebrew](https://brew.sh/)

    ```bash
    brew install make
    ```

   * On Windows, you can use [Chocolatey](https://chocolatey.org/)

    ```bash
    choco install make
    ```

3. Clone the repository

    ```bash
    git clone https://github.com/wildanfaz/go-starter-kit.git
    ```

4. Change to the project directory

    ```bash
    cd go-starter-kit
    ```

## Usage

1. Start the application using docker

    ```bash
    docker-compose up
    ```

## Commands

1. Install all dependencies
    ```bash
    make install
    ```

2. Start the application without docker
    ```bash
    make start
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.