all: build

install:
	go install

build:
	go build

test:
	go test -v

clean:
	rm -f tftrigger

.PHONY: install build test clean