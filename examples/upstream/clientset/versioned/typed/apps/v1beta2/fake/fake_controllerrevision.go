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
	json "encoding/json"
	"fmt"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	upstreamappsv1beta2client "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	"k8s.io/client-go/testing"
	appsv1beta2 "k8s.io/code-generator/examples/upstream/applyconfiguration/apps/v1beta2"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/apps/v1beta2"
)

var controllerrevisionsResource = v1beta2.SchemeGroupVersion.WithResource("controllerrevisions")

var controllerrevisionsKind = v1beta2.SchemeGroupVersion.WithKind("ControllerRevision")

// controllerRevisionsClusterClient implements controllerRevisionInterface
type controllerRevisionsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *controllerRevisionsClusterClient) Cluster(clusterPath logicalcluster.Path) *kcp.ControllerRevisionNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &controllerRevisionsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ControllerRevisions that match those selectors.
func (c *controllerRevisionsClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.ControllerRevisionList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(controllerrevisionsResource, controllerrevisionsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta2.ControllerRevisionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.ControllerRevisionList{ListMeta: obj.(*v1beta2.ControllerRevisionList).ListMeta}
	for _, item := range obj.(*v1beta2.ControllerRevisionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested controllerRevisions across all clusters.
func (c *controllerRevisionsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(controllerrevisionsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type controllerRevisionsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *controllerRevisionsNamespacer) Namespace(namespace string) upstreamappsv1beta2client.ControllerRevisionInterface {
	return &configMapsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type controllerRevisionsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *controllerRevisionsClient) Create(ctx context.Context, controllerRevision *v1beta2.ControllerRevision, opts metav1.CreateOptions) (*v1beta2.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(controllerrevisionsResource, c.ClusterPath, c.Namespace, controllerRevision), &v1beta2.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ControllerRevision), err
}

func (c *controllerRevisionsClient) Update(ctx context.Context, controllerRevision *v1beta2.ControllerRevision, opts metav1.CreateOptions) (*v1beta2.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(controllerrevisionsResource, c.ClusterPath, c.Namespace, controllerRevision), &v1beta2.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ControllerRevision), err
}

func (c *controllerRevisionsClient) UpdateStatus(ctx context.Context, controllerRevision *v1beta2.ControllerRevision, opts metav1.CreateOptions) (*v1beta2.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(controllerrevisionsResource, c.ClusterPath, "status", c.Namespace, controllerRevision), &v1beta2.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ControllerRevision), err
}

func (c *controllerRevisionsClient) Delete(ctx context.Context, name string, opts metav1.CreateOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(controllerrevisionsResource, c.ClusterPath, c.Namespace, name, opts), &v1beta2.ControllerRevision{})
	return err
}

func (c *controllerRevisionsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(controllerrevisionsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.ControllerRevisionList{})
	return err
}

func (c *controllerRevisionsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta2.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(controllerrevisionsResource, c.ClusterPath, c.Namespace, name), &v1beta2.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ControllerRevision), err
}

// List takes label and field selectors, and returns the list of v1beta2.ControllerRevision that match those selectors.
func (c *controllerRevisionsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta2.ControllerRevisionList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(controllerrevisionsResource, controllerrevisionsKind, c.ClusterPath, c.Namespace, opts), &v1beta2.ControllerRevisionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.ControllerRevisionList{ListMeta: obj.(*v1beta2.ControllerRevisionList).ListMeta}
	for _, item := range obj.(*v1beta2.ControllerRevisionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *controllerRevisionsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(controllerrevisionsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *controllerRevisionsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta2.ControllerRevision, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(controllerrevisionsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1beta2.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ControllerRevision), err
}

func (c *controllerRevisionsClient) Apply(ctx context.Context, applyConfiguration *appsv1beta2.ControllerRevisionApplyConfiguration, opts metav1.ApplyOptions) (*v1beta2.ControllerRevision, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(controllerrevisionsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1beta2.ControllerRevision{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ControllerRevision), err
}
