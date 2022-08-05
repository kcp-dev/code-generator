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
KUBE_LISTER_GEN_VER := v0.24.0
KUBE_LISTER_GEN_BIN := lister-gen
KUBE_INFORMER_GEN_VER := v0.24.0
KUBE_INFORMER_GEN_BIN := informer-gen

KUBE_CLIENT_GEN := $(GOBIN_DIR)/$(KUBE_CLIENT_GEN_BIN)-$(KUBE_CLIENT_GEN_VER)
KUBE_LISTER_GEN := $(GOBIN_DIR)/$(KUBE_LISTER_GEN_BIN)-$(KUBE_LISTER_GEN_VER)
KUBE_INFORMER_GEN := $(GOBIN_DIR)/$(KUBE_INFORMER_GEN_BIN)-$(KUBE_INFORMER_GEN_VER)


$(CONTROLLER_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) sigs.k8s.io/controller-tools/cmd/controller-gen $(CONTROLLER_GEN_BIN) $(CONTROLLER_GEN_VER)

$(KUBE_CLIENT_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) k8s.io/code-generator/cmd/client-gen $(KUBE_CLIENT_GEN_BIN) $(KUBE_CLIENT_GEN_VER)
$(KUBE_LISTER_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) k8s.io/code-generator/cmd/lister-gen $(KUBE_LISTER_GEN_BIN) $(KUBE_LISTER_GEN_VER)
$(KUBE_INFORMER_GEN):
	GOBIN=$(GOBIN_DIR) $(GO_INSTALL) k8s.io/code-generator/cmd/informer-gen $(KUBE_INFORMER_GEN_BIN) $(KUBE_INFORMER_GEN_VER)


.PHONY: build
build: ## Build the project
	mkdir -p bin
	go build -o bin

.PHONY: install
install:
	go install

.PHONY: codegen
codegen: $(CONTROLLER_GEN) $(KUBE_CLIENT_GEN) $(KUBE_LISTER_GEN) $(KUBE_INFORMER_GEN) build
	# Generate deepcopy functions
	${CONTROLLER_GEN} object paths=./examples/pkg/apis/...

	# Generate standard clientset
	$(KUBE_CLIENT_GEN) \
		--clientset-name versioned \
		--go-header-file hack/boilerplate/boilerplate.generatego.txt \
		--input-base github.com/kcp-dev/code-generator/examples/pkg/apis \
		--input example/v1 \
		--input example/v1alpha1 \
		--input example/v1beta1 \
		--input example/v2 \
		--input example3/v1 \
		--input secondexample/v1 \
		--input existinginterfaces/v1 \
		--output-base . \
		--output-package github.com/kcp-dev/code-generator/examples/pkg/generated/clientset \
		--trim-path-prefix github.com/kcp-dev/code-generator

	$(KUBE_LISTER_GEN) \
		--go-header-file hack/boilerplate/boilerplate.generatego.txt \
		--input-dirs github.com/kcp-dev/code-generator/examples/pkg/apis/existinginterfaces/v1 \
		--output-base . \
		--output-package github.com/kcp-dev/code-generator/examples/pkg/generated/listers \
		--trim-path-prefix github.com/kcp-dev/code-generator

	$(KUBE_INFORMER_GEN) \
		--versioned-clientset-package github.com/kcp-dev/code-generator/examples/pkg/generated/clientset/versioned \
		--listers-package github.com/kcp-dev/code-generator/examples/pkg/generated/listers \
		--go-header-file hack/boilerplate/boilerplate.generatego.txt \
		--input-dirs github.com/kcp-dev/code-generator/examples/pkg/apis/existinginterfaces/v1 \
		--output-base . \
		--output-package github.com/kcp-dev/code-generator/examples/pkg/generated/informers \
		--trim-path-prefix github.com/kcp-dev/code-generator

	# Generate cluster informers and listers
	bin/code-generator \
		lister,informer \
		--go-header-file hack/boilerplate/boilerplate.generatego.txt \
		--clientset-api-path github.com/kcp-dev/code-generator/examples/pkg/generated/clientset/versioned \
		--input-dir ./examples/pkg/apis \
		--output-dir ./examples/pkg \
		--group-versions example:v1 \
		--group-versions example:v2 \
		--group-versions example:v1alpha1 \
		--group-versions example:v1beta1 \
		--group-versions secondexample:v1 \
		--group-versions example3:v1

	# Generate cluster informers and listers that are compatible with upstream interfaces
	bin/code-generator \
		lister,informer \
		--go-header-file hack/boilerplate/boilerplate.generatego.txt \
		--clientset-api-path github.com/kcp-dev/code-generator/examples/pkg/generated/clientset/versioned \
		--listers-package github.com/kcp-dev/code-generator/examples/pkg/generated/listers \
		--informers-package github.com/kcp-dev/code-generator/examples/pkg/generated/informers/externalversions \
		--input-dir ./examples/pkg/apis \
		--output-dir ./examples/pkg/legacy \
		--group-versions existinginterfaces:v1

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

$(TOOLS_DIR)/verify_boilerplate.py:
	mkdir -p $(TOOLS_DIR)
	curl --fail --retry 3 -L -o $(TOOLS_DIR)/verify_boilerplate.py https://raw.githubusercontent.com/kubernetes/repo-infra/master/hack/verify_boilerplate.py
	chmod +x $(TOOLS_DIR)/verify_boilerplate.py

.PHONY: verify-boilerplate
verify-boilerplate: $(TOOLS_DIR)/verify_boilerplate.py
	$(TOOLS_DIR)/verify_boilerplate.py --boilerplate-dir=hack/boilerplate
