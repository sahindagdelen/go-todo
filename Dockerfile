FROM golang:latest
# create a working directory
WORKDIR /go-todo/server

#get dependencies
RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v github.com/graphql-go/graphql
RUN go get -d -v go.mongodb.org/mongo-driver/bson

# add source code.
ADD server server
# run main.go
CMD ["go", "run", "server/main.go"]