package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/AndriyKalashnykov/gqlgen-graphql-tutorial/domain"
	"github.com/AndriyKalashnykov/gqlgen-graphql-tutorial/graphql"
	customMiddleware "github.com/AndriyKalashnykov/gqlgen-graphql-tutorial/middleware"
	"github.com/AndriyKalashnykov/gqlgen-graphql-tutorial/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "meetmeup_dev",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userRepo := postgres.UsersRepo{DB: DB}

	router := chi.NewRouter()

	//router.Use(cors.New(cors.Options{
	//	AllowedOrigins:   []string{"http://localhost:8080"},
	//	AllowCredentials: true,
	//	Debug:            true,
	//}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	d := domain.NewDomain(userRepo, postgres.MeetupsRepo{DB: DB})

	c := graphql.Config{Resolvers: &graphql.Resolver{Domain: d}}

	queryHandler := handler.GraphQL(graphql.NewExecutableSchema(c))

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", graphql.DataloaderMiddleware(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
