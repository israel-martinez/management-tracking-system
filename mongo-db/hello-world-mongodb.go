package main

import (
    "context"
    "fmt"
    "log"
    "reflect"

    //"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
    Name string
    Age  int
    City string
}

func startMongoConnection() (*mongo.Client, error) {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")

    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")
    return client, err
}

func closeMongoConnection(client *mongo.Client){
    err := client.Disconnect(context.TODO())

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")
}

func insertOneDocument(collection *mongo.Collection, trainer Trainer){
    insertResult, err := collection.InsertOne(context.TODO(), trainer)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
    fmt.Println(reflect.TypeOf(insertResult))
}

func insertManyDocument(collection *mongo.Collection, documents []interface{}){
    insertManyResult, err := collection.InsertMany(context.TODO(), documents)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func main() {
    // Rest of the code will go here
    fmt.Println("- Hello World -")
    fmt.Println("Mongo db connection with Go ...")

    ash := Trainer{"Ash", 10, "Pallet Town"}
    misty := Trainer{"Misty", 10, "Cerulean City"}
    brock := Trainer{"Brock", 15, "Pewter City"}
    trainers := []interface{}{ash, misty, brock}
    

    client , err := startMongoConnection()
    fmt.Println(reflect.TypeOf(client))
    fmt.Println(reflect.TypeOf(err))

    if err != nil {
        log.Fatal(err)
    } else {
        collection := client.Database("test").Collection("trainers")
        fmt.Println(reflect.TypeOf(collection))
        fmt.Println(reflect.TypeOf(ash))
        fmt.Println(reflect.TypeOf(misty))
        fmt.Println(reflect.TypeOf(brock))
        fmt.Println(reflect.TypeOf(trainers))

        insertOneDocument(collection, ash)
        insertManyDocument(collection, trainers)

        closeMongoConnection(client)
    }
    fmt.Println("End Main ..")
}