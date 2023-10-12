package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"signup/interfaces"
	"signup/models"
)

type User struct {
	ctx        context.Context
	Collection *mongo.Collection
}

func Inituser(ctx context.Context, Collection *mongo.Collection) interfaces.Isignup {
	return &User{ctx, Collection}
}

func (user *User)SignUP(account *models.Signup) (*mongo.InsertOneResult, error) {
	result, err := user.Collection.InsertOne(user.ctx, account)
	if err != nil {
		return nil, err
	}
	return result, nil
}
