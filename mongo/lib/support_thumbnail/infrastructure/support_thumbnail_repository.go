package infrastructure

import (
	domain "../domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type SupportThumbnailRepository struct {
}

func (repo SupportThumbnailRepository) GetSupportThumbnails() []domain.SupportThumbnail {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("content-mock").Collection("product_attribute_values")
	cur, errColl := collection.Find(ctx, bson.M{})
	if errColl != nil {
		log.Fatal(errColl)
	}
	defer cur.Close(context.Background())

	var results []domain.SupportThumbnail
	errAll := cur.All(context.Background(), &results)
	if errAll != nil {
		panic(errAll)
	}
	return results
}
