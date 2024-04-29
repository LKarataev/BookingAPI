all: run

run:
	go run cmd/app/main.go 

test:
	go test -cover ./internal/handlers
#	go test ./internal/handlers -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html