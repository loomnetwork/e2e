all:
	go build -o e2e-test github.com/loomnetwork/e2e

deps:
	dep ensure --vendor-only 
