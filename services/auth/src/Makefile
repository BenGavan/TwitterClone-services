all: test build

test:
	go test

vet:
	go vet ./...

#fmt:
	#go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	#test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)
#
#lint:
#    go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

hello:
	echo "Hello"

build:
	go build -o main

run: build
	./main
