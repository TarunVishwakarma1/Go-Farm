package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"myfarm/farm"
	"myfarm/storage"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type App struct {
	DB *sql.DB
}

func main() {
	db, err := storage.InitDB("./farm.db")
	if err != nil {
		log.Fatalf("Failed to initlized database: %v", err)
	}
	defer db.Close()

	app := &App{DB: db}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", app.welcomeHandler)

	r.Route("/api", func(r chi.Router) {
		r.Get("/farmers", app.getFarmerHandler)
		r.Get("/animals", app.getAnimalHandler)
		r.Post("/farmers", app.addFarmerHandler)
		r.Post("/animals", app.addAnimalHandler)
		r.Get("/animals/{animalId}", app.getAnimalHandlerById)
	})

	fmt.Println("API server is running at http://localhost:8080")
	http.ListenAndServe(":8080", r)

}

func (app *App) welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to the Farm API</h1><p>Your data is now saved in farm.db!</p>")
}

func (app *App) getAnimalHandler(w http.ResponseWriter, r *http.Request) {
	animals, err := storage.GetAnimals(app.DB)
	if err != nil {
		http.Error(w, "Failed to fetch animals", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(animals)
}

func (app *App) getFarmerHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	farmers, err := storage.GetFarmers(ctx, app.DB)
	if err != nil {
		http.Error(w, "Failed to fetch farmers", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(farmers)
}

func (app *App) addFarmerHandler(w http.ResponseWriter, r *http.Request) {
	var farmer farm.Farmer
	if err := json.NewDecoder(r.Body).Decode(&farmer); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if farmer.Name == "" {
		http.Error(w, "Name is Required", http.StatusBadRequest)
		return
	}

	if err := storage.AddFarmer(app.DB, &farmer); err != nil {
		http.Error(w, "Could not save farmer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(farmer)
}

type addAnimalRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (app *App) addAnimalHandler(w http.ResponseWriter, r *http.Request) {
	var req addAnimalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var animal farm.Animal
	switch req.Type {
	case "Cow":
		animal = &farm.Cow{Name: req.Name, Type: "Cow"}
	case "Chicken":
		animal = &farm.Chicken{Name: req.Name, Type: "Chicken"}
	default:
		http.Error(w, "Invalid animal type. Must be 'Cow' or 'Chicken'", http.StatusBadRequest)
		return
	}

	if err := storage.AddAnimal(app.DB, animal); err != nil {
		http.Error(w, "Could not save animal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(animal)
}

func (app *App) getAnimalHandlerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "animalId"))

	if err != nil {
		http.Error(w, "Wrong id provided could not parse id", http.StatusBadRequest)
		return
	}
	animal, err := storage.GetAnimalByID(app.DB, id)

	if err != nil {
		http.Error(w, "Some error occurred while getting animal by id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(animal)
}
