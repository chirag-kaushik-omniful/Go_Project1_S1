package dbconn

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB_Instance *mongo.Client
var DB_Ctx context.Context

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	fmt.Println("Going...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err == nil {
		fmt.Println(Ping(client, ctx))
	}

	DB_Instance = client
	DB_Ctx = ctx

	return client, ctx, cancel, err
}

func Ping(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
