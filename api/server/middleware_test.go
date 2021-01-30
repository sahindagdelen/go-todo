package server

import (
	"context"
	"github.com/sahindagdelen/go-todo/api/types/todo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

type mockCollection struct {
}

func (m *mockCollection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{}
	c.InsertedID = "123456"
	return c, nil
}

func (m *mockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	c := &mongo.UpdateResult{}
	c.ModifiedCount = 1
	return c, nil
}

func (m *mockCollection) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c := &mongo.DeleteResult{}
	c.DeletedCount = 1
	return c, nil
}

func (m *mockCollection) DeleteMany(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c := &mongo.DeleteResult{}
	c.DeletedCount = 3
	return c, nil
}

func (m *mockCollection) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	c := &mongo.SingleResult{}
	return c
}

func (m *mockCollection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (*mongo.Cursor, error) {
	c := &mongo.Cursor{}
	return c, nil
}

func Test_CreateOneTask(t *testing.T) {
	collection := &mockCollection{}
	response, error := createOneTask(collection, todo.Todo{primitive.NewObjectID(), "Sleep", false})
	assert.Nil(t, error)
	assert.NotNil(t, response)
	assert.IsType(t, &mongo.InsertOneResult{}, response)
	assert.Equal(t, "123456", response.InsertedID)
}

func Test_UpdateTaskStatus(t *testing.T) {
	collection := &mockCollection{}
	response, error := updateTaskStatus(collection, "insertedId", true)
	assert.Nil(t, error)
	assert.NotNil(t, response)
	assert.Equal(t, int64(1), response.ModifiedCount)
}

func Test_DeleteOneTask(t *testing.T) {
	collection := &mockCollection{}
	response, error := deleteOneTask(collection, "insertedId")
	assert.Nil(t, error)
	assert.NotNil(t, response)
	assert.Equal(t, int64(1), response.DeletedCount)
}

func Test_DeleteAllTasks(t *testing.T) {
	collection := &mockCollection{}
	response, error := deleteAllTasks(collection)
	assert.Nil(t, error)
	assert.NotNil(t, response)
	assert.Equal(t, int64(3), response.DeletedCount)
}

func Test_GetTask_NoDocumentsFound(t *testing.T) {
	collection := &mockCollection{}
	response, error := getOneTask(collection, "123456")
	assert.NotNil(t, response)
	assert.NotNil(t, error)
	assert.IsType(t, todo.Todo{}, response)
	assert.EqualError(t, error, "mongo: no documents in result")
}
