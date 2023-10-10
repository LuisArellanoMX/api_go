package common

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDBCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func InitDB() error {
	// Connecting to MongoDB
	uri := os.Getenv("CONNECTION_STRING")
	if uri == "" {
		return errors.New("error to get enviroment 'CONNECTION_STRING' ")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Inicializando la base de datos
	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		return errors.New("error to get enviroment 'DB_NAME' ")
	}

	db = client.Database(db_name)

	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
}
