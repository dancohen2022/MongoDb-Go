package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver"
	"github.com/mongodb/mongo-go-driver/mongo/primitive"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	FirstName string             `json:"firstname, omitempty" bson:"firstname, omitempty"`
	LastName  string             `json:"lasttname, omitempty" bson:"lasttname, omitempty"`
}

var client *mongo.client

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	http.ListenAndServe(":12345", router)
}
