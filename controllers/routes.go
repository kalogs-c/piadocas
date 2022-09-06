package controllers

import "github.com/kalogs-c/piadocas/middlewares"

func (server *Server) initializeRoutes() {
	server.Router.HandlerFunc("/", middlewares.FormatToJSON(server.Home)).Methods("GET")

	server.Router.HandlerFunc("/joke", middlewares.FormatToJSON(server.CreateJoke)).Methods("POST")
	server.Router.HandlerFunc("/joke/{ownerId}", middlewares.FormatToJSON(server.GetJokes)).Methods("GET")
	server.Router.HandlerFunc("/joke/{id}", middlewares.FormatToJSON(server.DeleteJoke)).Methods("DELETE")
}