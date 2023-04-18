package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Animal interface {
	Eat()
	Run()
}

type Dog struct {
	Name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating\n", d.Name)
}

func (d *Dog) Run() {
	fmt.Printf("%s is running\n", d.Name)
}

func ShowEat(animal Animal) {
	animal.Eat()
}

func ShowRun(animal Animal) {
	animal.Run()
}

type Cat struct {
	Name string
}

func (c *Cat) Eat() {
	fmt.Printf("%s is eating\n", c.Name)
}

func (c *Cat) Run() {
	fmt.Printf("%s is running\n", c.Name)
}

func main() {

	// var dataSlic []int = foo()

	// var interfaceSlice []interface{} datadataSlic

	animals := [...]Animal{
		&Dog{Name: "Kenny"},
		&Cat{Name: "Nicole"},
	}

	for _, animal := range animals {
		fmt.Println(animal)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	usersCollection := client.Database("testing").Collection("users")

	// insert a single document into a collection
	// create a bson.D object
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	users := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}},
		bson.D{{"fullName", "User 3"}, {"age", 20}},
		bson.D{{"fullName", "User 4"}, {"age", 28}},
	}
	// insert the bson object slice using InsertMany()
	results, err := usersCollection.InsertMany(context.TODO(), users)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)

}
