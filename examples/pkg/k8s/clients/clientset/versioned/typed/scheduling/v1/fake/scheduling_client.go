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

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	schedulingv1 "acme.corp/pkg/generated/clientset/versioned/typed/scheduling/v1"
	kcpschedulingv1 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/scheduling/v1"
)

var _ kcpschedulingv1.SchedulingV1ClusterInterface = (*SchedulingV1ClusterClient)(nil)

type SchedulingV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *SchedulingV1ClusterClient) Cluster(clusterPath logicalcluster.Path) schedulingv1.SchedulingV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &SchedulingV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *SchedulingV1ClusterClient) PriorityClasses() kcpschedulingv1.PriorityClassClusterInterface {
	return &priorityClassesClusterClient{Fake: c.Fake}
}

var _ schedulingv1.SchedulingV1Interface = (*SchedulingV1Client)(nil)

type SchedulingV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *SchedulingV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *SchedulingV1Client) PriorityClasses() schedulingv1.PriorityClassInterface {
	return &priorityClassesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
