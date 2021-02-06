![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sahindagdelen/golangtodo?style=for-the-badge)

![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/sahindagdelen/golangtodo?style=for-the-badge)

![Docker Pulls](https://img.shields.io/docker/pulls/sahindagdelen/golangtodo?style=for-the-badge)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sahindagdelen/go-todo?style=for-the-badge)

![Go](https://github.com/sahindagdelen/go-todo/workflows/Go/badge.svg?branch=master)

## General info

Graphql api for crud operations of todo app using golang, mongodb atlas as db.

## Installation

## Docker Instructions

docker pull sahindagdelen/golangtodo

MongoDb Atlas configuration :
Mongodb connection string must be modified with username and password. <admin>, <password>
Dns errors might occur when running application, cd /etc sudo resolv.conf modify your dns to 8.8.8.8 You should also add
your ip to whitelist on mongodb atlas admin panel network access tab.

## Features

* Golang graphql api (CRUD operations)
* Graphiql implementation added. (http://localhost:8080/api/graphiql)
* Modularized golang app. (go.mod)
* Package structure edited complying with golang package structure standarts.
* Docker image size kept at minimum.
* Multistage docker file.
* Mongo DB Atlas used as database.
* Properties read from config file.

To-do list:

* Clean code refactor.

## Screenshots

![Screenshot](/examples/screenshots/getAllTasks.png?raw=true "Get all tasks")
![Screenshot](/examples/screenshots/getOneTask.png?raw=true "Get task" )
![Screenshot](/examples/screenshots/createTask.png?raw=true "Create Task")
![Screenshot](/examples/screenshots/updateTaskStatus.png?raw=true "Update task status")
![Screenshot](/examples/screenshots/deleteTask.png?raw=true "Delete task")
![Screenshot](/examples/screenshots/deleteAllTasks.png?raw=true "Delete all tasks")

## Status

Project is: _in progress_ , more features will be developed and integrated into the repo.

## Inspiration

https://morioh.com/p/82b11315afa1

https://github.com/graphql-go/graphql/blob/master/examples/crud/main.go

https://www.bradcypert.com/using-mongos-objectids-with-go-graphql/

https://github.com/graphql-go/graphql/blob/62a7bb0a9839309631e20fc4ed032cbf00f88544/examples/http-post/main.go#L20

https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

https://levelup.gitconnected.com/graphql-with-go-simple-server-tutorial-8678dbba20b9

https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c

https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742

https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2

https://github.com/motty93/Golang/tree/424e25c26b2313ef991552b9d978426b7caa417b/youtube/mongo_unit_test

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Contact

Created by [@sahindagdelen](https://twitter.com/sdgdln) - feel free to contact me!

## License

[MIT](https://choosealicense.com/licenses/mit/)
