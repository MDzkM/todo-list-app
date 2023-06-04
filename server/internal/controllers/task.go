package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/mdzkm/todo-list-app/server/internal/database"
	"github.com/mdzkm/todo-list-app/server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasks() []primitive.M {
	cur, err := database.GetCollection().Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func CreateTask(task models.Task) {
	insertResult, err := database.GetCollection().InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Task", insertResult.InsertedID)
}

func UpdateTask(task string, description string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"description": description}}
	result, err := database.GetCollection().UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

func CompleteTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := database.GetCollection().UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

// delete one task from the DB, delete by ID
func DeleteTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := database.GetCollection().DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Task", d.DeletedCount)
}
