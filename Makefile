.PHONY: vendor

run:
	go run main.go

test:
	go test -v -mod=vendor ./...

vendor:
	go get
	go mod vendor
