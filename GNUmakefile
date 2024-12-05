DAY ?= one

default: test 

solve:
	go run ./$(DAY)/main.go ./$(DAY)/solver.go

test:
	go test ./...

.PHONY: \
	solve \
	test