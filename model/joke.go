package model

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Joke struct {
	ID        uint32    `json:"id" gorm:"primary_key;unique;auto_increment"`
	Call      string    `json:"call" gorm:"size:255;not null"`
	Finish    string    `json:"finish" gorm:"size:100;not null"`
	Owner     string    `json:"owner" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (j *Joke) prepare() {
	j.Call = html.EscapeString(strings.TrimSpace(j.Call))
	j.Finish = html.EscapeString(strings.TrimSpace(j.Finish))
	j.CreatedAt = time.Now()
}

func (j *Joke) validate() error {
	if j.Call == "" {
		return errors.New("field 'Call' is required")
	}
	if j.Finish == "" {
		return errors.New("field 'Finish' is required")
	}
	return nil
}

func (j *Joke) Save(db *gorm.DB) (*Joke, error) {
	j.prepare()

	var err error

	err = j.validate()
	if err != nil {
		return &Joke{}, err
	}
	err = db.Debug().Create(&j).Error
	if err != nil {
		return &Joke{}, err
	}
	return j, nil
}

func (j *Joke) Delete(db *gorm.DB) error {
	err := db.Debug().Delete(&j).Error
	if err != nil {
		return err
	}
	return nil
}

func (j *Joke) CollectUserJokes(db *gorm.DB) (*[]Joke, error) {
	var err error

	jokes := []Joke{}

	fmt.Println(j.Owner)
	err = db.Debug().Model(&Joke{}).Where("owner = ?", j.Owner).Find(&jokes).Error
	if err != nil {
		return &[]Joke{}, err
	}

	return &jokes, nil
}
