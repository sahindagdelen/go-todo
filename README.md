![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sahindagdelen/golangtodo?style=for-the-badge)

![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/sahindagdelen/golangtodo?style=for-the-badge)

![Docker Pulls](https://img.shields.io/docker/pulls/sahindagdelen/golangtodo?style=for-the-badge)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sahindagdelen/go-todo?style=for-the-badge)

## General info

Graphql api for crud operations of todo app using golang, mongodb atlas as db.

## Installation

## Docker Run Steps

* BUILD

docker build -t *yourdockerhubusername* / *imagenameyouchoose* .

* RUN

docker run -p 8080:8080 *dockerhubusername* / *imagenameyouchoose*

Application will be on http://localhost:8080/api/graphql

## Docker Pull

docker pull sahindagdelen/golangtodo

## Features

* Golang restful api (CRUD operations)
* Modularized golang app. (go.mod)
* Graphql implementation.
* Docker image size kept at minimum.
* Mongo DB Atlas used as database.

To-do list:

* Swagger implementation
* Unit tests
* Properties file.
* Clean code refactor.
* Sample requests.

## Status

Project is: _in progress_ , more features will be developed and integrated into the repo.

## Inspiration

https://morioh.com/p/82b11315afa1

https://github.com/graphql-go/graphql/blob/master/examples/crud/main.go

https://www.bradcypert.com/using-mongos-objectids-with-go-graphql/

https://github.com/graphql-go/graphql/blob/62a7bb0a9839309631e20fc4ed032cbf00f88544/examples/http-post/main.go#L20

https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

https://levelup.gitconnected.com/graphql-with-go-simple-server-tutorial-8678dbba20b9

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Contact

Created by [@sahindagdelen](https://twitter.com/sdgdln) - feel free to contact me!

## License

[MIT](https://choosealicense.com/licenses/mit/)
