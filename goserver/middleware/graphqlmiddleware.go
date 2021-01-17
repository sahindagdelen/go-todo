package middleware

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/sahindagdelen/go-todo/goserver/models"
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

var TodoList []models.Todo

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
				/*
					curl --location --request POST 'http://localhost:8080/api/graphql' \
					--header 'Content-Type: application/json' \
					--data-raw '{ "query": "{ todo(id:\"5ff8b31b0f3e6c816ed77838\") { id status task } }" }'
				*/
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						return getOneTask(id), nil
					}
					return models.Todo{}, nil
				},
			},
			"todolist": &graphql.Field{
				Type:        graphql.NewList(todoType),
				Description: "List of todos",
				/*
						curl --location --request POST 'http://localhost:8080/api/graphql' \
					--header 'Content-Type: application/json' \
					--data-raw '{ "query": "{ todolist{ id status task } }" }'
				*/
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
		/*
				curl --location --request POST 'http://localhost:8080/api/graphql' \
			--header 'Content-Type: application/json' \
			--data-raw '{
			    "query": "mutation { createTask(task:\"Golang yaz13\")  }"
			}'
		*/

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
				newTask := models.Todo{
					ID:     primitive.NewObjectID(),
					Task:   task,
					Status: false,
				}
				return createOneTask(newTask), nil
			},
		},
		/*
			curl --location --request POST 'http://localhost:8080/api/graphql' \
			--header 'Content-Type: application/json' \
			--data-raw '{
			    "query": "mutation { updateTaskStatus(status:true , id : \"5ff8b31b0f3e6c816ed77838\"  )  }"
			}'
		*/
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
		/*
		   curl --location --request POST 'http://localhost:8080/api/graphql' \
		   --header 'Content-Type: application/json' \
		   --data-raw '{
		       "query": "mutation { deleteTask( id : \"5ffe0d1a45006ab82e126402\"  )  }"
		   }'
		*/
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
		/*
			 curl --location --request POST 'http://localhost:8080/api/graphql' \
			--header 'Content-Type: application/json' \
			--data-raw '{
			    "query": "mutation { deleteAllTasks }"
			}'
		*/

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

func executeQuery(postData models.PostData, schema graphql.Schema) *graphql.Result {
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
