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

vendor:
	go get
	go mod vendor
