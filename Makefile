GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GORUN := $(GOCMD) run

build:
	$(GOBUILD) -o bin/MovieRental

run:
	$(GORUN) cmd/main.go

test:
	$(GOTEST) ./tests/...

clean:
	$(GOCLEAN)
	rm -f bin/MovieRental