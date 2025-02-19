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

CONTROLLER_GEN_VER := v0.16.5
CONTROLLER_GEN_BIN := controller-gen
CONTROLLER_GEN := $(GOBIN_DIR)/$(CONTROLLER_GEN_BIN)-$(CONTROLLER_GEN_VER)
export CONTROLLER_GEN

GOLANGCI_LINT_VER := v1.62.2
GOLANGCI_LINT_BIN := golangci-lint
GOLANGCI_LINT := $(GOBIN_DIR)/$(GOLANGCI_LINT_BIN)-$(GOLANGCI_LINT_VER)

KUBE_CLIENT_GEN_VER := v0.32.0
KUBE_CLIENT_GEN_BIN := client-gen
KUBE_LISTER_GEN_VER := v0.32.0
KUBE_LISTER_GEN_BIN := lister-gen
KUBE_INFORMER_GEN_VER := v0.32.0
KUBE_INFORMER_GEN_BIN := informer-gen
KUBE_APPLYCONFIGURATION_GEN_VER := v0.32.0
KUBE_APPLYCONFIGURATION_GEN_BIN := applyconfiguration-gen

KUBE_CLIENT_GEN := $(GOBIN_DIR)/$(KUBE_CLIENT_GEN_BIN)-$(KUBE_CLIENT_GEN_VER)
export KUBE_CLIENT_GEN
KUBE_LISTER_GEN := $(GOBIN_DIR)/$(KUBE_LISTER_GEN_BIN)-$(KUBE_LISTER_GEN_VER)
export KUBE_LISTER_GEN
KUBE_INFORMER_GEN := $(GOBIN_DIR)/$(KUBE_INFORMER_GEN_BIN)-$(KUBE_INFORMER_GEN_VER)
export KUBE_INFORMER_GEN
KUBE_APPLYCONFIGURATION_GEN := $(GOBIN_DIR)/$(KUBE_APPLYCONFIGURATION_GEN_BIN)-$(KUBE_APPLYCONFIGURATION_GEN_VER)
export KUBE_APPLYCONFIGURATION_GEN

OPENSHIFT_GOIMPORTS_VER := c70783e636f2213cac683f6865d88c5edace3157
OPENSHIFT_GOIMPORTS_BIN := openshift-goimports
OPENSHIFT_GOIMPORTS := $(TOOLS_DIR)/$(OPENSHIFT_GOIMPORTS_BIN)-$(OPENSHIFT_GOIMPORTS_VER)
export OPENSHIFT_GOIMPORTS # so hack scripts can use it

$(OPENSHIFT_GOIMPORTS):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) github.com/openshift-eng/openshift-goimports $(OPENSHIFT_GOIMPORTS_BIN) $(OPENSHIFT_GOIMPORTS_VER)

imports: $(OPENSHIFT_GOIMPORTS)
	$(OPENSHIFT_GOIMPORTS) -m github.com/kcp-dev/code-generator
	$(OPENSHIFT_GOIMPORTS) --path ./examples -m acme.corp
.PHONY: imports

$(CONTROLLER_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) sigs.k8s.io/controller-tools/cmd/$(CONTROLLER_GEN_BIN) $(CONTROLLER_GEN_BIN) $(CONTROLLER_GEN_VER)

$(KUBE_CLIENT_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) k8s.io/code-generator/cmd/$(KUBE_CLIENT_GEN_BIN) $(KUBE_CLIENT_GEN_BIN) $(KUBE_CLIENT_GEN_VER)
$(KUBE_LISTER_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) k8s.io/code-generator/cmd/$(KUBE_LISTER_GEN_BIN) $(KUBE_LISTER_GEN_BIN) $(KUBE_LISTER_GEN_VER)
$(KUBE_INFORMER_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) k8s.io/code-generator/cmd/$(KUBE_INFORMER_GEN_BIN) $(KUBE_INFORMER_GEN_BIN) $(KUBE_INFORMER_GEN_VER)
$(KUBE_APPLYCONFIGURATION_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) k8s.io/code-generator/cmd/$(KUBE_APPLYCONFIGURATION_GEN_BIN) $(KUBE_APPLYCONFIGURATION_GEN_BIN) $(KUBE_APPLYCONFIGURATION_GEN_VER)


.PHONY: build
build: ## Build the project
	mkdir -p bin
	go build -o bin

.PHONY: install
install:
	go install

.PHONY: codegen
codegen: $(CONTROLLER_GEN) $(KUBE_CLIENT_GEN) $(KUBE_LISTER_GEN) $(KUBE_INFORMER_GEN) $(KUBE_APPLYCONFIGURATION_GEN) build
	./hack/update-codegen.sh
	$(MAKE) imports

$(GOLANGCI_LINT):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint $(GOLANGCI_LINT_BIN) $(GOLANGCI_LINT_VER)

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --timeout=10m ./...
	cd examples && $(GOLANGCI_LINT) run --timeout=10m ./...

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

$(TOOLS_DIR)/verify_boilerplate.py:
	mkdir -p $(TOOLS_DIR)
	curl --fail --retry 3 -L -o $(TOOLS_DIR)/verify_boilerplate.py https://raw.githubusercontent.com/kubernetes/repo-infra/master/hack/verify_boilerplate.py
	chmod +x $(TOOLS_DIR)/verify_boilerplate.py

.PHONY: verify-boilerplate
verify-boilerplate: $(TOOLS_DIR)/verify_boilerplate.py
	$(TOOLS_DIR)/verify_boilerplate.py --boilerplate-dir=hack/boilerplate
