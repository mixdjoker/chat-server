package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/mixdjoker/chat-server/internal/config"
	"github.com/mixdjoker/chat-server/internal/lib/handler"
	"github.com/mixdjoker/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.MustConfig()
	iLog := log.New(os.Stdout, color.CyanString("[INFO] "), log.LstdFlags)

	iLog.Println("Starting chat service...")

	url := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.GRPCPort)

	lis, err := net.Listen("tcp", url)
	if err != nil {
		errStr := fmt.Sprintf("failed to listen: %v", err)
		iLog.Fatalf(color.RedString(errStr))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	h := handler.NewChatHandlerV1(iLog)

	chat_v1.RegisterChat_V1Server(s, h)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.Serve(lis); err != nil {
			errStr := fmt.Sprintf("failed to serve: %v", err)
			iLog.Fatalf(color.RedString(errStr))
		}
	}()

	iLog.Println(color.GreenString("Chat service started successfully "), color.BlueString(url))

	<-done
	s.GracefulStop()
	iLog.Println(color.YellowString("Chat service stopped"))
}
