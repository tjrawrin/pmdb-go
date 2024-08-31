package server

import (
	"net/http"
	"pmdb-go/internal/components"
	"pmdb-go/internal/database"
	"pmdb-go/internal/pages"
	"pmdb-go/web"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func (s *FiberServer) RegisterHandlers() {
	s.App.Get("/", s.IndexHandler)
	s.App.Get("/movies", s.MoviesHandler)
	s.App.Get("/movies/random", s.MoviesRandomHandler)
	s.App.Use("/public", filesystem.New(filesystem.Config{
		Root:       http.FS(web.WebFS),
		PathPrefix: "dist",
		Browse:     false,
	}))

	api := s.App.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/movies", s.MovieListHandler)
	v1.Get("/movies/random", s.MovieItemHandler)
	v1.Post("/movies/search", s.MovieSearchHandler)

	s.App.Use(s.NotFoundHandler)
}

func (s *FiberServer) IndexHandler(c *fiber.Ctx) error {
	return Render(c, pages.Index())
}

func (s *FiberServer) MoviesHandler(c *fiber.Ctx) error {
	view := c.Query("view")
	var movies []database.Movie
	if len(strings.TrimSpace(view)) == 0 || view == "all" {
		movies = s.DB.GetAllMovies()
	} else {
		movies = s.DB.GetMoviesByFirstLetter(view)
	}
	if len(strings.TrimSpace(view)) == 0 {
		c.RedirectToRoute("movies", fiber.Map{
			"queries": map[string]string{"view": "all"},
		})
	}
	return Render(c, pages.Movies(movies, view, len(movies), s.DB.Total))
}

func (s *FiberServer) MoviesRandomHandler(c *fiber.Ctx) error {
	movie := s.DB.GetRandomMovie()
	return Render(c, pages.MoviesRandom(movie))
}

func (s *FiberServer) MovieListHandler(c *fiber.Ctx) error {
	movies := s.DB.GetAllMovies()
	return Render(c, components.MovieList(movies, len(movies), s.DB.Total))
}

func (s *FiberServer) MovieItemHandler(c *fiber.Ctx) error {
	movie := s.DB.GetRandomMovie()
	return Render(c, components.MovieItem(movie))
}

func (s *FiberServer) MovieSearchHandler(c *fiber.Ctx) error {
	if err := c.BodyParser(c); err != nil {
		innerErr := c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		if innerErr != nil {
			panic("Could not send error in HelloWebHandler: " + innerErr.Error())
		}
	}
	search := c.FormValue("search")
	view := c.Query("view")
	var movies []database.Movie
	if len(strings.TrimSpace(view)) == 0 || view == "all" {
		movies = s.DB.GetAllMovies()
	} else {
		movies = s.DB.GetMoviesByFirstLetter(view)
	}
	var filteredMovies []database.Movie
	for _, movie := range movies {
		if strings.Contains(strings.ToLower(movie.SearchableTitle), strings.ToLower(search)) {
			filteredMovies = append(filteredMovies, movie)
		}
	}
	return Render(c, components.MovieList(filteredMovies, len(filteredMovies), s.DB.Total))
}

func (s *FiberServer) NotFoundHandler(c *fiber.Ctx) error {
	return Render(c, pages.NotFound())
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
