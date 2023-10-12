package config

import (
	"context"
	"log"
	"time"
	"signup/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

)

func Connectdb() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	mongoConnection := options.Client().ApplyURI(constants.ConnectionString)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Fatal(err.Error())
		return nil,err
	}
	err=mongoClient.Ping(ctx,readpref.Primary())
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}

	return mongoClient, nil
}

func Getcollection(client *mongo.Client,dbname string,collname string)(*mongo.Collection){
	collection:=client.Database(dbname).Collection(collname)
	return collection
}
