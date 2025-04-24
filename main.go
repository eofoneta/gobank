package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Todo struct {
	ID        int    `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error fetching env file")
	}

	MONGODB_URL := os.Getenv("MONGODB_URI")
	clientOption := options.Client().ApplyURI(MONGODB_URL)

	client, err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB ATLAS")

	collection = client.Database("goBank").Collection("todos")

	// app := fiber.New()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8000"
	}

	// app.Get("/api/todos", getTodos)
	// app.Post("/api/todos", createTodos)
	// app.Patch("/api/todos/:id", updateTodos)
	// app.Delete("/api/todos/:id", deleteTodos)
}

// func getTodos (c *fiber.Ctx) error {
// 	var todos []Todo;


// }