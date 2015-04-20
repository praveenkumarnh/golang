package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"os"
	"web/controllers/frontend"
	"web/helpers/mynegroni"
)

func init() {
	env := os.Getenv("ENV")

	if env == "" {
		panic("ENV not set")
	}

	fmt.Printf("[negroni] %s\n", env)
}

func main() {

	// Negroni
	n := mynegroni.New()

	// Routing
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(mynegroni.NotFound)

	// Routing /
	router.HandleFunc("/", frontend.Index)

	// Routing /login
	router.HandleFunc("/login", frontend.Login)

	// Routing for /profile
	profileRouter := mux.NewRouter()
	profileRouter.HandleFunc("/profile", frontend.Profile)

	// Login authentication required middleware for /profile
	router.Handle("/profile", negroni.New(
		mynegroni.LoginRequired,
		negroni.Wrap(profileRouter),
	))

	render := render.New(render.Options{Layout: "layout"})

	// Add router
	n.UseHandler(router)
	n.UseHandler(frontend.MyHandler(render))

	// Run negroni run!
	n.Run(":3000")
}
