# Copyright 2022 The KCP Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := /usr/bin/env bash

GO_INSTALL = ./hack/go-install.sh

TOOLS_DIR=hack/tools
GOBIN_DIR := $(abspath $(TOOLS_DIR))
TMPDIR := $(shell mktemp -d)

CONTROLLER_GEN_VER := v0.8.0
CONTROLLER_GEN_BIN := controller-gen
CONTROLLER_GEN := $(TOOLS_DIR)/$(CONTROLLER_GEN_BIN)-$(CONTROLLER_GEN_VER)

GOLANGCI_LINT_VER := v1.44.2
GOLANGCI_LINT_BIN := golangci-lint
GOLANGCI_LINT := $(GOBIN_DIR)/$(GOLANGCI_LINT_BIN)-$(GOLANGCI_LINT_VER)

KUBE_CLIENT_GEN_VER := v0.24.0
KUBE_CLIENT_GEN_BIN := client-gen
KUBE_CLIENT_GEN := $(GOBIN_DIR)/$(KUBE_CLIENT_GEN_BIN)-$(KUBE_CLIENT_GEN_VER)

$(KUBE_CLIENT_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) k8s.io/code-generator/cmd/client-gen $(KUBE_CLIENT_GEN_BIN) $(KUBE_CLIENT_GEN_VER)

$(CONTROLLER_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) sigs.k8s.io/controller-tools/cmd/controller-gen $(CONTROLLER_GEN_BIN) $(CONTROLLER_GEN_VER)

.PHONY: build
build: ## Build the project
	mkdir -p bin
	go build -o bin

.PHONY: install
install:
	go install

.PHONY: codegen
codegen: $(CONTROLLER_GEN) $(KUBE_CLIENT_GEN) build
	# Generate deepcopy functions
	${CONTROLLER_GEN} object paths=./examples/pkg/apis/...

	# Generate standard clientset
	$(KUBE_CLIENT_GEN) \
		--clientset-name versioned \
		--go-header-file hack/boilerplate/boilerplate.generatego.txt \
		--input-base github.com/kcp-dev/code-generator/examples/pkg/apis \
		--input example/v1 \
		--output-base . \
		--output-package github.com/kcp-dev/code-generator/examples/pkg/generated/clientset \
		--trim-path-prefix github.com/kcp-dev/code-generator

	# Generate cluster clientset and listers
	bin/code-generator \
		client,lister,informer \
		--clientset-name clusterclient \
		--go-header-file hack/boilerplate/boilerplate.generatego.txt \
		--clientset-api-path github.com/kcp-dev/code-generator/examples/pkg/generated/clientset/versioned \
		--input-dir ./examples/pkg/apis \
		--output-dir ./examples/pkg \
		--group-versions example:v1

$(GOLANGCI_LINT):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint $(GOLANGCI_LINT_BIN) $(GOLANGCI_LINT_VER)

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --timeout=10m ./...

.PHONY: test
test:
	go test ./...

# Note, running this locally if you have any modified files, even those that are not generated,
# will result in an error. This target is mostly for CI jobs.
.PHONY: verify-codegen
verify-codegen:
	if [[ -n "${GITHUB_WORKSPACE}" ]]; then \
		mkdir -p $$(go env GOPATH)/src/github.com/kcp-dev; \
		ln -s ${GITHUB_WORKSPACE} $$(go env GOPATH)/src/github.com/kcp-dev/code-generator; \
	fi

	$(MAKE) codegen

	if ! git diff --quiet HEAD; then \
		git diff; \
		echo "You need to run 'make codegen' to update generated files and commit them"; \
		exit 1; \
	fi
