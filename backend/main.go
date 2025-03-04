package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Task struct
type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     string             `json:"due_date"`
	Status      string             `json:"status"`
}

var client *mongo.Client

// Connect to MongoDB
func ConnectDB() *mongo.Collection {
	return client.Database("taskdb").Collection("tasks")
}

// sahilghadiya7331
// e6WZozLiLY1vPt8q
func initMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://parth:1205@task.rv0kd.mongodb.net/?retryWrites=true&w=majority&appName=TASK")
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

// Create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid task data", http.StatusBadRequest)
		log.Println("Failed to decode task:", err)
		return
	}

	task.ID = primitive.NewObjectID()

	collection := ConnectDB()
	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		http.Error(w, "Failed to insert task", http.StatusInternalServerError)
		log.Println("MongoDB Insert error:", err)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// Get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tasks []Task
	collection := ConnectDB()

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch", http.StatusInternalServerError)
		log.Println("MongoDB Find error:", err)
		return
	}

	for cur.Next(context.TODO()) {
		var task Task
		err := cur.Decode(&task)
		if err != nil {
			log.Println("Error decoding:", err)
		}
		tasks = append(tasks, task)
	}

	json.NewEncoder(w).Encode(tasks)
}

// Update a task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		log.Println("Failed to decode:", err)
		return
	}

	collection := ConnectDB()
	objID, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": task}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		log.Println("MongoDB Update error:", err)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// Delete a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	collection := ConnectDB()
	objID, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": objID}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		log.Println("MongoDB Delete error:", err)
		return
	}

	json.NewEncoder(w).Encode("Task Deleted...")
}

func main() {
	initMongoDB()
	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", GetTasks).Methods("GET")
	r.HandleFunc("/api/tasks", CreateTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", UpdateTask).Methods("PUT")
	r.HandleFunc("/api/tasks/{id}", DeleteTask).Methods("DELETE")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	fmt.Println("Server port : 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
