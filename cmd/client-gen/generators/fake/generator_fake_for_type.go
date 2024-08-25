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

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"

	"k8s.io/code-generator/cmd/client-gen/generators/util"
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
	return imports
}

// Ideally, we'd like genStatus to return true if there is a subresource path
// registered for "status" in the API server, but we do not have that
// information, so genStatus returns true if the type has a status field.
func genStatus(t *types.Type) bool {
	// Default to true if we have a Status member
	hasStatus := false
	for _, m := range t.Members {
		if m.Name == "Status" {
			hasStatus = true
			break
		}
	}

	tags := util.MustParseClientGenTags(append(t.SecondClosestCommentLines, t.CommentLines...))
	return hasStatus && !tags.NoStatus
}

// GenerateType makes the body of a file implementing the individual typed client for type t.
func (g *genFakeForType) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	pkg := path.Base(t.Name.Package)
	tags, err := util.ParseClientGenTags(append(t.SecondClosestCommentLines, t.CommentLines...))
	if err != nil {
		return err
	}

	const pkgClientGoTesting = "k8s.io/client-go/testing"
	m := map[string]interface{}{
		"type":               t,
		"inputType":          t,
		"resultType":         t,
		"subresourcePath":    "",
		"package":            pkg,
		"Package":            namer.IC(pkg),
		"namespaced":         !tags.NonNamespaced,
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
	m["inputApplyConfig"] = types.Ref(path.Join(g.applyConfigurationPackage, gvString), t.Name.Name+"ApplyConfiguration")
	m["kcpClusterResultType"] = types.Ref(g.typedClientPackage, t.Name.Name)
	m["upstreamClientInterface"] = "upstream" + strings.ToLower(g.groupGoName+g.version+"client.") + t.Name.Name + "Interface"

	namespaced := !tags.NonNamespaced

	sw.Do(resource, m)
	sw.Do(kind, m)

	sw.Do(structClusterGeneric, m)

	sw.Do(methodClusterClusterTemplate, m)
	sw.Do(methodClusterListTemplate, m)

	sw.Do(methodClusterWatchTemplate, m)

	if namespaced {
		sw.Do(structNamespacer, m)
		sw.Do(methodNamespacerNamespaceTemplate, m)

		sw.Do(structClient, m)
		sw.Do(methodClientCreateTemplate, m)
		sw.Do(methodClientUpdateTemplate, m)
		sw.Do(methodClientUpdateStatusTemplate, m)
		sw.Do(methodClientDeleteTemplate, m)
		sw.Do(methodClientDeleteCollectionTemplate, m)
		sw.Do(methodClientGetTemplate, m)
		sw.Do(methodClientListTemplate, m)
		sw.Do(methodClientWatchTemplate, m)
		sw.Do(methodClientPatchTemplate, m)
		if generateApply {
			sw.Do(methodClientApplyTemplate, m)
		}
	}

	_, typeGVString := util.ParsePathGroupVersion(g.inputPackage)

	// generate extended client methods
	for _, e := range tags.Extensions {
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

			sw.Do(adjustTemplate(e.VerbName, e.VerbType, methodClusterListTemplate), m)
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

var methodClusterClusterTemplate = `
// Cluster scopes the client down to a particular cluster.
func (c *$.type|privatePlural$ClusterClient) Cluster(clusterPath logicalcluster.Path) *kcp.$.kcpClusterResultType|public$Namespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &$.type|privatePlural$Namespacer{Fake: c.Fake, ClusterPath: clusterPath}
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

var methodNamespacerNamespaceTemplate = `
func (n *$.type|privatePlural$Namespacer) Namespace(namespace string) $.upstreamClientInterface$ {
	return &configMapsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}
`

var structClient = `
type $.type|privatePlural$Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}
`

var methodClientCreateTemplate = `
func (c *$.type|privatePlural$Client) Create(ctx context.Context, $.type|private$ *$.type|raw$, opts metav1.CreateOptions) (*$.type|raw$, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, $.type|private$), &$.type|raw${})
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

var methodClientUpdateTemplate = `
func (c *$.type|privatePlural$Client) Update(ctx context.Context, $.type|private$ *$.type|raw$, opts metav1.CreateOptions) (*$.type|raw$, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, $.type|private$), &$.type|raw${})
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

var methodClientUpdateStatusTemplate = `
func (c *$.type|privatePlural$Client) UpdateStatus(ctx context.Context, $.type|private$ *$.type|raw$, opts metav1.CreateOptions) (*$.type|raw$, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, "status", c.Namespace, $.type|private$), &$.type|raw${})
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

var methodClientDeleteTemplate = `
func (c *$.type|privatePlural$Client) Delete(ctx context.Context, name string, opts metav1.CreateOptions)  error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, name, opts), &$.type|raw${})
	return err
}
`

var methodClientDeleteCollectionTemplate = `
func (c *$.type|privatePlural$Client) DeleteCollection(ctx context.Context,  opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &$.type|raw$List{})
	return err
}
`

var methodClientGetTemplate = `
func (c *$.type|privatePlural$Client) Get(ctx context.Context, name string, options metav1.GetOptions) (*$.type|raw$, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, name), &$.type|raw${})
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

var methodClientListTemplate = `
// List takes label and field selectors, and returns the list of $.type|raw$ that match those selectors.
func (c *$.type|privatePlural$Client) List(ctx context.Context, opts metav1.ListOptions) (*$.type|raw$List, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction($.type|allLowercasePlural$Resource, $.type|allLowercasePlural$Kind, c.ClusterPath, c.Namespace, opts), &$.type|raw$List{})
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
var methodClientWatchTemplate = `
func (c *$.type|privatePlural$Client) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, opts))
}
`
var methodClientPatchTemplate = `
func (c *$.type|privatePlural$Client) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*$.type|raw$, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &$.type|raw${})
	if obj == nil {
		return nil, err
	}
	return obj.(*$.type|raw$), err
}
`

//TODO: Apply config needs to be core based.

var methodClientApplyTemplate = `
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction($.type|allLowercasePlural$Resource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &$.resultType|raw${})
	if obj == nil {
		return nil, err
	}
	return obj.(*$.resultType|raw$), err
}
`
