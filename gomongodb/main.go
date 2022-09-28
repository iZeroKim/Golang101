package main

import (
	"fmt"
	"context"
	"log"

	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
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

	context := context.TODO()
	//Connect 
	client, err := mongo.Connect(context, options)

	if err != nil {
		log.Fatal(err)
	}

	//check connection
	connection_err := client.Ping(context, nil)

    if connection_err != nil{
		log.Fatal(connection_err)
	} else {
		fmt.Println("Connected to MongoDB")
	}

	booksCollection := client.Database("books").Collection("bookList")
	defer booksCollection.Drop(context)

	// Insert one
	book1 := Book{Title: "Go 101", Author: "Dennis Maina", Pages: 314}
	insertRes, err := booksCollection.InsertOne(context, book1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Book %v successfully inserted!\n ", insertRes)




	//Fetch all
	cursor, err := booksCollection.Find(context, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context)
	for cursor.Next(context) {
		var book bson.M
		if err = cursor.Decode(&book); err!= nil {
			log.Fatal(err)
		}
		fmt.Println("Book \n\t Name:", book["title"], "\n\t Author:", book["author"] , "\n\t Pages:", book["pages" ])
	}
	
}