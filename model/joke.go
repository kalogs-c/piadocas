package model

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Language string

const (
	PTBR = "pt-br"
	EN   = "en"
)

type Joke struct {
	ID        uint32    `json:"id"         gorm:"primary_key;unique;auto_increment"`
	Call      string    `json:"call"       gorm:"size:255;not null"`
	Finish    string    `json:"finish"     gorm:"size:100;not null"`
	Owner     string    `json:"owner"      gorm:"not null"`
	Language  Language  `json:"language"   gorm:"not null;size:5"`
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

func (j *Joke) CollectUserJokes(db *gorm.DB) ([]Joke, error) {
	var err error

	jokes := []Joke{}

	err = db.Debug().
		Model(&Joke{}).
		Where("owner = ?", j.Owner).
		Find(&jokes).
		Error
	if err != nil {
		return []Joke{}, err
	}

	return jokes, nil
}

func (j *Joke) CollectJokesByLang(db *gorm.DB) ([]Joke, error) {
	var err error

	jokes := []Joke{}

	err = db.Debug().
		Model(&Joke{}).
		Where("language = ?", j.Language).
		Find(&jokes).
		Error
	if err != nil {
		return []Joke{}, err
	}

	return jokes, nil
}

func (j *Joke) CollectJokesByTimeRange(db *gorm.DB) ([]Joke, error) {
	var err error

	jokes := []Joke{}

	err = db.Debug().
		Model(&Joke{}).
		Where("created_at > ?", j.CreatedAt).
		Find(&jokes).
		Error
	if err != nil {
		return []Joke{}, err
	}

	return jokes, nil
}
