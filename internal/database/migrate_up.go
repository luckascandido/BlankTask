package database

import (
	"blanktask/internal/models"
	"context"
	"log"
)

func Migration_up(ctx context.Context) error {
	db, err := Mypg(ctx)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		return err
	}

	log.Println("Migração realizada com sucesso!")
	return nil
}
