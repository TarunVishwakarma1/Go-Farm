package storage

import (
	"database/sql"
	"fmt"
	"myfarm/farm"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS farmers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS animals (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			type TEXT
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetFarmers(db *sql.DB) ([]farm.Farmer, error) {
	rows, err := db.Query("SELECT id, name FROM farmers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var farmers []farm.Farmer
	for rows.Next() {
		var f farm.Farmer
		if err := rows.Scan(&f.ID, &f.Name); err != nil {
			return nil, err
		}
		farmers = append(farmers, f)
	}
	return farmers, nil
}

func GetAnimals(db *sql.DB) ([]farm.Animal, error) {
	rows, err := db.Query("SELECT id, name, type FROM animals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []farm.Animal
	for rows.Next() {
		var id int
		var name, animalType string
		if err := rows.Scan(&id, &name, &animalType); err != nil {
			return nil, err
		}

		animals = append(animals, &farm.AnimalReponse{ID: id, Name: name, Type: animalType})

	}
	return animals, nil
}

func AddAnimal(db *sql.DB, animal farm.Animal) error {
	res, err := db.Exec("INSERT INTO animals (name, type) VALUES (?, ?)", animal.GetName(), animal.GetType())
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	animal.SetID(int(id))
	return nil
}

func AddFarmer(db *sql.DB, farmer *farm.Farmer) error {
	// Use Exec for INSERT and '?' for parameters to prevent SQL injection.
	res, err := db.Exec("INSERT INTO farmers (name) VALUES (?)", farmer.Name)
	if err != nil {
		return err
	}

	// Get the new ID from the database and update our farmer object.
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	farmer.ID = int(id)
	return nil
}

func GetAnimalByID(db *sql.DB, id int) (farm.Animal, error) {
	var name, animalType string

	err := db.QueryRow("SELECT name, type FROM animals WHERE id = ?", id).Scan(&name, &animalType)
	if err != nil {
		return nil, err
	}

	var animal farm.Animal
	switch animalType {
	case "Cow":
		animal = &farm.Cow{ID: id, Name: name, Type: "Cow"}
	case "Chicken":
		animal = &farm.Chicken{ID: id, Name: name, Type: "Chicken"}
	default:
		return nil, fmt.Errorf("unknown animal type: %s", animalType)
	}
	return animal, nil
}
