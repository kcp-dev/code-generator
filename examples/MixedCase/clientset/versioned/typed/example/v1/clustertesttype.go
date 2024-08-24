/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	examplev1 "k8s.io/code-generator/examples/MixedCase/apis/example/v1"
)

// ClusterTestTypesClusterGetter has a method to return a ClusterTestTypeClusterInterface.
// A group's client should implement this interface.
type ClusterTestTypesClusterGetter interface {
	ClusterTestTypes() ClusterTestTypeClusterInterface
}

// ClusterTestTypeClusterInterface has methods to work with ClusterTestType resources.
type ClusterTestTypeClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*examplev1.ClusterTestTypeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) upstreamNodeMagic

	ClusterTestTypeExpansion
}

type clusterTestTypesClusterInterface struct {
	clientCache kcpclient.Cache[*ExampleV1Client]
}
