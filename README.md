# Go-PostgreSQL REST API

This project is a RESTful API built with Go (Golang) and PostgreSQL. It performs CRUD operations on user data and includes Swagger documentation for easy interaction and testing.

## Features

- **CRUD Operations**: Create, read, update, and delete user records.
- **Swagger Integration**: Interactive API documentation at `http://localhost:8081/swagger/index.html`.
- **Modular Design**: Clean and scalable codebase.
- **Environment Config**: Uses environment variables for configuration.

## Tech Stack

- **Go (Golang)**
- **PostgreSQL**
- **Gorilla Mux**
- **pq (PostgreSQL driver)**
- **Swagger**
- **godotenv**

## Getting Started

### Prerequisites

- Go 1.16+
- PostgreSQL

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/tharinduyasantha/go-postgresql-rest-api.git
   cd go-postgresql-rest-api
2. **Set up the environment variables**:
   Create a `.env` file in the root directory and add your PostgreSQL credentials:
   ```env
   DB_HOST=your_db_host
   DB_PORT=your_db_port
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
3. **Install dependencies**:
   ```bash
   go mod tidy
3. **Run the application**:
   ```bash
   go run main.go
