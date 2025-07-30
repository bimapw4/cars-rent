# car-rent.
This is a backend API for a Car Rental System, built using Go and PostgreSQL. It provides complete CRUD functionality for managing cars and rental orders, including features like JWT-based authentication, role-based access control, and automatic price calculation based on rental duration.


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
| **Auth**      |Provides the authentication |
| **Cars**          | Allows creating new cars just role admin, and user only read   |
| **Orders**     | All role can order the cars and than will be count the total payment   |

you can see the api collection on [Link Download](https://drive.google.com/file/d/13rB5j2u0qbHPhRFJebeSWIWHtSr3hrb1/view?usp=sharing)

####  👤 Default Users (Seeded)
| Username     | Password      | Role    |
| ------------ | ------------- | ------- |
| `admin_user` | `akuntest123` | Admin   |
| `user_one`   | `akuntest123` | Regular |
| `user_two`   | `akuntest123` | Regular |
| `user_three` | `akuntest123` | Regular |


* All passwords are securely hashed with bcrypt in the database.

* You can use these accounts to test authentication and role-based access (e.g., admin-only routes).

### 3. Run the Project
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

### 4. Migration
```
run the project go run main.go or use docker then the migration will run automatically
```


### 5. Technology Stack
* Golang (1.21+)

* Fiber (HTTP Framework)

* SQLX + PostgreSQL

* Docker / Docker Compose

* Goose Migration


### 6. Env Example
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