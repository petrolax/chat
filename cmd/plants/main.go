package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/oklog/oklog/pkg/group"
	"github.com/petrolax/chat/config"
	privateApi "github.com/petrolax/chat/pkg/api/private"
	"github.com/petrolax/chat/pkg/plants"
	"github.com/petrolax/chat/pkg/repository"
	"github.com/petrolax/chat/pkg/transport/endpoints/private"
	"github.com/petrolax/chat/pkg/transport/endpoints/public"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Can't parse config")
	}
	fmt.Println(cfg)

	db, err := gorm.Open("postgres", cfg.DB.GetDSN())
	if err != nil {
		log.Fatalln(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// init repository
	repo := repository.NewRepository(db, logger.With(zap.String("component", "repository")))

	// init service
	service := plants.NewService(logger.With(zap.String("component", "service")), repo)

	// init controller
	publicController := public.NewController(logger.With(zap.String("component", "public")), service)
	privateController := private.NewController(logger.With(zap.String("component", "private")), service)

	var g group.Group
	server := &http.Server{
		Addr:    cfg.HttpPort,
		Handler: publicController.Endpoints(),
	}
	grpcListener, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.Fatalln("Can't listen grpc port")
	}
	grpcServer := grpc.NewServer()

	// ctrl+c/ctrl+z and other terminate hotkeys, which kill proccess
	trap := make(chan struct{})
	g.Add(
		func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("%s", sig)
			case <-trap:
				return nil
			}
		},
		func(err error) {
			close(trap)
		},
	)

	// http server (public)
	g.Add(
		func() error {
			logger.Info("run http server", zap.String("transport", "http"), zap.String("port", cfg.HttpPort))
			return server.ListenAndServe()
		}, func(err error) {
			logger.Error(err.Error())
			server.Close()
		},
	)

	// grpc server (private)
	g.Add(
		func() error {
			logger.Info("run grpc server", zap.String("transport", "grpc"), zap.String("port", cfg.GrpcPort))
			privateApi.RegisterPlantsApiServer(grpcServer, privateController)
			reflection.Register(grpcServer) // for postman and other request utils
			return grpcServer.Serve(grpcListener)
		}, func(err error) {
			logger.Error(err.Error())
			grpcListener.Close()
		},
	)

	logger.Error("Application exit", zap.Error(g.Run()))
}
