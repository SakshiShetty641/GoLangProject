#Go Parameters

GOCMD :=go
GOBUILD :=$(GOCMD) build
GOCLEAN :=$(GOCMD) clean
GOTEST :=$(GOCMD) test
GORUN :=$(GOCMD) run


build:
    go build -o bin/MovieRental.go

run:
    go run cmd/main.go

test:
    go test ./tests/...

clean:
    rm -f app