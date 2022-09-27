package main

import (
	"fmt"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Book struct{
	Title string
	Author string
	Pages int
}

func main(){
	fmt.Println("Books App - Mongodb")

	// Set mongo db options
	options := options.Client().ApplyURI("mongodb+srv://kimadmin:16577561KIMsam@cluster1.lt4izot.mongodb.net/?retryWrites=true&w=majority")

	//Connect 
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		log.Fatal(err)
	}
}