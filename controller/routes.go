package controller

import "github.com/kalogs-c/piadocas/middlewares"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", middlewares.FormatToJSON(server.HealthChecker)).Methods("GET")

	server.Router.HandleFunc("/joke", middlewares.FormatToJSON(server.CreateJoke)).Methods("POST")
	server.Router.HandleFunc("/joke/{owner}", middlewares.FormatToJSON(server.GetJokes)).Methods("GET")
	server.Router.HandleFunc("/joke/{id}", middlewares.FormatToJSON(server.DeleteJoke)).Methods("DELETE")
}
