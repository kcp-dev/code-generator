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

package internal

const wrappedInterfacesTempl = `
package generated

import (
	"github.com/kcp-dev/kcp-client-wrappers/kcp"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"

	{{$name := .InputPath}}
	{{ range .APIs }}
	{{.PkgName}}{{.Version}} "{{$name}}/pkg/apis/{{.PkgName}}/{{.Version}}"
	{{ end }}
)

func NewForConfig(config *rest.Config) (*ClusterClient, error) {
	client, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP client: %w", err)
	}

	clusterRoundTripper := kcp.NewClusterRoundTripper(client.Transport)
	client.Transport = clusterRoundTripper

	delegate, err := kubernetes.NewForConfigAndClient(config, client)
	if err != nil {
		return nil, fmt.Errorf("error creating delegate clientset: %w", err)
	}

	return &ClusterClient{
		delegate: delegate,
	}, nil
}

type ClusterClient struct {
	delegate {{.InterfaceName}}.Interface
}

func (c *ClusterClient) Cluster(cluster string) {{.InterfaceName}}.Interface {
	return &wrappedInterface{
		cluster:  cluster,
		delegate: c.delegate,
	}
}

type wrappedInterface struct {
	cluster  string
	delegate kubernetes.Interface
}

func (w *wrappedInterface) Discovery() discovery.DiscoveryInterface {
	return w.delegate.Discovery()
}

{{ range .APIs }}
func (w *wrappedInterface) {{.PkgNameUpperFirst}}{{.VersionUpperFirst}}() {{.PkgName}}{{.Version}}.{{.PkgNameUpperFirst}}{{.VersionUpperFirst}}Interface {
	return &wrapped{{.PkgNameUpperFirst}}{{.VersionUpperFirst}}{
		cluster:  w.cluster,
		delegate: w.delegate.{{.PkgNameUpperFirst}}{{.VersionUpperFirst}}(),
	}
}
{{ end }}

`

const commonTempl = `
package generated

import (
	"ctx"
	{{.Name}}api{{.Version}} "{{.APIPath}}"
	{{.Name}}{{.Version}} "{{.ClientPath}}typed/{{.Name}}/{{.Version}}"

	"github.com/kcp-dev/kcp-client-wrappers/kcp"
)

type wrapped{{.NameUpperFirst}}{{.VersionUpperFirst}} struct {
	cluster  string
	delegate {{.Name}}{{.Version}}.{{.NameUpperFirst}}{{.VersionUpperFirst}}Interface
}

func (w *wrapped{{.NameUpperFirst}}{{.VersionUpperFirst}}) RESTClient() rest.Interface {
	return w.delegate.RESTClient()
}

func (w *wrapped{{.Name}}) checkCluster(ctx context.Context) (context.Context, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		return kcp.WithCluster(ctx, w.cluster), nil
	} else if ctxCluster != w.cluster {
		return ctx, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return ctx, nil
}
`

const wrapperMethodsTempl = `
func (w *wrapped{{.PkgNameUpperFirst}}{{.VersionUpperFirst}}) {{.Name}}s() {{.PkgName}}{{.Version}}.{{.Name}}Interface {
	return &wrapped{{.Name}}{
		cluster:  w.cluster,
		delegate: w.delegate.{{.Name}}s(),
	}
}

type wrapped{{.Name}} struct {
	cluster  string
	delegate {{.PkgName}}.{{.Name}}Interface
}

func (w *wrapped{{.Name}}) Create(ctx context.Context, {{.NameLowerFirst}} *{{.PkgName}}api{{.Version}}.{{.Name}}, opts metav1.CreateOptions) (*{{.PkgName}}api{{.Version}}.{{.Name}}, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, {{.NameLowerFirst}}}, opts)
}

func (w *wrapped{{.Name}}) Update(ctx context.Context, {{.NameLowerFirst}}} *{{.PkgName}}api{{.Version}}.{{.Name}}, opts metav1.UpdateOptions) (*{{.PkgName}}api{{.Version}}.{{.Name}}, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, {{.NameLowerFirst}}}, opts)
}

func (w *wrapped{{.Name}}) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w *wrapped{{.Name}}) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Delete(ctx, opts, listOpts)
}

func (w *wrapped{{.Name}}) Get(ctx context.Context, name string, opts metav1.GetOptions) (*{{.PkgName}}api{{.Version}}.{{.Name}}, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, opts, listOpts)
}

func (w *wrapped{{.Name}}) List(ctx context.Context, opts metav1.ListOptions) (*{{.PkgName}}api{{.Version}}.{{.Name}}List, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w *wrapped{{.Name}}) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w *wrapped{{.Name}}) Patch(ctx context.Context, name string, pt apiTypes.PatchapiType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *{{.PkgName}}api{{.Version}}.{{.Name}}, err error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources)
}
`

// Do we make this configurable?
// TODO: reformat it to be able to configure Copyright Year
const HeaderText = `
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
// +build !ignore_autogenerated
`
