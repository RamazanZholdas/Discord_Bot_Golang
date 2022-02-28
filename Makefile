build:
	go build -o bin/main.exe

test:
	go test discordTestBot/tests -v	

run:
	go run main.go	