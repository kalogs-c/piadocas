package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kalogs-c/piadocas/model"
)

var jokes []model.Joke = []model.Joke{
	{
		Call:   "Setembrochove?",
		Finish: "Que? Pruuuf",
		Owner:  "kalogs-c",
	},
	{
		Call:   "Qual o animal que só da pra ver com a lupa?",
		Finish: "O micro leão dourado",
		Owner:  "kalogs-c",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&model.Joke{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&model.Joke{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for _, joke := range jokes {
		err = db.Debug().Model(&model.Joke{}).Create(&joke).Error
		if err != nil {
			log.Fatalf("cannot seed jokes table: %v", err)
		}
	}
}
