package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"rarebnb/internal/config"
	"rarebnb/internal/driver"
	"rarebnb/internal/handlers"
	"rarebnb/internal/helpers"
	"rarebnb/internal/models"
	"rarebnb/internal/render"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

//main is the main application function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	defer close(app.MailChan)

	fmt.Println("Starting mail listener...")
	listenForMail()

	// msg := models.MailData {
	// 	To: "ssmal@niceguy.com",
	// 	From: "bill@bulletz.com", 
	// 	Subject: "Hows ya Muvva?",
	// 	Content: "",
	// }

	// app.MailChan <- msg
		

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	//change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	//Connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=grandmastahash password=")
	if err != nil {
		log.Fatal("Unable to connect to database. Terminating process...")
	}
	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
		return db, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
