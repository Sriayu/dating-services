package main

import (
	"dating-services/src/handler"
	"dating-services/src/handler/response"
	"dating-services/src/infra"
	"dating-services/src/infra/persistence/postgres"
	"dating-services/src/usecases"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := infra.Make()
	postgresDb := postgres.New(config.SqlDb)
	userRepo := postgres.NewUsersPersistence(postgresDb.DB)
	dateRepo := postgres.NewDatesPersistence(postgresDb.DB)
	packagesRepo := postgres.NewPackagesPersistence(postgresDb.DB)
	packageOrdersRepo := postgres.NewPackageOrdersPersistence(postgresDb.DB)
	cartsRepo := postgres.NewCartsPersistence(postgresDb.DB)
	allUsecases := usecases.AllUseCases{
		Users: usecases.NewUsersUsecase(userRepo, packagesRepo, cartsRepo, packageOrdersRepo),
		Dates: usecases.NewDatesUsecase(userRepo, dateRepo),
	}
	r := makeRoute(allUsecases)
	log.Println("Server Running : http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func makeRoute(usecases usecases.AllUseCases) *chi.Mux {
	r := chi.NewRouter()
	respClient := response.NewResponseClient()
	usersHandler := handler.NewUsersHandler(usecases.Users, respClient)
	datesHandler := handler.NewDatesHandler(usecases.Dates, respClient)

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", usersHandler.RegisterUsers)
		r.Post("/login", usersHandler.LoginUser)
		r.Patch("/{id}/upgrade", usersHandler.UpgradeUser)
		r.Patch("/{id}/pay-upgrade", usersHandler.PayUpgradeUser)
	})

	r.Route("/dating", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/list", datesHandler.GetDateList)
			r.Patch("/update-status", datesHandler.UpdateStatusDate)

		})
		r.Get("/{datingid}/detail", datesHandler.GetDateDetail)
	})

	return r
}
