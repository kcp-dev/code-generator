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

	policyv1 "acme.corp/pkg/apis/policy/v1"
	policyv1client "acme.corp/pkg/generated/clientset/versioned/typed/policy/v1"
)

// PodDisruptionBudgetsClusterGetter has a method to return a PodDisruptionBudgetClusterInterface.
// A group's cluster client should implement this interface.
type PodDisruptionBudgetsClusterGetter interface {
	PodDisruptionBudgets() PodDisruptionBudgetClusterInterface
}

// PodDisruptionBudgetClusterInterface can operate on PodDisruptionBudgets across all clusters,
// or scope down to one cluster and return a PodDisruptionBudgetsNamespacer.
type PodDisruptionBudgetClusterInterface interface {
	Cluster(logicalcluster.Path) PodDisruptionBudgetsNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*policyv1.PodDisruptionBudgetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type podDisruptionBudgetsClusterInterface struct {
	clientCache kcpclient.Cache[*policyv1client.PolicyV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *podDisruptionBudgetsClusterInterface) Cluster(clusterPath logicalcluster.Path) PodDisruptionBudgetsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &podDisruptionBudgetsNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all PodDisruptionBudgets across all clusters.
func (c *podDisruptionBudgetsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*policyv1.PodDisruptionBudgetList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PodDisruptionBudgets(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all PodDisruptionBudgets across all clusters.
func (c *podDisruptionBudgetsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PodDisruptionBudgets(metav1.NamespaceAll).Watch(ctx, opts)
}

// PodDisruptionBudgetsNamespacer can scope to objects within a namespace, returning a policyv1client.PodDisruptionBudgetInterface.
type PodDisruptionBudgetsNamespacer interface {
	Namespace(string) policyv1client.PodDisruptionBudgetInterface
}

type podDisruptionBudgetsNamespacer struct {
	clientCache kcpclient.Cache[*policyv1client.PolicyV1Client]
	clusterPath logicalcluster.Path
}

func (n *podDisruptionBudgetsNamespacer) Namespace(namespace string) policyv1client.PodDisruptionBudgetInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).PodDisruptionBudgets(namespace)
}
