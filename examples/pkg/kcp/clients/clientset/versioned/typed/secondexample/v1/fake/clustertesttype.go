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

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	secondexamplev1 "acme.corp/pkg/apis/secondexample/v1"
	applyconfigurationssecondexamplev1 "acme.corp/pkg/generated/applyconfigurations/secondexample/v1"
	secondexamplev1client "acme.corp/pkg/generated/clientset/versioned/typed/secondexample/v1"
)

var clusterTestTypesResource = schema.GroupVersionResource{Group: "secondexample", Version: "v1", Resource: "clustertesttypes"}
var clusterTestTypesKind = schema.GroupVersionKind{Group: "secondexample", Version: "v1", Kind: "ClusterTestType"}

type clusterTestTypesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterTestTypesClusterClient) Cluster(clusterPath logicalcluster.Path) secondexamplev1client.ClusterTestTypeInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &clusterTestTypesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors across all clusters.
func (c *clusterTestTypesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*secondexamplev1.ClusterTestTypeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterTestTypesResource, clusterTestTypesKind, logicalcluster.Wildcard, opts), &secondexamplev1.ClusterTestTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &secondexamplev1.ClusterTestTypeList{ListMeta: obj.(*secondexamplev1.ClusterTestTypeList).ListMeta}
	for _, item := range obj.(*secondexamplev1.ClusterTestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ClusterTestTypes across all clusters.
func (c *clusterTestTypesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterTestTypesResource, logicalcluster.Wildcard, opts))
}

type clusterTestTypesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *clusterTestTypesClient) Create(ctx context.Context, clusterTestType *secondexamplev1.ClusterTestType, opts metav1.CreateOptions) (*secondexamplev1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(clusterTestTypesResource, c.ClusterPath, clusterTestType), &secondexamplev1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.ClusterTestType), err
}

func (c *clusterTestTypesClient) Update(ctx context.Context, clusterTestType *secondexamplev1.ClusterTestType, opts metav1.UpdateOptions) (*secondexamplev1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(clusterTestTypesResource, c.ClusterPath, clusterTestType), &secondexamplev1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.ClusterTestType), err
}

func (c *clusterTestTypesClient) UpdateStatus(ctx context.Context, clusterTestType *secondexamplev1.ClusterTestType, opts metav1.UpdateOptions) (*secondexamplev1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(clusterTestTypesResource, c.ClusterPath, "status", clusterTestType), &secondexamplev1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.ClusterTestType), err
}

func (c *clusterTestTypesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(clusterTestTypesResource, c.ClusterPath, name, opts), &secondexamplev1.ClusterTestType{})
	return err
}

func (c *clusterTestTypesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(clusterTestTypesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &secondexamplev1.ClusterTestTypeList{})
	return err
}

func (c *clusterTestTypesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*secondexamplev1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(clusterTestTypesResource, c.ClusterPath, name), &secondexamplev1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.ClusterTestType), err
}

// List takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors.
func (c *clusterTestTypesClient) List(ctx context.Context, opts metav1.ListOptions) (*secondexamplev1.ClusterTestTypeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterTestTypesResource, clusterTestTypesKind, c.ClusterPath, opts), &secondexamplev1.ClusterTestTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &secondexamplev1.ClusterTestTypeList{ListMeta: obj.(*secondexamplev1.ClusterTestTypeList).ListMeta}
	for _, item := range obj.(*secondexamplev1.ClusterTestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *clusterTestTypesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterTestTypesResource, c.ClusterPath, opts))
}

func (c *clusterTestTypesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*secondexamplev1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterTestTypesResource, c.ClusterPath, name, pt, data, subresources...), &secondexamplev1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.ClusterTestType), err
}

func (c *clusterTestTypesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationssecondexamplev1.ClusterTestTypeApplyConfiguration, opts metav1.ApplyOptions) (*secondexamplev1.ClusterTestType, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterTestTypesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &secondexamplev1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.ClusterTestType), err
}

func (c *clusterTestTypesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationssecondexamplev1.ClusterTestTypeApplyConfiguration, opts metav1.ApplyOptions) (*secondexamplev1.ClusterTestType, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterTestTypesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &secondexamplev1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.ClusterTestType), err
}
