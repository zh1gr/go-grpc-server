go_mod:
	go mod tidy && go mod vendor && go mod verify

.PHONY: run
run:
	go run server.go