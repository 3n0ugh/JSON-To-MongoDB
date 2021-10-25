package database

import (
	"context"
	"fmt"
	"time"

	"github.com/3n0ugh/JSON-to-MongoDB/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance

// Connect configures the MongoDB client and initializes the database connection.
func Connect() error {

	config, err := config.GetConfig()
	if err != nil {
		return err
	}
	mongoURI := fmt.Sprintf("mongodb://%s:%s/%s", config.DatabaseHost, config.DatabasePort, config.DatabaseName)
	client, _ := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(config.DatabaseName)

	if err != nil {
		return err
	}

	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
