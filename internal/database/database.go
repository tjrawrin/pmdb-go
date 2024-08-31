package database

import (
	"embed"
	"encoding/json"
	"log"
	"math/rand"
	"strings"
)

//go:embed "database.json"
var DatabaseFS embed.FS

type Movie struct {
	ID              int      `json:"id"`
	Title           string   `json:"title"`
	SearchableTitle string   `json:"searchable_title"`
	Format          []string `json:"format"`
}

type Movies struct {
	Movies []Movie `json:"movies"`
}

type Database struct {
	Movies []Movie
	Total  int
}

func New() *Database {
	data, err := DatabaseFS.ReadFile("database.json")
	if err != nil {
		log.Fatal("Error reading database file " + err.Error())
	}
	var m Movies
	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Fatal("Error unmarshalling database " + err.Error())
	}
	return &Database{
		Movies: m.Movies,
		Total:  len(m.Movies),
	}
}

func (db *Database) GetAllMovies() []Movie {
	return db.Movies
}

func (db *Database) GetMoviesByFirstLetter(c string) []Movie {
	var movies []Movie
	for _, movie := range db.Movies {
		if strings.HasPrefix(strings.ToLower(movie.Title), strings.ToLower(c)) {
			movies = append(movies, movie)
		}
	}
	return movies
}

func (db *Database) GetRandomMovie() Movie {
	return db.Movies[rand.Intn(db.Total)]
}
