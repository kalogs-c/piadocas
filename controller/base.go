package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/kalogs-c/piadocas/model"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	var err error

	server.DB, err = gorm.Open("postgres", os.Getenv("DB_CONN"))
	if err != nil {
		fmt.Println("Cannot connect to postgres database.")
		log.Fatal("Failed to connect to db: ", err)
	}
	fmt.Println("Connected to mysql database.")

	server.DB.Debug().AutoMigrate(&model.Joke{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
