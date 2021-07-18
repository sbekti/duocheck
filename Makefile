build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o duocheck main.go

run:
	go run main.go

clean:
	rm -f duocheck

all: clean build
