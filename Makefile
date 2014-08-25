
all: install test

install:
	cd databench_go; go install

test:
	cd databench_go; go test
