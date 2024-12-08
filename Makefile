.PHONY: test
.PHONY: all
.PHONY: clean
.PHONY: build
build:
	@go build main.go	

run:    build
	@./main	-m ./test.txt && rm ./main

pipe:    build
	@cat ./test.txt | ./main
