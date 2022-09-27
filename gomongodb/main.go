package main

import (
	"fmt"
	"context"

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

}