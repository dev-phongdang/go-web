package main

import (
	"log"
	"net/http"
	user_handler "web-api/delivery/http/user"
	user_repository "web-api/repository"
	user_uc "web-api/usecase/user"
)

// main is the entry point of the application.
func main() {
	userRepo := user_repository.NewUserMemoryRepository()

	userUserCase := user_uc.NewUserUsecase(userRepo)

	userhandler := user_handler.NewUserHandler(userUserCase)

	router := user_handler.NewRouter(userhandler)

	log.Println("Server starting on port 8087...")
	if err := http.ListenAndServe(":8087", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
