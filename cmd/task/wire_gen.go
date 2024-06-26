// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"sonar-bat/internal/conf"
	"sonar-bat/internal/runtime"
	"sonar-bat/internal/task/biz"
	"sonar-bat/internal/task/data"
	"sonar-bat/internal/task/server"
	"sonar-bat/internal/task/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, info runtime.Info) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	taskRepo := data.NewTaskRepo(dataData, logger)
	taskUseCase := biz.NewTaskUseCase(taskRepo, logger)
	taskService := service.NewTaskService(taskUseCase, info)
	subtaskService := service.NewSubtaskService(taskRepo)
	grpcServer := server.NewGRPCServer(confServer, taskService, subtaskService, logger)
	httpServer := server.NewHTTPServer(confServer, taskService, subtaskService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
