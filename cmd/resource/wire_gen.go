// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"sonar-bat/internal/conf"
	"sonar-bat/internal/resource/biz"
	"sonar-bat/internal/resource/data"
	"sonar-bat/internal/resource/server"
	"sonar-bat/internal/resource/service"
	"sonar-bat/internal/runtime"
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
	hostRepo := data.NewHostRepo(dataData, logger)
	resource := biz.NewHostManager(hostRepo, logger)
	resourceService := service.NewResourceService(resource, info)
	grpcServer := server.NewGRPCServer(confServer, resourceService, logger)
	httpServer := server.NewHTTPServer(confServer, resourceService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
