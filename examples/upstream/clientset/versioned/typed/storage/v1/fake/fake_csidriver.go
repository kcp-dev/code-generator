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

package fake

import (
	"context"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v3"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/storage/v1"
)

var csidriversResource = v1.SchemeGroupVersion.WithResource("csidrivers")

var csidriversKind = v1.SchemeGroupVersion.WithKind("CSIDriver")

// cSIDriversClusterClient implements cSIDriverInterface
type cSIDriversClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *cSIDriversClusterClient) Cluster(clusterPath logicalcluster.Path) *kcp.CSIDriverNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &cSIDriversNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of CSIDrivers that match those selectors.
func (c *cSIDriversClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CSIDriverList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(csidriversResource, csidriversKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1.CSIDriverList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CSIDriverList{ListMeta: obj.(*v1.CSIDriverList).ListMeta}
	for _, item := range obj.(*v1.CSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cSIDrivers across all clusters.
func (c *cSIDriversClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(csidriversResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}
