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
	fmt.Println("\n\nINSERT ONE OPERATION")
	book1 := Book{Title: "Go 101", Author: "Dennis Maina", Pages: 314}
	insertRes, err := booksCollection.InsertOne(context, book1)
	if err != nil && insertRes != nil {
		log.Fatal(err)
	}
	fmt.Printf("One book successfully inserted!\n ")
	//Fetch all
	fmt.Println("BookList after insert one:")
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

	//Insert many at a go
	fmt.Println("\n\nINSERT MANY OPERATION")
	book2 := Book{Title: "The Giver", Author: "Alex Allan", Pages: 314}
	book3 := Book{Title: "Goroutines", Author: "Sheila Sharon", Pages: 200}
	book4 := Book{Title: "The Alchemist", Author: "Dennis Maina", Pages: 394}
	book5 := Book{Title: "Kafka", Author: "Samuel Kimani", Pages: 520}

	manyBooks := []interface{}{book2, book3, book4, book5}
	insertMultipleRes, err := booksCollection.InsertMany(context, manyBooks)
	if err != nil && insertMultipleRes != nil {
		log.Fatal(err)
	}
	fmt.Printf("Multiple books successfully inserted!\n ")



	//Fetch all
	fmt.Println("BookList after insert many:")
	cursor1, err := booksCollection.Find(context, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor1.Close(context)
	for cursor1.Next(context) {
		var book bson.M
		if err = cursor1.Decode(&book); err!= nil {
			log.Fatal(err)
		}
		fmt.Println("Book \n\t Name:", book["title"], "\n\t Author:", book["author"] , "\n\t Pages:", book["pages" ])
	}

	//Update one document
	fmt.Println("\n\nUPDATE OPERATION")
	filter := bson.D{{"title", "Kafka"}}

	update := bson.D{
		{"$inc", bson.D{
				{"pages", 200},
		}},
	}
	
	updateRes, err := booksCollection.UpdateOne(context, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v documents matched, %v documents updated.\n ", updateRes.MatchedCount, updateRes.ModifiedCount)
	//Fetch all

	fmt.Println("BookList after update:")
	cursor2, err := booksCollection.Find(context, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor2.Close(context)
	for cursor2.Next(context) {
		var book bson.M
		if err = cursor2.Decode(&book); err!= nil {
			log.Fatal(err)
		}
		fmt.Println("Book \n\t Name:", book["title"], "\n\t Author:", book["author"] , "\n\t Pages:", book["pages" ])
	}
	

	//Delete document - using collection.DeleteOne() or collection.DeleteMany()
	fmt.Println("\n\nDELETE OPERATION")
	deleteFilter := bson.D{{"title", "The Alchemist"}}

	deleteOneRes, err := booksCollection.DeleteOne(context, deleteFilter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents\n", deleteOneRes.DeletedCount)

	//Fetch all

	fmt.Println("BookList after delete:")
	cursor3, err := booksCollection.Find(context, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor3.Close(context)
	for cursor3.Next(context) {
		var book bson.M
		if err = cursor3.Decode(&book); err!= nil {
			log.Fatal(err)
		}
		fmt.Println("Book \n\t Name:", book["title"], "\n\t Author:", book["author"] , "\n\t Pages:", book["pages" ])
	}
}