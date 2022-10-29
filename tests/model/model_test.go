package model_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/kalogs-c/piadocas/controller"
	"github.com/kalogs-c/piadocas/model"
)

var server = controller.Server{}
var jokeInstance = &model.Joke{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("./../../.env"))
	if err != nil {
		log.Fatalf("Error loading env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	DbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("TEST_DB_NAME"))
	server.DB, err = gorm.Open("mysql", DbURL)
	if err != nil {
		fmt.Println("Cannot connect to the database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Println("We are connected to the database")
	}
}

func refreshJokeTable() error {
	var err error
	err = server.DB.DropTableIfExists(&model.Joke{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&model.Joke{}).Error
	if err != nil {
		return err
	}

	log.Printf("Table refreshed sucessfully.")

	return nil
}

func createJoke() (*model.Joke, error) {
	refreshJokeTable()

	joke := model.Joke{
		Call:   "I don't know how to tell a joke",
		Finish: "Really.",
		Owner:  "testing",
	}

	var err error
	jokeInstance, err = joke.Save(server.DB)
	if err != nil {
		return nil, err
	}

	return jokeInstance, nil
}

func listUsersJoke() (*[]model.Joke, error) {
	usersJokes, err := jokeInstance.CollectUserJokes(server.DB)
	if err != nil {
		return nil, err
	}

	return usersJokes, nil
}

func deleteJoke() (bool, error) {
	err := jokeInstance.Delete(server.DB)
	if err != nil {
		return false, err
	}

	return true, nil
}
