package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"graph/logger"
	"graph/pb"
	"graph/service"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	ctx := context.Background()

	var listen   string
	var logLevel string


	flag.StringVar(&listen, "listen", ":9007", "server listener")
	flag.StringVar(&logLevel, "logLevel", "Info", "log level")
	flag.Parse()

	if !strings.HasPrefix(listen, ":"){
		listen = ":" + listen
	}

	// Initialize logger
	logConf := logger.Conf{
		LogLevel:   logLevel,
	}
	logger.InitLogger(logConf)

	// Initialize service
	var graphService service.GraphService
	graphService = service.GraphServiceImpl{}

	// Interrupt handler - a channel for that emits error.
	errChannel := make(chan error)

	go func() {
		logger.GetLogger().Infof("Starting server with Handler %s and port %s.", "gRCP", listen)
		listener, err := net.Listen("tcp", listen)
		if err != nil {
			errChannel <- err
			return
		}
		handler := service.NewGRPCServer(ctx, graphService)
		gRPCServer := grpc.NewServer()
		pb.RegisterGraphServiceServer(gRPCServer, handler)
		logger.GetLogger().Infom("Server started");
		errChannel <- gRPCServer.Serve(listener)
	}()


	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChannel <- fmt.Errorf("%s", <-c)
	}()

	// Run! - Log startup failures if any or interruption to exit
	logger.GetLogger().Error("exit", <-errChannel)
}
