package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sahindagdelen/goserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const connectionString = "mongodb+srv://<admin>:<password>@cluster0.klm1m.mongodb.net/test?retryWrites=true&w=majority"

const dbName = "test"

const collectionName = "todolist"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance created!")
}

func ExecuteQueryGraphql(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var postData models.PostData
	if err := json.NewDecoder(r.Body).Decode(&postData); err != nil {
		w.WriteHeader(400)
		return
	}
	payload := executeQuery(postData, schema)
	json.NewEncoder(w).Encode(payload)
}

//get all task from DB and return it
func getAllTask() []models.Todo {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Todo
	for cur.Next(context.Background()) {
		var result models.Todo
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

func getOneTask(task string) models.Todo {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}

	cur := collection.FindOne(context.Background(), filter)
	if cur.Err() != nil {
		log.Fatal(cur.Err())
	}
	var result models.Todo
	e := cur.Decode(&result)
	if e != nil {
		log.Fatal(e)
	}
	return result
}

func createOneTask(task models.Todo) string {
	insertResult, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Created document id, %s ", insertResult.InsertedID.(primitive.ObjectID).String())
}

func taskUpdateStatus(task string, done bool) string {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": done}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count : ", result.ModifiedCount)
	return fmt.Sprintf("modified count :  %d", result.ModifiedCount)
}

func deleteOneTask(task string) string {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted document", d.DeletedCount)
	return fmt.Sprintf("Deleted %d document", d.DeletedCount)
}

func deleteAllTasks() string {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document %s", d.DeletedCount)
	return fmt.Sprintf("Deleted %d documents", d.DeletedCount)
}