package server

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/sahindagdelen/go-todo/api/types/postdata"
	"github.com/sahindagdelen/go-todo/api/types/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var todoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: ObjectID,
			},
			"task": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var TodoList []todo.Todo

var ObjectID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "BSON",
	Description: "bson scalar type represents a BSON Object",
	//serialize bson.ObjectId to string
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case primitive.ObjectID:
			return value.Hex()
		case *primitive.ObjectID:
			v := value
			return v.Hex()
		default:
			return nil
		}
	},
	//ParseValue parses Graphql variable from string to bson.ObjectId
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			id, _ := primitive.ObjectIDFromHex(value)
			return id
		case *string:
			id, _ := primitive.ObjectIDFromHex(*value)
			return id
		default:
			return nil
		}
		return nil
	},


	//Parseliteral parses Graphql AST to bson.ObjectID
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			id, _ := primitive.ObjectIDFromHex(valueAST.Value)
			return id
		}
		return nil
	},
})

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"todo": &graphql.Field{
				Type:        todoType,
				Description: "Get todo by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						return getOneTask(id), nil
					}
					return todo.Todo{}, nil
				},
			},
			"todolist": &graphql.Field{
				Type:        graphql.NewList(todoType),
				Description: "List of todos",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return getAllTask(), nil
				},
			},
		},
	},
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createTask": &graphql.Field{
			Type:        graphql.String, //return type
			Description: "Create task",
			Args: graphql.FieldConfigArgument{
				"task": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				//marshall and cast the argument value
				task := params.Args["task"].(string)
				//perform mutation  create a new task save
				newTask := todo.Todo{
					ID:     primitive.NewObjectID(),
					Task:   task,
					Status: false,
				}
				return createOneTask(newTask), nil
			},
		},
		"updateTaskStatus": &graphql.Field{
			Type:        graphql.String,
			Description: "Update task",
			Args: graphql.FieldConfigArgument{
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				status := params.Args["status"].(bool)
				id := params.Args["id"].(string)
				return taskUpdateStatus(id, status), nil
			},
		},
		"deleteTask": &graphql.Field{
			Type:        graphql.String,
			Description: "Delete task",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(string)
				return deleteOneTask(id), nil
			},
		},
		"deleteAllTasks": &graphql.Field{
			Type:        graphql.String,
			Description: "Delete all tasks",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return deleteAllTasks(), nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

func executeQuery(postData postdata.PostData, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  postData.Query,
		VariableValues: postData.Variables,
		OperationName:  postData.Operation,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}
