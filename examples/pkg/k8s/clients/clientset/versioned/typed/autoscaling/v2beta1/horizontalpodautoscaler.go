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

package v2beta1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	autoscalingv2beta1 "acme.corp/pkg/apis/autoscaling/v2beta1"
	autoscalingv2beta1client "acme.corp/pkg/generated/clientset/versioned/typed/autoscaling/v2beta1"
)

// HorizontalPodAutoscalersClusterGetter has a method to return a HorizontalPodAutoscalerClusterInterface.
// A group's cluster client should implement this interface.
type HorizontalPodAutoscalersClusterGetter interface {
	HorizontalPodAutoscalers() HorizontalPodAutoscalerClusterInterface
}

// HorizontalPodAutoscalerClusterInterface can operate on HorizontalPodAutoscalers across all clusters,
// or scope down to one cluster and return a HorizontalPodAutoscalersNamespacer.
type HorizontalPodAutoscalerClusterInterface interface {
	Cluster(logicalcluster.Path) HorizontalPodAutoscalersNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*autoscalingv2beta1.HorizontalPodAutoscalerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type horizontalPodAutoscalersClusterInterface struct {
	clientCache kcpclient.Cache[*autoscalingv2beta1client.AutoscalingV2beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *horizontalPodAutoscalersClusterInterface) Cluster(clusterPath logicalcluster.Path) HorizontalPodAutoscalersNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &horizontalPodAutoscalersNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all HorizontalPodAutoscalers across all clusters.
func (c *horizontalPodAutoscalersClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*autoscalingv2beta1.HorizontalPodAutoscalerList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).HorizontalPodAutoscalers(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all HorizontalPodAutoscalers across all clusters.
func (c *horizontalPodAutoscalersClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).HorizontalPodAutoscalers(metav1.NamespaceAll).Watch(ctx, opts)
}

// HorizontalPodAutoscalersNamespacer can scope to objects within a namespace, returning a autoscalingv2beta1client.HorizontalPodAutoscalerInterface.
type HorizontalPodAutoscalersNamespacer interface {
	Namespace(string) autoscalingv2beta1client.HorizontalPodAutoscalerInterface
}

type horizontalPodAutoscalersNamespacer struct {
	clientCache kcpclient.Cache[*autoscalingv2beta1client.AutoscalingV2beta1Client]
	clusterPath logicalcluster.Path
}

func (n *horizontalPodAutoscalersNamespacer) Namespace(namespace string) autoscalingv2beta1client.HorizontalPodAutoscalerInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).HorizontalPodAutoscalers(namespace)
}
