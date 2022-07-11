EXEC_NAME=git-frontend

all: build

build:
	go build -o $(EXEC_NAME) ./cmd/**/*.go
clean:
	rm -rf $(EXEC_NAME)