package main

import (
	"log"
	"micro_go/cmd/services1"
	"net/http"
	"os"
	"time"

	// "github.com/gin-gonic/gin"
	"github.com/go-playground/form/v4"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	// g := gin.Default()
	// g.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// g.Run() // listen and serve on 0.0.0.0:8080

	l := log.New(os.Stdout, "cmdh", log.LstdFlags)
	e := form.NewDecoder()

	// create the handlers
	auth := services1.NewAuth(l)
	// create the handlers
	form := services1.NewForm(e)

	// create a new Gorilla Mux router
	r := mux.NewRouter()
	// Contoh log dengan struktur yang diinginkan

	// Define the routes for each handler using Gorilla Mux
	r.HandleFunc("/", auth.ServeHTTP).Methods("GET")
	r.HandleFunc("/form", form.ServeHTTP).Methods("POST")

	// Set up a server using the Gorilla Mux routervsjdbvojdvbwjvbjboeujv
	srv := &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(srv.ListenAndServe(), "localhost:8080")

}
