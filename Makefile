build:
	go build -o bin/aoc2024 .

test:
	go test ./...

clean:
	rm -rf bin
