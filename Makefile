.PHONY: vendor

builddir = build

testflags = -cover -coverprofile=${builddir}/coverage.out -mod=vendor 

run:
	go run main.go

clean:
	rm -rf ${builddir}

prepare:
	mkdir -p ${builddir}

test: clean prepare
	go test ${testflags} ./...

cover: test
	go tool cover -html=${builddir}/coverage.out -o ${builddir}/coverage.html

vendor:
	go get
	go mod vendor
