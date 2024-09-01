/*
Copyright 2015 The Kubernetes Authors.

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

package fake

import (
	"io"
	"path"
	"strings"
	"sync"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
	"k8s.io/klog/v2"

	"k8s.io/code-generator/cmd/client-gen/generators/util"
	"k8s.io/code-generator/pkg/static"
)

// genFakeForType produces a file for each top-level type.
type genFakeForType struct {
	generator.GoGenerator
	outputPackage                        string // Must be a Go import-path
	group                                string
	version                              string
	groupGoName                          string
	inputPackage                         string
	typeToMatch                          *types.Type
	imports                              namer.ImportTracker
	applyConfigurationPackage            string
	typedClientPackage                   string
	singleClusterTypedClientsPackagePath string
	singleClusterApplyConfigPackagePath  string
	staticFakeExpansions                 []string
}

var _ generator.Generator = &genFakeForType{}

var titler = cases.Title(language.Und)

// Filter ignores all but one type because we're making a single file per type.
func (g *genFakeForType) Filter(c *generator.Context, t *types.Type) bool { return t == g.typeToMatch }

func (g *genFakeForType) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *genFakeForType) Imports(c *generator.Context) (imports []string) {
	imports = g.imports.ImportLines()
	imports = append(imports, "kcptesting \"github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing\"")

	if len(g.singleClusterTypedClientsPackagePath) > 0 {
		imports = append(imports, "upstream"+strings.ToLower(g.groupGoName+g.version+"client \""+g.singleClusterTypedClientsPackagePath+"/"+g.groupGoName+"/"+g.version+"\""))
	}

	imports = append(imports, "kcp "+"\""+g.typedClientPackage+"\"")
	imports = append(imports, "restclient \"k8s.io/client-go/rest\"")
	imports = append(imports, "fakerest \"k8s.io/client-go/rest/fake\"")
	imports = append(imports, "policyv1 \"k8s.io/api/policy/v1\"")
	imports = append(imports, "policyv1beta1 \"k8s.io/api/policy/v1beta1\"")
	imports = append(imports, "autoscalingv1 \"k8s.io/api/autoscaling/v1\"")
	imports = append(imports, "authenticationv1 \"k8s.io/api/authentication/v1\"")
	imports = append(imports, "metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"")
	imports = append(imports, "applyconfigurationsautoscalingv1 \"k8s.io/client-go/applyconfigurations/autoscaling/v1\"")

	return imports
}

// GenerateType makes the body of a file implementing the individual typed client for type t.
func (g *genFakeForType) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	pkg := path.Base(t.Name.Package)
	tags, err := util.ParseClientGenTags(append(t.SecondClosestCommentLines, t.CommentLines...))
	if err != nil {
		return err
	}

	namespaced := !tags.NonNamespaced

	const pkgClientGoTesting = "k8s.io/client-go/testing"
	m := map[string]interface{}{
		"type":               t,
		"inputType":          t,
		"resultType":         t,
		"subresourcePath":    "",
		"package":            pkg,
		"Package":            namer.IC(pkg),
		"namespaced":         namespaced,
		"Group":              namer.IC(g.group),
		"GroupGoName":        g.groupGoName,
		"Version":            namer.IC(g.version),
		"version":            g.version,
		"SchemeGroupVersion": c.Universe.Type(types.Name{Package: t.Name.Package, Name: "SchemeGroupVersion"}),
		"CreateOptions":      c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "CreateOptions"}),
		"DeleteOptions":      c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "DeleteOptions"}),
		"GetOptions":         c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "GetOptions"}),
		"ListOptions":        c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ListOptions"}),
		"PatchOptions":       c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "PatchOptions"}),
		"ApplyOptions":       c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ApplyOptions"}),
		"UpdateOptions":      c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "UpdateOptions"}),
		"Everything":         c.Universe.Function(types.Name{Package: "k8s.io/apimachinery/pkg/labels", Name: "Everything"}),
		"PatchType":          c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/types", Name: "PatchType"}),
		"ApplyPatchType":     c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/types", Name: "ApplyPatchType"}),
		"watchInterface":     c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/watch", Name: "Interface"}),
		"jsonMarshal":        c.Universe.Type(types.Name{Package: "encoding/json", Name: "Marshal"}),

		"NewRootListActionWithOptions":              c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootListActionWithOptions"}),
		"NewListActionWithOptions":                  c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewListActionWithOptions"}),
		"NewRootGetActionWithOptions":               c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootGetActionWithOptions"}),
		"NewGetActionWithOptions":                   c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewGetActionWithOptions"}),
		"NewRootDeleteActionWithOptions":            c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootDeleteActionWithOptions"}),
		"NewDeleteActionWithOptions":                c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewDeleteActionWithOptions"}),
		"NewRootDeleteCollectionActionWithOptions":  c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootDeleteCollectionActionWithOptions"}),
		"NewDeleteCollectionActionWithOptions":      c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewDeleteCollectionActionWithOptions"}),
		"NewRootUpdateActionWithOptions":            c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootUpdateActionWithOptions"}),
		"NewUpdateActionWithOptions":                c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewUpdateActionWithOptions"}),
		"NewRootCreateActionWithOptions":            c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootCreateActionWithOptions"}),
		"NewCreateActionWithOptions":                c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewCreateActionWithOptions"}),
		"NewRootWatchActionWithOptions":             c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootWatchActionWithOptions"}),
		"NewWatchActionWithOptions":                 c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewWatchActionWithOptions"}),
		"NewCreateSubresourceActionWithOptions":     c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewCreateSubresourceActionWithOptions"}),
		"NewRootCreateSubresourceActionWithOptions": c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootCreateSubresourceActionWithOptions"}),
		"NewUpdateSubresourceActionWithOptions":     c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewUpdateSubresourceActionWithOptions"}),
		"NewGetSubresourceActionWithOptions":        c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewGetSubresourceActionWithOptions"}),
		"NewRootGetSubresourceActionWithOptions":    c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootGetSubresourceActionWithOptions"}),
		"NewRootUpdateSubresourceActionWithOptions": c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootUpdateSubresourceActionWithOptions"}),
		"NewRootPatchActionWithOptions":             c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootPatchActionWithOptions"}),
		"NewPatchActionWithOptions":                 c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewPatchActionWithOptions"}),
		"NewRootPatchSubresourceActionWithOptions":  c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewRootPatchSubresourceActionWithOptions"}),
		"NewPatchSubresourceActionWithOptions":      c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "NewPatchSubresourceActionWithOptions"}),
		"ExtractFromListOptions":                    c.Universe.Function(types.Name{Package: pkgClientGoTesting, Name: "ExtractFromListOptions"}),
	}

	generateApply := len(g.applyConfigurationPackage) > 0

	_, gvString := util.ParsePathGroupVersion(g.inputPackage)
	if g.singleClusterTypedClientsPackagePath != "" {
		m["inputApplyConfig"] = types.Ref(path.Join(g.singleClusterApplyConfigPackagePath, gvString), t.Name.Name+"ApplyConfiguration")
	} else {
		m["inputApplyConfig"] = types.Ref(path.Join(g.applyConfigurationPackage, gvString), t.Name.Name+"ApplyConfiguration")
	}

	m["kcpClusterResultType"] = types.Ref(g.typedClientPackage, t.Name.Name)
	m["upstreamClientInterface"] = "upstream" + strings.ToLower(g.groupGoName+g.version+"client.") + t.Name.Name + "Interface"

	// Global structs for the types.
	sw.Do(resource, m)
	sw.Do(kind, m)

	sw.Do(structClusterGeneric, m)

	if namespaced {
		sw.Do(methodNamespacedClusterClusterTemplate, m)
		sw.Do(methodNamespacedClusterListTemplate, m)

		sw.Do(methodNamespacedClusterWatchTemplate, m)
		sw.Do(structNamespacer, m)
		sw.Do(methodNamespacerNamespaceTemplate, m)

		sw.Do(structNamespacedClient, m)
		sw.Do(methodNamespacedClientCreateTemplate, m)
		sw.Do(methodNamespacedClientUpdateTemplate, m)
		sw.Do(methodNamespacedClientUpdateStatusTemplate, m)
		sw.Do(methodNamespacedClientDeleteTemplate, m)
		sw.Do(methodNamespacedClientDeleteCollectionTemplate, m)
		sw.Do(methodNamespacedClientGetTemplate, m)
		sw.Do(methodNamespacedClientListTemplate, m)
		sw.Do(methodNamespacedClientWatchTemplate, m)
		sw.Do(methodNamespacedClientPatchTemplate, m)
		if generateApply {
			sw.Do(methodNamespacedClientApplyTemplate, m)
			sw.Do(methodNamespacedClientApplyStatusTemplate, m)
		}
	} else {
		// Non namespaced types.
		sw.Do(methodClusterClusterTemplate, m)
		sw.Do(methodClusterListTemplate, m)
		sw.Do(methodClusterWatchTemplate, m)

		sw.Do(structClient, m)

		sw.Do(methodNamespacedClientCreateTemplate, m)
		sw.Do(methodNamespacedClientUpdateTemplate, m)
		sw.Do(methodNamespacedClientUpdateStatusTemplate, m)
		sw.Do(methodNamespacedClientDeleteTemplate, m)
		sw.Do(methodNamespacedClientDeleteCollectionTemplate, m)
		sw.Do(methodNamespacedClientGetTemplate, m)
		sw.Do(methodNamespacedClientListTemplate, m)
		sw.Do(methodNamespacedClientWatchTemplate, m)
		sw.Do(methodNamespacedClientPatchTemplate, m)
		if generateApply {
			sw.Do(methodNamespacedClientApplyTemplate, m)
			sw.Do(methodNamespacedClientApplyStatusTemplate, m)
		}
	}

	if len(g.staticFakeExpansions) > 0 {
		for _, expansion := range g.staticFakeExpansions {
			source, target := strings.Split(expansion, ":")[0], strings.Split(expansion, ":")[1]
			if source == t.String() {
				klog.Infof("Found fake expansion override for %s", t.String())
				override := static.GetClientSetFakeExpansions(target)
				sw.Do(override, m)
			}
		}
	}

	_, typeGVString := util.ParsePathGroupVersion(g.inputPackage)

	// generate extended client methods
	var once sync.Once
	for _, e := range tags.Extensions {
		if e.SubResourcePath == "scale" {
			// add scale methods.
			once.Do(func() {
				sw.Do(methodNamespacedClientGetScaleTemplate, m)
				sw.Do(methodNamespacedClientUpdateScaleTemplate, m)
				sw.Do(methodNamespacedClientApplyScaleTemplate, m)
			})
		}

		if e.HasVerb("apply") && !generateApply {
			continue
		}
		inputType := *t
		resultType := *t
		inputGVString := typeGVString
		if len(e.InputTypeOverride) > 0 {
			if name, pkg := e.Input(); len(pkg) > 0 {
				_, inputGVString = util.ParsePathGroupVersion(pkg)
				newType := c.Universe.Type(types.Name{Package: pkg, Name: name})
				inputType = *newType
			} else {
				inputType.Name.Name = e.InputTypeOverride
			}
		}
		if len(e.ResultTypeOverride) > 0 {
			if name, pkg := e.Result(); len(pkg) > 0 {
				newType := c.Universe.Type(types.Name{Package: pkg, Name: name})
				resultType = *newType
			} else {
				resultType.Name.Name = e.ResultTypeOverride
			}
		}
		m["inputType"] = &inputType
		m["resultType"] = &resultType
		m["subresourcePath"] = e.SubResourcePath
		if e.HasVerb("apply") {
			m["inputApplyConfig"] = types.Ref(path.Join(g.applyConfigurationPackage, inputGVString), inputType.Name.Name+"ApplyConfiguration")
		}

		if e.HasVerb("list") {

			sw.Do(adjustTemplate(e.VerbName, e.VerbType, methodNamespacedClusterListTemplate), m)
		}

	}

	return sw.Error()
}

// adjustTemplate adjust the origin verb template using the expansion name.
// TODO: Make the verbs in templates parametrized so the strings.Replace() is
// not needed.
func adjustTemplate(name, verbType, template string) string {
	return strings.ReplaceAll(template, " "+titler.String(verbType), " "+name)
}

// template for the struct that implements the type's interface
var structClusterGeneric = `
// $.type|privatePlural$ClusterClient implements $.type|private$Interface
type $.type|privatePlural$ClusterClient struct {
	*kcptesting.Fake
}
`

var resource = `
var $.type|allLowercasePlural$Resource = $.SchemeGroupVersion|raw$.WithResource("$.type|resource$")
`

var kind = `
var $.type|allLowercasePlural$Kind = $.SchemeGroupVersion|raw$.WithKind("$.type|singularKind$")
`

var methodNamespacedClusterClusterTemplate = `
// Cluster scopes the client down to a particular cluster.
func (c *$.type|privatePlural$ClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.$.kcpClusterResultType|public$Namespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &$.type|privatePlural$Namespacer{Fake: c.Fake, ClusterPath: clusterPath}
}
`

var methodClusterClusterTemplate = `
// Cluster scopes the client down to a particular cluster.
func (c *$.type|privatePlural$ClusterClient) Cluster(clusterPath logicalcluster.Path) $.upstreamClientInterface$ {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &$.type|privatePlural$Client{Fake: c.Fake, ClusterPath: clusterPath}
}
`

var methodNamespacedClusterListTemplate = `
// List takes label and field selectors, and returns the list of $.type|publicPlural$ that match those selectors.
func (c *$.type|privatePlural$ClusterClient) List(ctx context.Context, opts $.ListOptions|raw$) (result *$.type|raw$List, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction($.type|allLowercasePlural$Resource, $.type|allLowercasePlural$Kind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &$.type|raw$List{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &$.type|raw$List{ListMeta: obj.(*$.type|raw$List).ListMeta}
	for _, item := range obj.(*$.type|raw$List).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
`

var methodClusterListTemplate = `
// List takes label and field selectors, and returns the list of $.type|publicPlural$ that match those selectors.
func (c *$.type|privatePlural$ClusterClient) List(ctx context.Context, opts $.ListOptions|raw$) (result *$.type|raw$List, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction($.type|allLowercasePlural$Resource, $.type|allLowercasePlural$Kind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &$.type|raw$List{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &$.type|raw$List{ListMeta: obj.(*$.type|raw$List).ListMeta}
	for _, item := range obj.(*$.type|raw$List).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
`

var methodNamespacedClusterWatchTemplate = `
// Watch returns a watch.Interface that watches the requested $.type|private$s across all clusters.
func (c *$.type|privatePlural$ClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction($.type|allLowercasePlural$Resource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}
`

var methodClusterWatchTemplate = `
// Watch returns a watch.Interface that watches the requested $.type|private$s across all clusters.
func (c *$.type|privatePlural$ClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction($.type|allLowercasePlural$Resource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}
`

var structNamespacer = `
type $.type|privatePlural$Namespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}
`
var structClient = `
type $.type|privatePlural$Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}
`

var methodNamespacerNamespaceTemplate = `
func (n *$.type|privatePlural$Namespacer) Namespace(namespace string) $.upstreamClientInterface$ {
	return &$.type|privatePlural$Client{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}
`

var structNamespacedClient = `
type $.type|privatePlural$Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}
`

var methodNamespacedClientCreateTemplate = `
func (c *$.type|privatePlural$Client) Create(ctx context.Context, $.type|private$ *$.type|raw$, opts metav1.CreateOptions) (*$.type|raw$, error) {
	$if .namespaced$obj, err := c.Fake.Invokes(kcptesting.NewCreateAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, $.type|private$), &$.type|raw${})
	$else$obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction($.type|allLowercasePlural$Resource, c.ClusterPath, $.type|private$), &$.type|raw${})$end$
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

var methodNamespacedClientUpdateTemplate = `
func (c *$.type|privatePlural$Client) Update(ctx context.Context, $.type|private$ *$.type|raw$, opts metav1.UpdateOptions) (*$.type|raw$, error) {
	$if .namespaced$obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, $.type|private$), &$.type|raw${})
	$else$obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction($.type|allLowercasePlural$Resource, c.ClusterPath, $.type|private$), &$.type|raw${})
	$end$
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

var methodNamespacedClientUpdateStatusTemplate = `
func (c *$.type|privatePlural$Client) UpdateStatus(ctx context.Context, $.type|private$ *$.type|raw$, opts metav1.UpdateOptions) (*$.type|raw$, error) {
	$if .namespaced$obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, "status", c.Namespace, $.type|private$), &$.type|raw${})
	$else$obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, "status", $.type|private$), &$.type|raw${})
	$end$
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

var methodNamespacedClientDeleteTemplate = `
func (c *$.type|privatePlural$Client) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	$if .namespaced$_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, name, opts), &$.type|raw${})
	$else$_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions($.type|allLowercasePlural$Resource, c.ClusterPath, name, opts), &$.type|raw${})
	$end$
	return err
}
`

var methodNamespacedClientDeleteCollectionTemplate = `
func (c *$.type|privatePlural$Client) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	$if .namespaced$action := kcptesting.NewDeleteCollectionAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, listOpts)
	$else$action := kcptesting.NewRootDeleteCollectionAction($.type|allLowercasePlural$Resource, c.ClusterPath, listOpts)
	$end$

	_, err := c.Fake.Invokes(action, &$.type|raw$List{})
	return err
}
`

var methodNamespacedClientGetTemplate = `
func (c *$.type|privatePlural$Client) Get(ctx context.Context, name string, options metav1.GetOptions) (*$.type|raw$, error) {
	$if .namespaced$obj, err := c.Fake.Invokes(kcptesting.NewGetAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, name), &$.type|raw${})
	$else$obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction($.type|allLowercasePlural$Resource, c.ClusterPath, name), &$.type|raw${})
	$end$
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

var methodNamespacedClientListTemplate = `
func (c *$.type|privatePlural$Client) List(ctx context.Context, opts metav1.ListOptions) (*$.type|raw$List, error) {
	$if .namespaced$obj, err := c.Fake.Invokes(kcptesting.NewListAction($.type|allLowercasePlural$Resource, $.type|allLowercasePlural$Kind, c.ClusterPath, c.Namespace, opts), &$.type|raw$List{})
	$else$obj, err := c.Fake.Invokes(kcptesting.NewRootListAction($.type|allLowercasePlural$Resource, $.type|allLowercasePlural$Kind, c.ClusterPath, opts), &$.type|raw$List{})
	$end$
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &$.type|raw$List{ListMeta: obj.(*$.type|raw$List).ListMeta}
	for _, item := range obj.(*$.type|raw$List).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
`

var methodNamespacedClientWatchTemplate = `
func (c *$.type|privatePlural$Client) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	$if .namespaced$return c.Fake.InvokesWatch(kcptesting.NewWatchAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, opts))
	$else$return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction($.type|allLowercasePlural$Resource, c.ClusterPath, opts))
	$end$
}
`

var methodNamespacedClientPatchTemplate = `
func (c *$.type|privatePlural$Client) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*$.type|raw$, error) {
	$if .namespaced$obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &$.type|raw${})
	$else$obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, name, pt, data, subresources...), &$.type|raw${})
	$end$
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

//TODO: Apply config needs to be core based.

var methodNamespacedClientApplyTemplate = `
func (c *$.type|privatePlural$Client) Apply(ctx context.Context, applyConfiguration *$.inputApplyConfig|raw$, opts metav1.ApplyOptions) (*$.resultType|raw$, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := $.jsonMarshal|raw$(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	$if .namespaced$
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &$.resultType|raw${})
	$else$
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, *name, types.ApplyPatchType, data), &$.resultType|raw${})
	$end$
	if obj == nil {
		return nil, err
	}
	return obj.(*$.resultType|raw$), err
}
`

var methodNamespacedClientApplyStatusTemplate = `
func (c *$.type|privatePlural$Client) ApplyStatus(ctx context.Context, applyConfiguration *$.inputApplyConfig|raw$, opts metav1.ApplyOptions) (*$.resultType|raw$, error) {
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
	$if .namespaced$
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &$.resultType|raw${})
	$else$
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &$.resultType|raw${})
	$end$
	return obj.(*$.resultType|raw$), err
}
`
var methodNamespacedClientGetScaleTemplate = `
func (c *$.type|privatePlural$Client) GetScale(ctx context.Context, replicationControllerName string, options metav1.GetOptions) (*autoscalingv1.Scale, error) {
	$if .namespaced$obj, err := c.Fake.Invokes(kcptesting.NewGetSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, "scale", c.Namespace, replicationControllerName), &autoscalingv1.Scale{})
	$else$obj, err := c.Fake.Invokes(kcptesting.NewRootGetSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, "scale", replicationControllerName), &autoscalingv1.Scale{})
	$end$
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}
`

var methodNamespacedClientUpdateScaleTemplate = `
func (c *$.type|privatePlural$Client) UpdateScale(ctx context.Context, replicationControllerName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error) {
	$if .namespaced$obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, "scale", c.Namespace, scale), &autoscalingv1.Scale{})
	$else$obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, "scale", scale), &autoscalingv1.Scale{})
	$end$
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}
`

var methodNamespacedClientApplyScaleTemplate = `
func (c *$.type|privatePlural$Client) ApplyScale(ctx context.Context, deploymentName string, applyConfiguration *applyconfigurationsautoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (*autoscalingv1.Scale, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &autoscalingv1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}
`
