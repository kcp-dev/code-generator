//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The KCP Authors.

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

// Code auto-generated. DO NOT EDIT.

package v1

import (
	"context"
	"fmt"
	rbacapiv1 "github.com/kcp-dev/client-gen/testdata/pkg/apis/rbac/v1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"

	kcp "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type WrappedRbacV1 struct {
	Cluster  logicalcluster.LogicalCluster
	Delegate rbacv1.RbacV1Interface
}

func (w *WrappedRbacV1) RESTClient() rest.Interface {
	return w.Delegate.RESTClient()
}

func (w *WrappedRbacV1) ClusterRoles() rbacv1.ClusterRoleInterface {
	return &wrappedClusterRole{
		cluster:  w.Cluster,
		delegate: w.Delegate.ClusterRoles(),
	}
}

type wrappedClusterRole struct {
	cluster  logicalcluster.LogicalCluster
	delegate rbacv1.ClusterRoleInterface
}

func (w *wrappedClusterRole) checkCluster(ctx context.Context) (context.Context, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		return kcp.WithCluster(ctx, w.cluster), nil
	} else if ctxCluster != w.cluster {
		return ctx, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return ctx, nil
}

func (w *wrappedClusterRole) Create(ctx context.Context, clusterRole *rbacapiv1.ClusterRole, opts metav1.CreateOptions) (*rbacapiv1.ClusterRole, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, clusterRole, opts)
}

func (w *wrappedClusterRole) Update(ctx context.Context, clusterRole *rbacapiv1.ClusterRole, opts metav1.UpdateOptions) (*rbacapiv1.ClusterRole, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, clusterRole, opts)
}

func (w *wrappedClusterRole) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w *wrappedClusterRole) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listopts metav1.ListOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listopts)
}

func (w *wrappedClusterRole) Get(ctx context.Context, name string, opts metav1.GetOptions) (*rbacapiv1.ClusterRole, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w *wrappedClusterRole) List(ctx context.Context, opts metav1.ListOptions) (*rbacapiv1.ClusterRoleList, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w *wrappedClusterRole) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w *wrappedClusterRole) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *rbacapiv1.ClusterRole, err error) {
	ctx, err = w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}

func (w *WrappedRbacV1) ClusterRoleBindings(namespace) rbacv1.ClusterRoleBindingInterface {
	return &wrappedClusterRoleBinding{
		cluster:  w.Cluster,
		delegate: w.Delegate.ClusterRoleBindings(namespace),
	}
}

type wrappedClusterRoleBinding struct {
	cluster  logicalcluster.LogicalCluster
	delegate rbacv1.ClusterRoleBindingInterface
}

func (w *wrappedClusterRoleBinding) checkCluster(ctx context.Context) (context.Context, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		return kcp.WithCluster(ctx, w.cluster), nil
	} else if ctxCluster != w.cluster {
		return ctx, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return ctx, nil
}

func (w *wrappedClusterRoleBinding) Create(ctx context.Context, clusterRoleBinding *rbacapiv1.ClusterRoleBinding, opts metav1.CreateOptions) (*rbacapiv1.ClusterRoleBinding, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, clusterRoleBinding, opts)
}

func (w *wrappedClusterRoleBinding) Update(ctx context.Context, clusterRoleBinding *rbacapiv1.ClusterRoleBinding, opts metav1.UpdateOptions) (*rbacapiv1.ClusterRoleBinding, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, clusterRoleBinding, opts)
}

func (w *wrappedClusterRoleBinding) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w *wrappedClusterRoleBinding) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listopts metav1.ListOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listopts)
}

func (w *wrappedClusterRoleBinding) Get(ctx context.Context, name string, opts metav1.GetOptions) (*rbacapiv1.ClusterRoleBinding, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w *wrappedClusterRoleBinding) List(ctx context.Context, opts metav1.ListOptions) (*rbacapiv1.ClusterRoleBindingList, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w *wrappedClusterRoleBinding) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w *wrappedClusterRoleBinding) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *rbacapiv1.ClusterRoleBinding, err error) {
	ctx, err = w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}
