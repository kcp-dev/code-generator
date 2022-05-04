#!/bin/bash

go build -a -o client-gen main.go

echo "built binary"

cp client-gen /Users/vnarsing/go/src/github.com/varshaprasad96/custom-crd-operator

echo "copied binary"
