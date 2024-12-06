.PHONY: test
.PHONY: all
.PHONY: clean
.PHONY: build
build:
	@go build main.go	

run:    build
	@./main	-c test.txt && rm ./main
