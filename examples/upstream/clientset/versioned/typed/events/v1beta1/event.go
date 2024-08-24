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

package v1beta1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/events/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreameventsv1beta1client "k8s.io/client-go/kubernetes/typed/events/v1beta1"
)

// EventsClusterGetter has a method to return a EventClusterInterface.
// A group's client should implement this interface.
type EventsClusterGetter interface {
	Events() EventClusterInterface
}

// EventClusterInterface has methods to work with Event resources.
type EventClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.EventList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) EventNamespacer
	EventExpansion
}

type eventsClusterInterface struct {
	clientCache kcpclient.Cache[*upstreameventsv1beta1client.EventsV1beta1Client]
}
