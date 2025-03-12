# Max Bench Press Calculator - calculator for Estimating One-Rep Max in Bench Press

**Max Bench Press Calculator** is a REST API service for calculating your maximum bench press based on input data. The service allows users to register and receive accurate calculations of their potential maximum in bench press, as well as compare their results with the average. The project is developed in **Golang**, which ensures high performance and reliability.

## 🔧 Key Features
- **Modern Technologies**: Uses the `go-chi/chi` HTTP router for flexible request routing.
- **PostgreSQL Database**: Reliable storage of user data and their results.
- **Swagger Documentation**: Automatic API documentation generation for easy integration.
- **Docker & Docker Compose**: Easy deployment and scaling.
- **Migrations**: Automatic database structure management.
- **Logging**: Integration with `slog` for efficient application monitoring.
- **Testing**: Use of mock objects for reliable functionality testing.
- **Configuration via ENV**: Flexible setup through environment variables.

## 📜 API Overview
- **POST /create**: Create a user and calculate the maximum bench press.

## 🗂️ Project Structure
```plaintext
.
├── calculator
├── cmd
│   ├── calculator
│   │   └── main.go
│   └── migrator
│       └── main.go
├── config
│   └── local.yaml
├── docker-compose.yaml
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── http
│   │   ├── handlers
│   │   └── responce
│   │       └── responce.go
│   ├── lib
│   │   └── logger
│   │       └── sl
│   │           └── sl.go
│   ├── model
│   │   ├── stat.go
│   │   └── user.go
│   ├── service
│   │   ├── calculator
│   │   │   └── service.go
│   │   └── service.go
│   ├── storage
│   │   ├── errors.go
│   │   ├── postgresql
│   │   │   ├── storage.go
│   │   │   └── userrepository.go
│   │   ├── storage.go
│   │   └── userrepository.go
│   └── transport
│       └── http
│           ├── create.go
│           ├── create_test.go
│           └── domain.go
├── local.env
├── Makefile
├── migrations
│   └── postgres
│       ├── 20250206114415_create_users.down.sql
│       └── 20250206114415_create_users.up.sql
└── README.md
```

## Installation and Running

To install and run the **Max Bench Press Calculator**, follow these steps:

### 1. Clone the repository:
```bash
git clone https://github.com/wehw93/bench-press-calculator.git
cd bench-press-calculator
```

### 2. Install dependencies:
```bash
go mod tidy
```

### 3. Configure:
- Use the `local.env` file with the necessary environment variables or edit `config/local.yaml`
- Default config file (`config/local.yaml`):
```yaml
env: "local"
db:
  host: "calculator_db"
  port: "5432"
  name: "calculator_db"
  user: "calc_user"
  password: "pwd123"
  sslmode: "disable"

  http_server:
    address: "0.0.0.0:8080"
    timeout: "4s"
    idle_timeout: "60s"
```

### 4. Start the database and run migrations:
```bash
docker-compose up -d --build
```


### 5. Test the API:
- You can use curl or Postman for testing:
- Create a user and calculate the maximum bench press:
  ```bash
  curl  -X POST http://localhost:8080/create -H "Content-Type: application/json" -d '{"email":"example@email.ru", "password":"password123", "weight":100, "quantity":10}'
  ```

### 6. API Documentation:
- Swagger UI is available at:
  ```
  http://localhost:8080/swagger/
  ```

## Skills I Developed While Working on This Project:

- **REST API**: Development and integration of RESTful services using HTTP methods.
- **Swagger**: Creating API documentation using annotations and automatic generation.
- **PostgreSQL**: Working with a relational database for storing user data.
- **Migrations**: Managing database structure using migration scripts.
- **Mock Testing**: Writing tests using mock objects for component isolation.
- **Docker & Docker Compose**: Setting up containerization to simplify deployment.
- **Environment Variables**: Managing application configuration through environment variables.
- **Logging**: Using the modern slog logger for efficient application monitoring.
- **Validation**: Checking input data to ensure correct service operation.
- **Dependency Injection**: Implementing dependency injections to improve code testability.
- **Password Hashing**: Secure storage of user data.
