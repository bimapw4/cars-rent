# car-rent

Loan car-rent System is a backend service designed to manage microloan payments in a weekly installment model. The system provides modules for managing borrowers, creating loans with flat interest rates, handling weekly payments, and tracking loan status including outstanding balances and delinquency conditions.

### 1. Project Structure
<pre>
├── bootstrap/                # App initialization (DB, DI, Migrations)
│   ├── db.go
│   ├── migrate.go
│   └── providers.go
│
├── internal/                 
│   ├── business/             # Business logic / usecases
│   ├── common/               # Common helpers (JWT, context, bcrypt, etc.)
│   ├── consts/                # Global constants
│   ├── entity/               # Request payloads / DTO (input layer)
│   ├── handlers/             # HTTP handlers (Fiber endpoints)
│   ├── middleware/           # Middleware (Audit log, Auth guard)
│   ├── presentations/        # DB models & API response structures
│   ├── provider/             # Dependency injection & service registry
│   ├── repositories/         # Data access layer (SQLX + PostgreSQL)
│   ├── response/             # API response wrapper (success / error)
│   ├── routes/                # HTTP route definitions (Fiber)
│   └── migrations/           # SQL migration scripts
│
├── pkg/                      
│   ├── databasex/            # Additional DB helper functions
│   └── meta/                 # Pagination, metadata utilities
│
├── .env                      # Environment variables
├── .env.example              # Sample environment file
├── docker-compose.yml        # Docker service setup
├── dockerfile                 # Dockerfile for app build
├── go.mod                     # Go modules
├── go.sum                     
├── main.go                   # Application entry point
└── readme.md                 
</pre>

### 👀 2. Features
> 
| Module             | Description                                        |
| ------------------ | -------------------------------------------------- |
| **Borrowers**      |Provides functionality to create new borrowers and retrieve a list of all registered borrowers. |
| **Loans**          | Allows creating new loans linked to specific borrowers and retrive loans list   |
| **Payment**     | Enables weekly loan repayments, automatically calculates the current week, and provides payment status checks including outstanding balance and delinquency detection.           |

### 3. Sample Endpoints
| Method | Endpoint               | Description                |
| ------ | ---------------------- | -------------------------- |
| POST   | `/borrower`            | Create the borrowers      |
| GET    | `/borrower`             | Get list the borrowers      |
| POST   | `/loan`                 | Create loans       |
| GET   | `/loan`                 | Get all loans       |
| POST   | `/payment`             | Create payment             |
| GET    | `/payment`               | list payment and schedulle with spesific loan        |
| GET    | `/payment/status/:id` | check deliquent status, deliquent week and outstanding  |

this is the postman collection [Link Download](https://drive.google.com/file/d/1vSGfGyXM6mE5MBrgpbi4KzNK_Gxg9ciQ/view?usp=sharing)


### 4. Run the Project
Without Docker
```
go run main.go
```
##### or
with docker
```
docker-compose build --no-cache
docker-compose up
```

### 5. Migration
```
run the project go run main.go or use docker then the migration will run automatically
```


### 6. Technology Stack
* Golang (1.21+)

* Fiber (HTTP Framework)

* SQLX + PostgreSQL

* Docker / Docker Compose

* Goose Migration


### 7. Env Example
```
APP_NAME = Payroll Payslip
PORT = 8083

DB_HOST = 
DB_USER = 
DB_PASSWORD = 
DB_NAME = 
DB_PORT = 

JWT_SECRET_KEY = 
JWT_LIFESPAN = 
```