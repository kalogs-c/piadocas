package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kalogs-c/piadocas/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open("mysql", DbUrl)
	if err != nil {
		fmt.Println("Cannot connect to mysql database.")
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
