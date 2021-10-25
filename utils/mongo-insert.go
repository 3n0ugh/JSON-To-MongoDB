package utils

import (
	"context"
	"fmt"

	"github.com/3n0ugh/JSON-to-MongoDB/config"
	"github.com/3n0ugh/JSON-to-MongoDB/database"
	"github.com/3n0ugh/JSON-to-MongoDB/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDataToMongo() error {
	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	collection := database.Mg.Db.Collection(config.DatabaseColName)
	ctx := context.Background()

	for i := 0; i < len(BookList.Books); i++ {

		insertRes, err := collection.InsertOne(ctx, BookList.Books[i])
		if err != nil {
			fmt.Println(err, BookList.Books[i].Title)
		}
		filter := bson.D{{Key: "_id", Value: insertRes.InsertedID}}
		createdRecord := collection.FindOne(ctx, filter)
		createdBook := &models.Book{}
		createdRecord.Decode(createdBook)
		fmt.Printf("%s added successfully.\n", createdBook.Title)
	}

	fmt.Printf("\n\n================================\nTotal %d book added successfully.\n================================\n", len(BookList.Books))
	return nil
}
