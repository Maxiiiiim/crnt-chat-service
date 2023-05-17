package main

import (
	"context"
	"fmt"
	"github.com/Maxiiiiim/crnt-chat-service/internal/chat"
	"github.com/Maxiiiiim/crnt-chat-service/internal/middleware/auth"
	"github.com/Maxiiiiim/crnt-chat-service/internal/repository/postgres/dialogs"
	"github.com/Maxiiiiim/crnt-chat-service/internal/repository/postgres/messages"
	"github.com/Maxiiiiim/crnt-chat-service/internal/repository/postgres/users"
	pb "github.com/Maxiiiiim/crnt-chat-service/pkg/api/chat"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	host       string
	port       string
	httpPort   string
	enableAuth string
)

func loadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if host = os.Getenv("HOST"); host == "" {
		log.Fatalln("$HOST environment variable should be specified")
	}
	if port = os.Getenv("PORT"); port == "" {
		log.Fatalln("$PORT environment variable should be specified")
	}
	if httpPort = os.Getenv("HTTP_PORT"); httpPort == "" {
		log.Fatalln("$HTTP_PORT environment variable should be specified")
	}
	if enableAuth = os.Getenv("ENABLE_AUTH"); enableAuth == "" {
		log.Fatalln("$ENABLE_AUTH environment variable should be specified")
	}
	if s := os.Getenv("SYMMETRIC_KEY"); s == "" {
		log.Fatalln("$SYMMETRIC_KEY environment variable should be specified")
	}
}

func createGrpcNetworkListener() net.Listener {
	addr := fmt.Sprintf("%s:%s", host, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start serving grpc with addr: %v\n", addr)
	return listener
}

func main() {
	loadEnvironmentVariables()

	db := getDBConn()
	defer db.Close(context.Background())

	usersRepository := users.NewRepository(db)
	dialogsRepository := dialogs.NewRepository(db)
	messagesRepository := messages.NewRepository(db)

	var interceptors []grpc.UnaryServerInterceptor
	if enableAuth != "false" {
		middleware := auth.NewAuthMiddleware()
		interceptors = append(interceptors, middleware.GetInterceptor())
	}
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors...))

	pb.RegisterChatServiceServer(server, chat.NewServer(&chat.Config{
		Users:    usersRepository,
		Dialogs:  dialogsRepository,
		Messages: messagesRepository,
	}))

	reflection.Register(server)

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		mux := runtime.NewServeMux()

		// creating swagger
		hmux := http.NewServeMux()
		hmux.Handle("/", mux)
		// mount the gRPC HTTP gateway to the root
		fs := http.FileServer(http.Dir("./swagger"))
		hmux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		err := pb.RegisterChatServiceHandlerFromEndpoint(ctx, mux, host+":"+port, opts)
		if err != nil {
			panic(err)
		}
		log.Printf("Start serving http with addr: " + host + ":" + httpPort)
		if err := http.ListenAndServe(":"+httpPort, hmux); err != nil {
			panic(err)
		}
	}()

	listener := createGrpcNetworkListener()
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
