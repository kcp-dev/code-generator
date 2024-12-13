package clientgen

import (
	"io"
	"strings"
	"text/template"

	"github.com/kcp-dev/code-generator/v2/pkg/parser"
	"github.com/kcp-dev/code-generator/v2/pkg/util"
)

type FakeClientset struct {
	// Name is the name of the clientset, e.g. "kubernetes"
	Name string

	// Groups are the groups in this client-set.
	Groups []parser.Group

	// PackagePath is the package under which this client-set will be exposed.
	// TODO(skuznets) we should be able to figure this out from the output dir, ideally
	PackagePath string

	// SingleClusterClientPackagePath is the root directory under which single-cluster-aware clients exist.
	// e.g. "k8s.io/client-go/kubernetes"
	SingleClusterClientPackagePath string
}

func (c *FakeClientset) WriteContent(w io.Writer) error {
	templ, err := template.New("fakeClientset").Funcs(template.FuncMap{
		"upperFirst": util.UpperFirst,
		"lowerFirst": util.LowerFirst,
		"toLower":    strings.ToLower,
	}).Parse(fakeClientset)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"name":                           c.Name,
		"packagePath":                    c.PackagePath,
		"groups":                         c.Groups,
		"singleClusterClientPackagePath": c.SingleClusterClientPackagePath,
	}
	return templ.Execute(w, m)
}

var fakeClientset = `
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	kcpfakediscovery "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/discovery/fake"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"

	client "{{.singleClusterClientPackagePath}}"
	clientscheme "{{.singleClusterClientPackagePath}}/scheme"

	kcpclient "{{.packagePath}}"
{{range .groups}}	{{.PackageAlias}} "{{$.singleClusterClientPackagePath}}/typed/{{.Group.PackageName}}/{{.Version.PackageName}}"
{{end -}}
{{range .groups}}	kcp{{.PackageAlias}} "{{$.packagePath}}/typed/{{.Group.PackageName}}/{{.Version.PackageName}}"
{{end -}}
{{range .groups}}	fake{{.PackageAlias}} "{{$.packagePath}}/typed/{{.Group.PackageName}}/{{.Version.PackageName}}/fake"
{{end -}}
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *ClusterClientset {
	o := kcptesting.NewObjectTracker(clientscheme.Scheme, clientscheme.Codecs.UniversalDecoder())
	o.AddAll(objects...)

	cs := &ClusterClientset{Fake: &kcptesting.Fake{}, tracker: o}
	cs.discovery = &kcpfakediscovery.FakeDiscovery{Fake: cs.Fake, ClusterPath: logicalcluster.Wildcard}
	cs.AddReactor("*", "*", kcptesting.ObjectReaction(o))
	cs.AddWatchReactor("*", kcptesting.WatchReaction(o))

	return cs
}

var _ kcpclient.ClusterInterface = (*ClusterClientset)(nil)

// ClusterClientset contains the clients for groups.
type ClusterClientset struct {
	*kcptesting.Fake
	discovery *kcpfakediscovery.FakeDiscovery
	tracker   kcptesting.ObjectTracker
}

// Discovery retrieves the DiscoveryClient
func (c *ClusterClientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *ClusterClientset) Tracker() kcptesting.ObjectTracker {
	return c.tracker
}

{{range .groups}}
// {{.GroupGoName}}{{.Version}} retrieves the {{.GroupGoName}}{{.Version}}ClusterClient.  
func (c *ClusterClientset) {{.GroupGoName}}{{.Version}}() kcp{{.PackageAlias}}.{{.GroupGoName}}{{.Version}}ClusterInterface {
	return &fake{{.PackageAlias}}.{{.GroupGoName}}{{.Version}}ClusterClient{Fake: c.Fake}
}
{{end -}}

// Cluster scopes this clientset to one cluster.
func (c *ClusterClientset) Cluster(clusterPath logicalcluster.Path) client.Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &Clientset{
		Fake: c.Fake,
		discovery: &kcpfakediscovery.FakeDiscovery{Fake: c.Fake, ClusterPath: clusterPath},
		tracker: c.tracker.Cluster(clusterPath),
		clusterPath: clusterPath,
	}
}

var _ client.Interface = (*Clientset)(nil)

// Clientset contains the clients for groups.
type Clientset struct {
	*kcptesting.Fake
	discovery *kcpfakediscovery.FakeDiscovery
	tracker   kcptesting.ScopedObjectTracker
	clusterPath logicalcluster.Path
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() kcptesting.ScopedObjectTracker {
	return c.tracker
}

{{range .groups}}
// {{.GroupGoName}}{{.Version}} retrieves the {{.GroupGoName}}{{.Version}}Client.  
func (c *Clientset) {{.GroupGoName}}{{.Version}}() {{.PackageAlias}}.{{.GroupGoName}}{{.Version}}Interface {
	return &fake{{.PackageAlias}}.{{.GroupGoName}}{{.Version}}Client{Fake: c.Fake, ClusterPath: c.clusterPath}
}
{{end -}}
`
