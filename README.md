# Inventory API

The Inventory API is a simple CRUD (Create, Read, Update, Delete) API for managing inventory items. It is built using the [gofr](https://gofr.dev) framework and MySQL database.

## Getting Started

These instructions will guide you through setting up and running the Inventory API on your local machine.

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine
- [Docker](https://www.docker.com/get-started) (optional, for running MySQL in a Docker container)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/inventory-api.git

2. Change into the project directory:

   ```bash
   cd inventory-api

3. Run the following command to install dependencies:

   ```bash
   go mod download

4. Set up your MySQL database. You can use Docker or your local MySQL server.

   ```bash
   docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3307:3306 -d mysql:8.0.28

5. Create a .env file in the project root and configure your database connection:

   ```bash
   DB_HOST=localhost
   DB_USER=root
   DB_PASSWORD=root123
   DB_NAME=test_db
   DB_PORT=3307
   DB_DIALECT=mysql

6. Run the application:

   ```bash
   go run cmd/main.go
