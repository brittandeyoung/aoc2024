CHALLENGE ?= one

default: test 

solve:
	go run ./$(CHALLENGE)/main.go ./$(CHALLENGE)/solver.go

test:
	go test ./...

.PHONY: \
	solve \
	test