//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	corev1 "acme.corp/pkg/apis/core/v1"
	corev1client "acme.corp/pkg/generated/clientset/versioned/typed/core/v1"
)

// NodesClusterGetter has a method to return a NodeClusterInterface.
// A group's cluster client should implement this interface.
type NodesClusterGetter interface {
	Nodes() NodeClusterInterface
}

// NodeClusterInterface can operate on Nodes across all clusters,
// or scope down to one cluster and return a corev1client.NodeInterface.
type NodeClusterInterface interface {
	Cluster(logicalcluster.Path) corev1client.NodeInterface
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.NodeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type nodesClusterInterface struct {
	clientCache kcpclient.Cache[*corev1client.CoreV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *nodesClusterInterface) Cluster(clusterPath logicalcluster.Path) corev1client.NodeInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).Nodes()
}

// List returns the entire collection of all Nodes across all clusters.
func (c *nodesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.NodeList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Nodes().List(ctx, opts)
}

// Watch begins to watch all Nodes across all clusters.
func (c *nodesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Nodes().Watch(ctx, opts)
}
