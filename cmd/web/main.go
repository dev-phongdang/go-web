package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	user_handler "web-api/delivery/http/user"
	user_repository "web-api/repository"
	user_uc "web-api/usecase/user"

	"github.com/go-chi/httplog/v3"
)

// main is the entry point of the application.
func main() {
	logFormat := httplog.SchemaECS.Concise(true)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: logFormat.ReplaceAttr,
	})).With(
		slog.String("app", "wep-api"),
		slog.String("version", "v1.0.0-a1fa420"),
		slog.String("env", "dev"),
	)
	userRepo := user_repository.NewUserMemoryRepository()

	userUserCase := user_uc.NewUserUsecase(userRepo)

	userhandler := user_handler.NewUserHandler(userUserCase)

	router := user_handler.NewRouter(userhandler, logger)

	log.Println("Server starting on port 8087...")
	if err := http.ListenAndServe(":8087", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
