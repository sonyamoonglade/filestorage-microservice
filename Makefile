build:
	go build -o ./bin/ cmd/main.go && cp .env ./bin/
run:
	cd bin && main.exe