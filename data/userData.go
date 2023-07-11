package data

import (
	"animal-game/initializers"
	"animal-game/models"
	"log"
)

func CreateUser(id string, userName string) error {

	user := models.User{ID: id, UserName: userName}
	insertResult := initializers.DB.Create(&user)

	if insertResult.Error != nil {
		log.Fatal("error inserting user in the database")
	}
	return nil
}
