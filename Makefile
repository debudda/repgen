build:
	go build -o repgen main.go

gen: build
	./repgen ./examples/example.go ./examples/example.rep.gen.go ./examples/example.gen.go ./examples/example.mock.gen.go

.PHONY:
	build