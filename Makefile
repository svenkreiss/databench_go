
all: install test

install:
	cd databench; go install

test:
	cd databench; go test
