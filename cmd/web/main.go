package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/furmedia/bookings/config"

	"github.com/furmedia/bookings/pkg/render"

	"github.com/furmedia/bookings/pkg/handler"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.PortNumber = ":8080"

	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/About", handler.Repo.About)

	fmt.Println("Starting server at localhost:8080")
	//_ = http.ListenAndServe(":8080", nil)
	srv := &http.Server{
		Addr:    app.PortNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
