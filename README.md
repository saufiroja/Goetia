# GO CQRS Microservice Boilerplate

This is a boilerplate for a microservice written in Go. It uses CQRS and Event Sourcing as architectural patterns.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [GoLand](https://www.jetbrains.com/go/) (optional)
- [Postman](https://www.getpostman.com/) (optional)

### Installing

1. Clone the repository

```bash
git clone
```

2. Run the application using Docker Compose

```bash
make docker-up
```

3. Run the application locally

```bash
make dev
```

4Open the application in your browser

```bash
http://localhost:8080
```

## Running the tests

```bash
go test ./...
```

## Built With

- [Go](https://golang.org/) - The programming language used
- [Docker](https://www.docker.com/) - Containerization
- [Docker Compose](https://docs.docker.com/compose/) - Container orchestration
- [GoLand](https://www.jetbrains.com/go/) - IDE
- [Postman](https://www.getpostman.com/) - API testing
- [PostgreSQL](https://www.postgresql.org/) - Database
- [Redis](https://redis.io/) - Cache
- [RabbitMQ](https://www.rabbitmq.com/) - Message broker
- [gRPC](https://grpc.io/) - Remote procedure call framework
- gRPC Gateway - gRPC to HTTP reverse proxy
- Http Router - HTTP request router
- CQRS - Command Query Responsibility Segregation
- Jeager - Distributed tracing
- Prometheus - Monitoring
- Grafana - Metrics visualization
- Logrus - Logging
- Loki - Log aggregation

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details