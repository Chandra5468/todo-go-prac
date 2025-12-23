# Domain driven design and Hexagonal Arch 

# To implement a clean, modular and scaled structuring of code

.
├── cmd/
│   └── api/
│       └── main.go           # The "Wiring" only. Connects modules together.
├── internal/
│   ├── platform/             # Infrastructure code (Technicals, not Business)
│   │   ├── postgres/         # Database connection setup (pgx pool)
│   │   ├── logger/           # Logging logic
│   │   └── middleware/       # Global middlewares (Auth, RateLimit)
│   │
│   └── modules/              # THE CORE DOMAIN (Your "Features")
│       ├── users/            # Everything related to Users
│       │   ├── delivery/     # HTTP Handlers & Routes
│       │   ├── service/      # Business Logic
│       │   ├── repository/   # Database Queries
│       │   └── models.go     # Structs & Interfaces specific to Users
│       │
│       ├── todos/            # Everything related to Todos
│       │   ├── delivery/
│       │   ├── service/
│       │   ├── repository/
│       │   └── models.go
│       │
│       └── orders/           # Easy to add new features later!
│           ├── ...
│
├── pkg/                      # Public code you might share with other projects (optional)
├── go.mod
└── .env


.
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── platform/                # Infrastructure
│   │   ├── db/
│   │   │   ├── postgres/
│   │   │   │   └── client.go
│   │   │   └── mongo/
│   │   │       └── client.go
│   │   ├── logger/
│   │   ├── http/
│   │   │   └── middleware/
│   │   └── config/
│   │
│   ├── modules/
│   │   ├── users/
│   │   │   ├── domain/           # Pure business logic
│   │   │   │   ├── entity.go
│   │   │   │   └── repository.go # Interfaces only
│   │   │   │
│   │   │   ├── service/          # Use cases
│   │   │   │   └── service.go
│   │   │   │
│   │   │   ├── transport/
│   │   │   │   └── http/
│   │   │   │       ├── handler.go
│   │   │   │       └── routes.go
│   │   │   │
│   │   │   └── repository/
│   │   │       └── postgres/
│   │   │           └── user_repo.go
│   │   │
│   │   ├── todos/
│   │   └── orders/
│   │
│   └── shared/
│       ├── errors/
│       ├── pagination/
│       └── auth/
│
├── pkg/                         # Truly reusable libs only
├── go.mod
└── .env
