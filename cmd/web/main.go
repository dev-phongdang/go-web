package main

import (
	"log"
	"net/http"
	"os"
	auth_router "web-api/delivery/http/auth"
	"web-api/delivery/http/health"
	"web-api/delivery/http/routers"
	user_handler "web-api/delivery/http/user"
	user_repository "web-api/repository"
	auth_uc "web-api/usecase/auth"
	user_uc "web-api/usecase/user"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

// main is the entry point of the application.
func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	routers := routers.NewRouters("/api/v1", chi.NewRouter())

	userRepo := user_repository.NewUserMemoryRepository()
	userUserCase := user_uc.NewUserUsecase(userRepo)
	userhandler := user_handler.NewUserHandler(userUserCase)
	user_router := user_handler.NewRouter(userhandler)

	auth_usecase, err := auth_uc.NewPasetoMaker(os.Getenv("APP_SECRET"))
	if err != nil {
		log.Fatalf("could not create auth usecase: %v", err)
		panic("could not create auth usecase")
	}

	auth_router := auth_router.NewRouter(auth_router.NewAuthHandler(auth_usecase))

	routers.Register("/health", health.NewRouter())
	routers.Register("/users", user_router)
	routers.Register("/auth", auth_router)
	log.Println("Server starting on port " + port + "...")
	if err := http.ListenAndServe(":"+port, routers.Root); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
