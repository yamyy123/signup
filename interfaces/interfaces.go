package interfaces

import (
	"signup/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Isignup interface{
	SignUP(*models.Signup)(*mongo.InsertOneResult,error)
}