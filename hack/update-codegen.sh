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
  --input secondexample/v1 \
  --input existinginterfaces/v1 \
  --output-base . \
  --output-package acme.corp/pkg/generated/clientset \
  --trim-path-prefix acme.corp

${KUBE_APPLYCONFIGURATION_GEN} \
  --go-header-file ./../hack/boilerplate/boilerplate.generatego.txt \
  --input-dirs acme.corp/pkg/apis/example/v1,acme.corp/pkg/apis/example/v1alpha1,acme.corp/pkg/apis/example/v1beta1,acme.corp/pkg/apis/example/v2,acme.corp/pkg/apis/example3/v1,acme.corp/pkg/apis/secondexample/v1,acme.corp/pkg/apis/existinginterfaces/v1 \
  --output-base . \
  --output-package acme.corp/pkg/generated/applyconfigurations \
  --trim-path-prefix acme.corp

${KUBE_LISTER_GEN} \
  --go-header-file ./../hack/boilerplate/boilerplate.generatego.txt \
  --input-dirs acme.corp/pkg/apis/example/v1,acme.corp/pkg/apis/example/v1alpha1,acme.corp/pkg/apis/example/v1beta1,acme.corp/pkg/apis/example/v2,acme.corp/pkg/apis/example3/v1,acme.corp/pkg/apis/secondexample/v1,acme.corp/pkg/apis/existinginterfaces/v1 \
  --output-base . \
  --output-package acme.corp/pkg/generated/listers \
  --trim-path-prefix acme.corp

${KUBE_INFORMER_GEN} \
  --versioned-clientset-package acme.corp/pkg/generated/clientset/versioned \
  --listers-package acme.corp/pkg/generated/listers \
  --go-header-file ./../hack/boilerplate/boilerplate.generatego.txt \
  --input-dirs acme.corp/pkg/apis/example/v1,acme.corp/pkg/apis/example/v1alpha1,acme.corp/pkg/apis/example/v1beta1,acme.corp/pkg/apis/example/v2,acme.corp/pkg/apis/example3/v1,acme.corp/pkg/apis/secondexample/v1,acme.corp/pkg/apis/existinginterfaces/v1 \
  --output-base . \
  --output-package acme.corp/pkg/generated/informers \
  --trim-path-prefix acme.corp

# Generate cluster-aware clients, informers and listers using generated single-cluster code
./../bin/code-generator \
    "client:standalone=true,outputPackagePath=acme.corp/pkg/kcpexisting/clients,apiPackagePath=acme.corp/pkg/apis,singleClusterClientPackagePath=acme.corp/pkg/generated/clientset/versioned,singleClusterApplyConfigurationsPackagePath=acme.corp/pkg/generated/applyconfigurations,headerFile=./../hack/boilerplate/boilerplate.go.txt" \
    "lister:apiPackagePath=acme.corp/pkg/apis,singleClusterListerPackagePath=acme.corp/pkg/generated/listers,headerFile=./../hack/boilerplate/boilerplate.go.txt" \
    "informer:standalone=true,outputPackagePath=acme.corp/pkg/kcpexisting/clients,apiPackagePath=acme.corp/pkg/apis,singleClusterClientPackagePath=acme.corp/pkg/generated/clientset/versioned,singleClusterListerPackagePath=acme.corp/pkg/generated/listers,singleClusterInformerPackagePath=acme.corp/pkg/generated/informers/externalversions,headerFile=./../hack/boilerplate/boilerplate.go.txt" \
    "paths=./pkg/apis/..." \
    "output:dir=./pkg/kcpexisting/clients"

# Generate cluster-aware clients, informers and listers assuming no single-cluster listers or informers
./../bin/code-generator \
  "client:standalone=true,outputPackagePath=acme.corp/pkg/kcp/clients,apiPackagePath=acme.corp/pkg/apis,singleClusterClientPackagePath=acme.corp/pkg/generated/clientset/versioned,singleClusterApplyConfigurationsPackagePath=acme.corp/pkg/generated/applyconfigurations,headerFile=./../hack/boilerplate/boilerplate.go.txt" \
  "lister:apiPackagePath=acme.corp/pkg/apis,headerFile=./../hack/boilerplate/boilerplate.go.txt" \
  "informer:standalone=true,outputPackagePath=acme.corp/pkg/kcp/clients,apiPackagePath=acme.corp/pkg/apis,singleClusterClientPackagePath=acme.corp/pkg/generated/clientset/versioned,headerFile=./../hack/boilerplate/boilerplate.go.txt" \
  "paths=./pkg/apis/..." \
  "output:dir=./pkg/kcp/clients"

popd