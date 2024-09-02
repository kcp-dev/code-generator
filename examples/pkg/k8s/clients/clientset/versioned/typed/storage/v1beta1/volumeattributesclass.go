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

package v1beta1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	storagev1beta1 "acme.corp/pkg/apis/storage/v1beta1"
	storagev1beta1client "acme.corp/pkg/generated/clientset/versioned/typed/storage/v1beta1"
)

// VolumeAttributesClassesClusterGetter has a method to return a VolumeAttributesClassClusterInterface.
// A group's cluster client should implement this interface.
type VolumeAttributesClassesClusterGetter interface {
	VolumeAttributesClasses() VolumeAttributesClassClusterInterface
}

// VolumeAttributesClassClusterInterface can operate on VolumeAttributesClasses across all clusters,
// or scope down to one cluster and return a storagev1beta1client.VolumeAttributesClassInterface.
type VolumeAttributesClassClusterInterface interface {
	Cluster(logicalcluster.Path) storagev1beta1client.VolumeAttributesClassInterface
	List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.VolumeAttributesClassList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type volumeAttributesClassesClusterInterface struct {
	clientCache kcpclient.Cache[*storagev1beta1client.StorageV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *volumeAttributesClassesClusterInterface) Cluster(clusterPath logicalcluster.Path) storagev1beta1client.VolumeAttributesClassInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).VolumeAttributesClasses()
}

// List returns the entire collection of all VolumeAttributesClasses across all clusters.
func (c *volumeAttributesClassesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.VolumeAttributesClassList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).VolumeAttributesClasses().List(ctx, opts)
}

// Watch begins to watch all VolumeAttributesClasses across all clusters.
func (c *volumeAttributesClassesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).VolumeAttributesClasses().Watch(ctx, opts)
}
