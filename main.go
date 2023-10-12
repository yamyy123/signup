package main

import (
	"context"
	"fmt"
	"log"
	"signup/config"
	"signup/service"
     "signup/routes"
	"signup/constants"
     "signup/controller"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx context.Context
	server *gin.Engine
)


func initAll(mongoClient *mongo.Client){
	ctx = context.TODO()
	Collection := config.Getcollection(mongoClient,constants.DatabaseName,constants.Collections)
	Service := service.Inituser(ctx, Collection)
	Controller := controller.Initsignup(Service)
	routes.Router(server,Controller)
}

func main(){
	server = gin.Default()
	mongoclient,err :=config.Connectdb()
	defer   mongoclient.Disconnect(ctx)
	if err!=nil{
		panic(err)
	}
	
	initAll(mongoclient)
	fmt.Println("server running on port",constants.Port)
	log.Fatal(server.Run(constants.Port))
}
