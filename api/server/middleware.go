package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sahindagdelen/go-todo/api/types/postdata"
	"github.com/sahindagdelen/go-todo/api/types/todo"
	"github.com/sahindagdelen/go-todo/config"
	"github.com/sahindagdelen/go-todo/test/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func prepareConnectionString() string {
	const connectionStringFormat = "mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority"
	return fmt.Sprintf(connectionStringFormat, config.Config.Database.Username, config.Config.Database.Password, config.Config.Database.Host, config.Config.Database.Name)
}

func getCollection() *mongo.Collection {
	connectionString := prepareConnectionString()
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
	collection := client.Database(config.Config.Database.Name).Collection(config.Config.Database.Collection)
	fmt.Println("Collection instance created!")
	return collection
}

func ExecuteQueryGraphql(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var postData postdata.PostData
	if err := json.NewDecoder(r.Body).Decode(&postData); err != nil {
		w.WriteHeader(400)
		return
	}
	payload := executeQuery(postData, schema)
	json.NewEncoder(w).Encode(payload)
}

//get all task from DB and return it
func getAllTask(collection types.CrudInterface) ([]todo.Todo, error) {
	var results []todo.Todo
	ctx := context.Background()
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return results, err
	}
	err = cur.All(ctx, &results)
	if err != nil {
		return results, err
	}
	return results, err
}

func getOneTask(collection types.CrudInterface, task string) (todo.Todo, error) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}

	cur := collection.FindOne(context.Background(), filter)
	if cur.Err() != nil {
		return todo.Todo{}, cur.Err()
	}
	var result todo.Todo
	e := cur.Decode(&result)
	if e != nil {
		return todo.Todo{}, e
	}
	return result, e
}

func createOneTask(collection types.CrudInterface, todo todo.Todo) (*mongo.InsertOneResult, error) {
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return insertResult, err
	}
	return insertResult, err
}

func updateTaskStatus(collection types.CrudInterface, task string, done bool) (*mongo.UpdateResult, error) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": done}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return result, err
	}
	fmt.Println("modified count : ", result.ModifiedCount)
	return result, err
}

func deleteOneTask(collection types.CrudInterface, task string) (*mongo.DeleteResult, error) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return d, err
	}
	fmt.Println("Deleted document", d.DeletedCount)
	return d, err
}

func deleteAllTasks(collection types.CrudInterface) (*mongo.DeleteResult, error) {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		return d, err
	}
	fmt.Println("Deleted document", d.DeletedCount)
	return d, err
}
