package model_test

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCreateJoke(t *testing.T) {
	joke, err := createJoke()
	if err != nil {
		t.Errorf("this is the error creating an joke: %v\n", err)
		return
	}

	assert.Equal(t, jokeInstance.Call, joke.Call)
	assert.Equal(t, jokeInstance.Finish, joke.Finish)
	assert.Equal(t, jokeInstance.Owner, joke.Owner)
}

func TestListUsersJoke(t *testing.T) {
	usersJoke, err := listUsersJoke()
	if err != nil {
		t.Errorf("this is the error listing the jokes: %v\n", err)
		return
	}

	assert.NotEqual(t, len(*usersJoke), 0)
}

func TestDeleteUser(t *testing.T) {
	isDeleted, err := deleteJoke()
	if err != nil {
		t.Errorf("this is the error listing the jokes: %v\n", err)
		return
	}

	assert.Equal(t, isDeleted, true)
}
