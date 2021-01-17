![Docker Build Status](https://img.shields.io/docker/build/sahindagdelen/golangtodo?style=for-the-badge)    

![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/sahindagdelen/golangtodo?style=for-the-badge)

![Docker Pulls](https://img.shields.io/docker/pulls/sahindagdelen/golangtodo)

# Go Rest API

Build on top of this sample go rest api https://morioh.com/p/82b11315afa1.

## Installation

## DOCKER RUN STEPS

1) BUILD

docker build -t sahindagdelen/golangtodo .


2) RUN

docker run -p 8080:8080 sahindagdelen/golangtodo

Application will be on http://localhost:8080/api/graphql   
 

graphql implementation is added.
https://levelup.gitconnected.com/graphql-with-go-simple-server-tutorial-8678dbba20b9

Swagger implementation will be added.
Example requests will be added.

## References

https://github.com/graphql-go/graphql/blob/master/examples/crud/main.go

https://www.bradcypert.com/using-mongos-objectids-with-go-graphql/

post with graphql
https://github.com/graphql-go/graphql/blob/62a7bb0a9839309631e20fc4ed032cbf00f88544/examples/http-post/main.go#L20


https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
