BINARY=engine
test:
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} app/*.go

serve:
	nodemon --exec go run app/main.go --signal SIGTERM