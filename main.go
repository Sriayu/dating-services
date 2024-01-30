package main

import (
	"dating-services/handler"
	"dating-services/usecases"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	usersUsecase := usecases.NewUsersUsecase()
	allUsecases := usecases.AllUseCases{
		Users: usersUsecase,
	}
	r := makeRoute(allUsecases)
	log.Println("Server Running : http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func makeRoute(usecases usecases.AllUseCases) *chi.Mux {
	r := chi.NewRouter()
	usersHandler := handler.NewUsersHandler(usecases.Users)

	r.Get("/register-users", usersHandler.RegisterUsers)
	return r
}
