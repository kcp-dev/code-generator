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
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	upstreambatchv1client "k8s.io/client-go/kubernetes/typed/batch/v1"
	"k8s.io/client-go/testing"
	batchv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/batch/v1"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/batch/v1"
)

var cronjobsResource = v1.SchemeGroupVersion.WithResource("cronjobs")

var cronjobsKind = v1.SchemeGroupVersion.WithKind("CronJob")

// cronJobsClusterClient implements cronJobInterface
type cronJobsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *cronJobsClusterClient) Cluster(clusterPath logicalcluster.Path) *kcp.CronJobNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &cronJobsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of CronJobs that match those selectors.
func (c *cronJobsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CronJobList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(cronjobsResource, cronjobsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1.CronJobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CronJobList{ListMeta: obj.(*v1.CronJobList).ListMeta}
	for _, item := range obj.(*v1.CronJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cronJobs across all clusters.
func (c *cronJobsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(cronjobsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type cronJobsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *cronJobsNamespacer) Namespace(namespace string) upstreambatchv1client.CronJobInterface {
	return &configMapsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type cronJobsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *cronJobsClient) Create(ctx context.Context, cronJob *v1.CronJob, opts metav1.CreateOptions) (*v1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(cronjobsResource, c.ClusterPath, c.Namespace, cronJob), &v1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CronJob), err
}

func (c *cronJobsClient) Update(ctx context.Context, cronJob *v1.CronJob, opts metav1.CreateOptions) (*v1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(cronjobsResource, c.ClusterPath, c.Namespace, cronJob), &v1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CronJob), err
}

func (c *cronJobsClient) UpdateStatus(ctx context.Context, cronJob *v1.CronJob, opts metav1.CreateOptions) (*v1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(cronjobsResource, c.ClusterPath, "status", c.Namespace, cronJob), &v1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CronJob), err
}

func (c *cronJobsClient) Delete(ctx context.Context, name string, opts metav1.CreateOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(cronjobsResource, c.ClusterPath, c.Namespace, name, opts), &v1.CronJob{})
	return err
}

func (c *cronJobsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(cronjobsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1.CronJobList{})
	return err
}

func (c *cronJobsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(cronjobsResource, c.ClusterPath, c.Namespace, name), &v1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CronJob), err
}

// List takes label and field selectors, and returns the list of v1.CronJob that match those selectors.
func (c *cronJobsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1.CronJobList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(cronjobsResource, cronjobsKind, c.ClusterPath, c.Namespace, opts), &v1.CronJobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CronJobList{ListMeta: obj.(*v1.CronJobList).ListMeta}
	for _, item := range obj.(*v1.CronJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *cronJobsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(cronjobsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *cronJobsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(cronjobsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CronJob), err
}

func (c *cronJobsClient) Apply(ctx context.Context, applyConfiguration *batchv1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (*v1.CronJob, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(cronjobsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CronJob), err
}
