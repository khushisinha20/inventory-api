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
   git clone https://github.com/khushisinha20/inventory-api.git

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

The API should now be running at http://localhost:9000.

### API Endpoints Testing

1. Add Inventory
   ![image](https://github.com/khushisinha20/inventory-api/assets/79047990/c27f84a5-aa5d-4286-b0a8-ea3375608711)

2. Get All Inventories
   ![image](https://github.com/khushisinha20/inventory-api/assets/79047990/d4e57f59-d5f1-4fd9-96f7-a8bfd88e0588)

3. Get Inventory by ID
   ![image](https://github.com/khushisinha20/inventory-api/assets/79047990/17ca7f2a-3e1e-47b8-927c-bcbe908a57f9)

4. Update Inventory
   ![image](https://github.com/khushisinha20/inventory-api/assets/79047990/1392d877-694f-49af-966c-f01b8e3c8df3)

5. Delete Inventory
   ![image](https://github.com/khushisinha20/inventory-api/assets/79047990/cda28074-c91e-4d8f-a936-f21e1a37f486)

   ![image](https://github.com/khushisinha20/inventory-api/assets/79047990/8622c175-5077-4e89-af4f-0d1dd635a5a3)

### Sequence Diagram
![image](https://github.com/khushisinha20/inventory-api/assets/79047990/22014389-1082-4567-b277-f6e572c4f3ff)

