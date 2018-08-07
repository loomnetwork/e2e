all:
	go build -o e2e-test github.com/loomnetwork/e2e/cmd

deps:
	dep ensure --vendor-only 