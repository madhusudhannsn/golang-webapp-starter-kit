package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/madhusudhannsn/go-web-app/app/routes"
	"github.com/madhusudhannsn/go-web-app/app/utils/db"
	"github.com/madhusudhannsn/go-web-app/app/utils/logger"
	"github.com/madhusudhannsn/go-web-app/app/utils/config"
)

func main() {
	router := httprouter.New()
	router.GET("/", routes.Index)
	logger.Info.Println("Starting the server on 3000 port")
	log.Fatal(http.ListenAndServe(":3000", router))
}

// Init : This function initializes the logger
func init() {
	log.Println("Environment", os.Getenv("env"))
	if os.Getenv("env") == "" {
		os.Setenv("env", "local")
	}
	log.Println("Initializing the Prerequistes")
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	confErr := config.Init()
	if confErr != nil {
		log.Fatal("Error occured while Loading configuration")
	}
	err, _ := db.GetSession()
	if err != nil {
		log.Fatal("Error occured while creating db session")
	}
}
