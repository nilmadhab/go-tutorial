# Go Learning Roadmap

Based on the e-commerce project built with Go backend and Vue frontend.

---

## Current Skills Covered

- Basic HTTP server with `net/http`
- REST API endpoints (GET, POST, PUT, DELETE)
- PostgreSQL integration with `database/sql`
- JSON encoding/decoding
- CORS middleware
- Environment variables

---

## Phase 1: Strengthen Fundamentals

**Goal:** Master Go basics before moving to advanced topics

- [ ] Go syntax deep dive (pointers, slices, maps, structs)
- [ ] Error handling patterns (custom errors, error wrapping with `%w`)
- [ ] Interfaces and type assertions
- [ ] Go modules and dependency management
- [ ] Unit testing with `testing` package
- [ ] Table-driven tests
- [ ] Benchmarking

**Resources:**
- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)

---

## Phase 2: Better Project Structure

**Goal:** Organize code for maintainability and scalability

**Current state:** Everything in `main.go`

**Target structure:**
```
/cmd/server/main.go        # Entry point
/internal/
  /handlers/               # HTTP handlers
  /models/                 # Data structures
  /repository/             # Database operations
  /middleware/             # Auth, logging, etc.
  /config/                 # Configuration
/pkg/utils/                # Reusable utilities
/migrations/               # Database migrations
```

**Topics to learn:**
- [ ] Package organization
- [ ] Configuration management (env files, Viper)
- [ ] Dependency injection
- [ ] Middleware patterns (logging, auth, rate limiting)
- [ ] Input validation
- [ ] Custom error types

---

## Phase 3: Database & ORM

**Goal:** Efficient and safe database operations

**Current state:** Raw SQL with `database/sql`

**Topics to learn:**
- [ ] GORM or sqlx for easier database operations
- [ ] Database migrations (golang-migrate)
- [ ] Transactions
- [ ] Connection pooling
- [ ] Query optimization
- [ ] Prepared statements
- [ ] Handling NULL values

**Example with GORM:**
```go
type Product struct {
    gorm.Model
    Name        string  `gorm:"not null"`
    Price       float64 `gorm:"not null"`
    Stock       int     `gorm:"default:0"`
}

// Auto-migrate
db.AutoMigrate(&Product{})

// Query
var products []Product
db.Where("stock > ?", 0).Find(&products)
```

---

## Phase 4: Web Frameworks & Routing

**Goal:** Use production-ready frameworks for better developer experience

**Current state:** `net/http` standard library

**Popular frameworks:**
| Framework | Description |
|-----------|-------------|
| Gin | Fast, popular, good middleware support |
| Echo | High performance, minimalist |
| Chi | Lightweight, idiomatic, composable |
| Fiber | Express-inspired, very fast |

**Topics to learn:**
- [ ] Router with path parameters
- [ ] Request binding and validation
- [ ] Grouped routes
- [ ] Swagger/OpenAPI documentation
- [ ] API versioning

**Example with Gin:**
```go
r := gin.Default()

r.GET("/products", getProducts)
r.GET("/products/:id", getProduct)
r.POST("/products", createProduct)
r.PUT("/products/:id", updateProduct)
r.DELETE("/products/:id", deleteProduct)

r.Run(":8080")
```

---

## Phase 5: Authentication & Security

**Goal:** Secure your application

**Topics to learn:**
- [ ] JWT authentication
- [ ] Password hashing (bcrypt)
- [ ] Session management
- [ ] OAuth2 integration (Google, GitHub login)
- [ ] HTTPS/TLS
- [ ] Input sanitization
- [ ] CSRF protection
- [ ] Rate limiting
- [ ] API key authentication

**JWT Example:**
```go
// Generate token
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "user_id": user.ID,
    "exp":     time.Now().Add(24 * time.Hour).Unix(),
})
tokenString, _ := token.SignedString([]byte(secretKey))

// Verify token
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte(secretKey), nil
})
```

---

## Phase 6: Concurrency

**Goal:** Leverage Go's concurrency primitives

**This is Go's superpower!**

**Topics to learn:**
- [ ] Goroutines
- [ ] Channels (buffered and unbuffered)
- [ ] Select statement
- [ ] sync package (Mutex, RWMutex, WaitGroup, Once)
- [ ] Context for cancellation and timeouts
- [ ] Worker pools
- [ ] Fan-out/Fan-in patterns

**Example - Worker Pool:**
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}
```

**Project ideas:**
- Background job processing
- Email notification service
- Image resizing pipeline
- Web scraper

---

## Phase 7: Advanced Topics

**Goal:** Build production-grade applications

### Caching with Redis
```go
rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})

// Set
rdb.Set(ctx, "product:1", productJSON, time.Hour)

// Get
val, err := rdb.Get(ctx, "product:1").Result()
```

### Message Queues
- [ ] RabbitMQ
- [ ] Apache Kafka
- [ ] NATS

### Real-time Features
- [ ] WebSockets with Gorilla WebSocket
- [ ] Server-Sent Events (SSE)

### Microservices
- [ ] gRPC for service-to-service communication
- [ ] Protocol Buffers
- [ ] Service discovery
- [ ] API Gateway

### GraphQL
- [ ] gqlgen for type-safe GraphQL

---

## Phase 8: DevOps & Deployment

**Goal:** Deploy and monitor your applications

### Containerization
```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

### Topics to learn:
- [ ] Docker containerization
- [ ] Docker Compose for local development
- [ ] CI/CD pipelines (GitHub Actions, GitLab CI)
- [ ] Kubernetes basics
- [ ] Health checks and graceful shutdown
- [ ] Monitoring (Prometheus, Grafana)
- [ ] Structured logging (zerolog, zap)
- [ ] Distributed tracing (Jaeger, OpenTelemetry)

---

## Suggested Project Progression

| # | Project | Skills Practiced |
|---|---------|-----------------|
| 1 | Add user auth to e-commerce | JWT, bcrypt, middleware |
| 2 | Build a CLI tool | Cobra, file I/O, flags |
| 3 | URL shortener | Redis caching, unique ID generation |
| 4 | Real-time chat | WebSockets, goroutines, channels |
| 5 | Task queue system | Redis/RabbitMQ, worker pools |
| 6 | Microservices app | gRPC, Docker, service discovery |
| 7 | Kubernetes deployment | K8s, Helm, monitoring |

---

## Recommended Resources

### Official Documentation
- [Go Documentation](https://go.dev/doc/)
- [Go Blog](https://go.dev/blog/)
- [Go Wiki](https://github.com/golang/go/wiki)

### Books
- "Learning Go" by Jon Bodner
- "Concurrency in Go" by Katherine Cox-Buday
- "Let's Go" by Alex Edwards
- "Let's Go Further" by Alex Edwards

### Online Courses
- [Exercism Go Track](https://exercism.org/tracks/go) - Free
- [Gophercises](https://gophercises.com/) - Free video course
- [Ardan Labs Ultimate Go](https://www.ardanlabs.com/training/)

### YouTube Channels
- Melkey
- Anthony GG
- NerdCademy

### Community
- [Gophers Slack](https://gophers.slack.com/)
- [r/golang](https://www.reddit.com/r/golang/)
- [Go Forum](https://forum.golangbridge.org/)

---

## Progress Tracker

### Completed
- [x] Basic HTTP server
- [x] REST API (CRUD)
- [x] PostgreSQL integration
- [x] JSON handling
- [x] CORS middleware

### In Progress
- [ ] _Update as you learn_

### Next Up
- [ ] _Plan your next topic_

---

*Last updated: August 2026*
