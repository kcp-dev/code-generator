clean:
	rm -rf bin/

build: clean
	go build -o bin/ ./cmd/*

codegen:
	cd examples && ./hack/update-codegen.sh

codegen-dev: build
	cd examples && ./hack/update-codegen.sh
