#!/usr/bin/env bash

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

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

if [[ -z "${MAKELEVEL:-}" ]]; then
  echo 'You must invoke this script via make'
  exit 1
fi

source cluster_codegen.sh

pushd ./examples

# Generate deepcopy functions
${CONTROLLER_GEN} object paths=./pkg/apis/...

# Generate standard clientset
${KUBE_CLIENT_GEN} \
  --clientset-name versioned \
  --go-header-file ./../hack/boilerplate/boilerplate.generatego.txt \
  --input-base acme.corp/pkg/apis \
  --input example/v1 \
  --input example/v1alpha1 \
  --input example/v1beta1 \
  --input example/v2 \
  --input example3/v1 \
  --input exampledashed/v1 \
  --input secondexample/v1 \
  --input existinginterfaces/v1 \
  --output-dir ./pkg/generated/clientset \
  --output-pkg acme.corp/pkg/generated/clientset

${KUBE_APPLYCONFIGURATION_GEN} \
  --go-header-file ./../hack/boilerplate/boilerplate.generatego.txt \
  --output-dir ./pkg/generated/applyconfigurations \
  --output-pkg acme.corp/pkg/generated/applyconfigurations \
  acme.corp/pkg/apis/example/v1 acme.corp/pkg/apis/example/v1alpha1 acme.corp/pkg/apis/example/v1beta1 acme.corp/pkg/apis/example/v2 acme.corp/pkg/apis/example3/v1 acme.corp/pkg/apis/exampledashed/v1 acme.corp/pkg/apis/secondexample/v1 acme.corp/pkg/apis/existinginterfaces/v1

${KUBE_LISTER_GEN} \
  --go-header-file ./../hack/boilerplate/boilerplate.generatego.txt \
  --output-dir ./pkg/generated/listers \
  --output-pkg acme.corp/pkg/generated/listers \
  acme.corp/pkg/apis/example/v1 acme.corp/pkg/apis/example/v1alpha1 acme.corp/pkg/apis/example/v1beta1 acme.corp/pkg/apis/example/v2 acme.corp/pkg/apis/example3/v1 acme.corp/pkg/apis/exampledashed/v1 acme.corp/pkg/apis/secondexample/v1 acme.corp/pkg/apis/existinginterfaces/v1

${KUBE_INFORMER_GEN} \
  --versioned-clientset-package acme.corp/pkg/generated/clientset/versioned \
  --listers-package acme.corp/pkg/generated/listers \
  --go-header-file ./../hack/boilerplate/boilerplate.generatego.txt \
  --output-dir ./pkg/generated/informers \
  --output-pkg acme.corp/pkg/generated/informers \
  acme.corp/pkg/apis/example/v1 acme.corp/pkg/apis/example/v1alpha1 acme.corp/pkg/apis/example/v1beta1 acme.corp/pkg/apis/example/v2 acme.corp/pkg/apis/example3/v1 acme.corp/pkg/apis/exampledashed/v1 acme.corp/pkg/apis/secondexample/v1 acme.corp/pkg/apis/existinginterfaces/v1

# Generate cluster-aware clients, informers and listers using generated single-cluster code
cluster::codegen::gen_client \
  --boilerplate ../hack/boilerplate/boilerplate.generatego.txt \
  --output-dir pkg/kcpexisting/clients \
  --output-pkg acme.corp/pkg/kcpexisting/clients \
  --versioned-clientset-dir pkg/kcpexisting/clients/clientset/versioned \
  --versioned-clientset-pkg acme.corp/pkg/kcpexisting/clients/clientset/versioned \
  --informers-dir pkg/kcpexisting/clients/informers/externalversions \
  --informers-pkg acme.corp/pkg/kcpexisting/clients/informers/externalversions \
  --with-watch \
  --single-cluster-versioned-clientset-pkg acme.corp/pkg/generated/clientset/versioned \
  --single-cluster-applyconfigurations-pkg acme.corp/pkg/generated/applyconfigurations \
  --single-cluster-listers-pkg acme.corp/pkg/generated/listers \
  --single-cluster-informers-pkg acme.corp/pkg/generated/informers/externalversions \
  pkg/apis

# Generate cluster-aware clients, informers and listers assuming no single-cluster listers or informers
cluster::codegen::gen_client \
  --boilerplate ../hack/boilerplate/boilerplate.generatego.txt \
  --output-dir pkg/kcp/clients \
  --output-pkg acme.corp/pkg/kcp/clients \
  --versioned-clientset-dir pkg/kcp/clients/clientset/versioned \
  --versioned-clientset-pkg acme.corp/pkg/kcp/clients/clientset/versioned \
  --informers-dir pkg/kcp/clients/informers/externalversions \
  --informers-pkg acme.corp/pkg/kcp/clients/informers/externalversions \
  --with-watch \
  --single-cluster-versioned-clientset-pkg acme.corp/pkg/generated/clientset/versioned \
  --single-cluster-applyconfigurations-pkg acme.corp/pkg/generated/applyconfigurations \
  pkg/apis

popd
