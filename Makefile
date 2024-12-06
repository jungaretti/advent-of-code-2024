build:
	go build -o bin/bacon .

test:
	go test ./...

clean:
	rm -rf bin
