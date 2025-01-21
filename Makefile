.DEFAULT_GOAL := build

init:
	git submodule update --init --recursive

docker:
	docker build -t "ruby-hdl" .

docker-run:
	docker run -d --restart always -p 127.0.0.1:43828:8080 ruby-hdl

build:
	go build -v -o build/serv cmd/serv/main.go

build-debug:
	go build -v -gcflags="all=-N -l" -o build/serv cmd/serv/main.go

postbuild:
	mkdir -p build/tmp/task
	cp config.json build/

run:
	./build/serv

.PHONY: docker docker-run build build-debug run
