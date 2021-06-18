.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: run
run:
	go run ./main.go