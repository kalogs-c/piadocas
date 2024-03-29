package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/kalogs-c/piadocas/model"
	"github.com/kalogs-c/piadocas/responses"
)

func (server *Server) CreateJoke(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	joke := model.Joke{}
	err = json.Unmarshal(body, &joke)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	jokeCreated, err := joke.Save(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, jokeCreated.ID))
	responses.JSON(w, http.StatusCreated, jokeCreated)
}

func (server *Server) GetUserJokes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["owner"]

	joke := model.Joke{Owner: username}
	jokesGotten, err := joke.CollectUserJokes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if len(jokesGotten) == 0 {
		responses.JSON(w, http.StatusNotFound, jokesGotten)
		return
	}

	responses.JSON(w, http.StatusOK, jokesGotten)
}

func (server *Server) GetJokesByLang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	language := vars["language"]

	joke := model.Joke{Language: model.Language(language)}
	jokesGotten, err := joke.CollectJokesByLang(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if len(jokesGotten) == 0 {
		responses.JSON(w, http.StatusNotFound, jokesGotten)
		return
	}

	responses.JSON(w, http.StatusOK, jokesGotten)
}

func (server *Server) GetJokesByTimeRange(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	timeRangeString := vars["time_range"]
	timeRange, err := time.Parse("2006-01-02", timeRangeString)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	joke := model.Joke{CreatedAt: timeRange}
	jokesGotten, err := joke.CollectJokesByTimeRange(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if len(jokesGotten) == 0 {
		responses.JSON(w, http.StatusNotFound, jokesGotten)
		return
	}

	responses.JSON(w, http.StatusOK, jokesGotten)
}

func (server *Server) DeleteJoke(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jokeId, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	joke := model.Joke{ID: uint32(jokeId)}
	err = joke.Delete(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", jokeId))
	responses.JSON(w, http.StatusOK, "Deleted sucessfully")
}
