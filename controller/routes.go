package controller

import "github.com/kalogs-c/piadocas/middlewares"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", middlewares.FormatToJSON(server.HealthChecker)).Methods("GET")

	server.Router.HandleFunc("/joke", middlewares.FormatToJSON(server.CreateJoke)).Methods("POST")
	server.Router.HandleFunc("/joke/lang/{language}", middlewares.FormatToJSON(server.GetJokesByLang)).
		Methods("GET")
	server.Router.HandleFunc("/joke/time/{time_range}", middlewares.FormatToJSON(server.GetJokesByTimeRange)).
		Methods("GET")
	server.Router.HandleFunc("/joke/user/{owner}", middlewares.FormatToJSON(server.GetUserJokes)).
		Methods("GET")
	server.Router.HandleFunc("/joke/{id}", middlewares.FormatToJSON(server.DeleteJoke)).
		Methods("DELETE")
}
