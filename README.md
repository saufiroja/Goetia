# Goetia
Goetia is a repository for Learning the Principles of using gRPC and Rest HTTP

The principle of using gRPC and Rest HTTP in 1 application is to use a grpc gateway. 
Grpc gateway is a proxy that can be used to access the gRPC service using the Rest HTTP protocol. 
This repository can be used as a reference to learn the principles of using gRPC and Rest HTTP in 1 application.

## Upcoming Features
- Unit Testing (on progress)
- Integration Testing
- End to End Testing
- Deployment to Kubernetes k3s (on research)

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
git clone https://github.com/saufiroja/Goetia.git
cd Goetia
```
2. Install dependencies

```bash
go mod tidy
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.0
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```

### Usage

1. To Generate proto file

```bash
make protoc
```

2. Run the application using docker-compose

```bash
make docker-up
```

3. Open grafana dashboard

```bash
username: admin
password: admin
http://localhost:3000
```

4. Run unit test

```bash
make unit-test
```

### Endpoints
| METHOD |        Endpoint         |       Description        |
|:------:|:-----------------------:|:------------------------:|
|  GET   |       {url}/todos       |    Get all todos list    |
|  GET   |    {url}/todos/{id}     |      Get todo by id      |
|  PUT   |    {url}/todos/{id}     |       Update todo        |
| PATCH  | {url}/todos/{id}/status | Update todo status by id |
|  POST  |       {url}/todos       |       Create todo        |
| DELETE |    {url}/todos/{id}     |    Delete todo by id     |


## Tech Stack

- ✅ Go - The programming language used
- ✅ Docker - Containerization
- ✅ Docker Compose - Container orchestration
- ✅ PostgresSQL- Database
- ✅ Redis- Cache
- ✅ gRPC - Remote procedure call framework
- ✅ gRPC Gateway - gRPC to HTTP reverse proxy
- ✅ Http Server - HTTP server
- ✅ Jaeger - Distributed tracing
- ✅ Prometheus - Metrics
- ✅ Grafana - Visualization
- ✅ Logrus - Logging 
- ✅ Loki - Log aggregation
- ✅ CI/CD - Continuous integration and delivery
- ✅ Unit Testing - Testing
- [ ] Kubernetes - Container orchestration (on research)
- [ ] Integration Testing - Testing
- [ ] End to End Testing - Testing

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
