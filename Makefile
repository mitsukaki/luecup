build:
	GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o bin/luecup src/*.go

init:
	go get ./...

dev:
	go build -o bin/luecup-dev src/*.go

run:
	./bin/luecup*

clean:
	rm -rf luecup
	rm -rf luecup-dev