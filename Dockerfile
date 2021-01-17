# build stage
FROM golang as builder


WORKDIR gotodoapp/goserver

COPY go.mod .
COPY go.sum .

RUN go mod download


COPY  goserver  .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

# final stage
FROM scratch
COPY --from=builder . .
EXPOSE 8080
ENTRYPOINT ["/main"]