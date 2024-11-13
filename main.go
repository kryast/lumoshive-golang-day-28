package main

import (
	"day-28/database"
	"day-28/handler"
	"day-28/repository"
	"day-28/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoTravel := repository.NewTravelRepository(db)
	TravelService := service.NewTravelService(repoTravel)
	TravelHandler := handler.NewTravelHandler(TravelService)

	r := chi.NewRouter()

	r.Get("/travel", TravelHandler.GetTravelHandler)
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
