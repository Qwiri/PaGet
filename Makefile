BINARY_NAME = paget.out
MAIN=paget.go

all: build test

build:
	go build -o ${BINARY_NAME} ./cmd/paget/${MAIN}


clean:
	go clean
	rm ${BINARY_NAME}