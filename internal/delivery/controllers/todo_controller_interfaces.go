package controllers

import "github.com/saufiroja/cqrs/internal/grpc"

type ITodoController interface {
	grpc.TodosServer
}
