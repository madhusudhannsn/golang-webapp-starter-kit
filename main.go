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
)

func main() {
	Init()
	router := httprouter.New()
	router.GET("/", routes.Index)
	logger.Info.Println("Starting the server on 3000 port")
	log.Fatal(http.ListenAndServe(":3000", router))
}

// Init : This function initializes the logger
func Init() {
	log.Println("Initializing the startups")
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	err, _ := db.GetSession()
	if err != nil {
		log.Fatal("Error occured while creating db session")
	}
}
