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

package generators

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

// genClientForType produces a file for each top-level type.
type genClientForType struct {
	generator.GoGenerator
	outputPackage             string // must be a Go import-path
	inputPackage              string
	clientsetPackage          string // must be a Go import-path
	applyConfigurationPackage string // must be a Go import-path
	group                     string
	version                   string
	groupGoName               string
	typeToMatch               *types.Type
	imports                   namer.ImportTracker

	singleClusterTypedClientsPackagePath string
}

var _ generator.Generator = &genClientForType{}

var titler = cases.Title(language.Und)

// Filter ignores all but one type because we're making a single file per type.
func (g *genClientForType) Filter(c *generator.Context, t *types.Type) bool {
	return t == g.typeToMatch
}

func (g *genClientForType) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *genClientForType) Imports(c *generator.Context) (imports []string) {
	imports = g.imports.ImportLines()
	imports = append(imports, "kcpclient \"github.com/kcp-dev/apimachinery/v2/pkg/client\"")

	if len(g.singleClusterTypedClientsPackagePath) > 0 {
		imports = append(imports, "upstream"+strings.ToLower(g.groupGoName+g.version+"client \""+g.singleClusterTypedClientsPackagePath+"/"+g.groupGoName+"/"+g.version+"\""))
	}
	imports = append(imports, "github.com/kcp-dev/logicalcluster/v3")
	imports = append(imports, "metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"")
	return
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
	return hasStatus && !util.MustParseClientGenTags(append(t.SecondClosestCommentLines, t.CommentLines...)).NoStatus
}

// GenerateType makes the body of a file implementing the individual typed client for type t.
func (g *genClientForType) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	generateApply := len(g.applyConfigurationPackage) > 0
	defaultVerbTemplates := buildDefaultVerbTemplates(generateApply)
	subresourceDefaultVerbTemplates := buildSubresourceDefaultVerbTemplates(generateApply)
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	pkg := path.Base(t.Name.Package)
	tags, err := util.ParseClientGenTags(append(t.SecondClosestCommentLines, t.CommentLines...))
	if err != nil {
		return err
	}
	type extendedInterfaceMethod struct {
		template string
		args     map[string]interface{}
	}
	_, typeGVString := util.ParsePathGroupVersion(g.inputPackage)
	extendedMethods := []extendedInterfaceMethod{}
	for _, e := range tags.Extensions {
		if e.HasVerb("apply") && !generateApply {
			continue
		}
		inputType := *t
		resultType := *t
		inputGVString := typeGVString
		// TODO: Extract this to some helper method as this code is copied into
		// 2 other places.
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
		var updatedVerbtemplate string
		if _, exists := subresourceDefaultVerbTemplates[e.VerbType]; e.IsSubresource() && exists {
			updatedVerbtemplate = e.VerbName + "(" + strings.TrimPrefix(subresourceDefaultVerbTemplates[e.VerbType], titler.String(e.VerbType)+"(")
		} else {
			updatedVerbtemplate = e.VerbName + "(" + strings.TrimPrefix(defaultVerbTemplates[e.VerbType], titler.String(e.VerbType)+"(")
		}
		extendedMethod := extendedInterfaceMethod{
			template: updatedVerbtemplate,
			args: map[string]interface{}{
				"type":          t,
				"inputType":     &inputType,
				"resultType":    &resultType,
				"CreateOptions": c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "CreateOptions"}),
				"GetOptions":    c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "GetOptions"}),
				"ListOptions":   c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ListOptions"}),
				"UpdateOptions": c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "UpdateOptions"}),
				"ApplyOptions":  c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ApplyOptions"}),
				"PatchType":     c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/types", Name: "PatchType"}),
				"jsonMarshal":   c.Universe.Type(types.Name{Package: "encoding/json", Name: "Marshal"}),
			},
		}
		if e.HasVerb("apply") {
			extendedMethod.args["inputApplyConfig"] = types.Ref(path.Join(g.applyConfigurationPackage, inputGVString), inputType.Name.Name+"ApplyConfiguration")
		}
		extendedMethods = append(extendedMethods, extendedMethod)
	}
	m := map[string]interface{}{
		"type":                             t,
		"inputType":                        t,
		"resultType":                       t,
		"package":                          pkg,
		"Package":                          namer.IC(pkg),
		"namespaced":                       !tags.NonNamespaced,
		"Group":                            namer.IC(g.group),
		"subresource":                      false,
		"subresourcePath":                  "",
		"GroupGoName":                      g.groupGoName,
		"Version":                          namer.IC(g.version),
		"CreateOptions":                    c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "CreateOptions"}),
		"DeleteOptions":                    c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "DeleteOptions"}),
		"GetOptions":                       c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "GetOptions"}),
		"ListOptions":                      c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ListOptions"}),
		"PatchOptions":                     c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "PatchOptions"}),
		"ApplyOptions":                     c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ApplyOptions"}),
		"UpdateOptions":                    c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "UpdateOptions"}),
		"PatchType":                        c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/types", Name: "PatchType"}),
		"ApplyPatchType":                   c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/types", Name: "ApplyPatchType"}),
		"watchInterface":                   c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/watch", Name: "Interface"}),
		"RESTClientInterface":              c.Universe.Type(types.Name{Package: "k8s.io/client-go/rest", Name: "Interface"}),
		"schemeParameterCodec":             c.Universe.Variable(types.Name{Package: path.Join(g.clientsetPackage, "scheme"), Name: "ParameterCodec"}),
		"jsonMarshal":                      c.Universe.Type(types.Name{Package: "encoding/json", Name: "Marshal"}),
		"resourceVersionMatchNotOlderThan": c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ResourceVersionMatchNotOlderThan"}),
		"CheckListFromCacheDataConsistencyIfRequested":      c.Universe.Function(types.Name{Package: "k8s.io/client-go/util/consistencydetector", Name: "CheckListFromCacheDataConsistencyIfRequested"}),
		"CheckWatchListFromCacheDataConsistencyIfRequested": c.Universe.Function(types.Name{Package: "k8s.io/client-go/util/consistencydetector", Name: "CheckWatchListFromCacheDataConsistencyIfRequested"}),
		"PrepareWatchListOptionsFromListOptions":            c.Universe.Function(types.Name{Package: "k8s.io/client-go/util/watchlist", Name: "PrepareWatchListOptionsFromListOptions"}),
		"Client":                                            c.Universe.Type(types.Name{Package: "k8s.io/client-go/gentype", Name: "Client"}),
		"ClientWithList":                                    c.Universe.Type(types.Name{Package: "k8s.io/client-go/gentype", Name: "ClientWithList"}),
		"ClientWithApply":                                   c.Universe.Type(types.Name{Package: "k8s.io/client-go/gentype", Name: "ClientWithApply"}),
		"ClientWithListAndApply":                            c.Universe.Type(types.Name{Package: "k8s.io/client-go/gentype", Name: "ClientWithListAndApply"}),
		"kcpClientCacheType":                                g.groupGoName + namer.IC(g.version) + "Client",
		"kcpExternalClientCacheType":                        "upstream" + strings.ToLower(g.groupGoName+g.version+"client.") + g.groupGoName + namer.IC(g.version) + "Client",
		"upstreamClientInterface":                           "upstream" + strings.ToLower(g.groupGoName+g.version+"client.") + t.Name.Name + "Interface",
	}

	if generateApply {
		// Generated apply configuration type references required for generated Apply function
		_, gvString := util.ParsePathGroupVersion(g.inputPackage)
		m["inputApplyConfig"] = types.Ref(path.Join(g.applyConfigurationPackage, gvString), t.Name.Name+"ApplyConfiguration")
	}

	namespaced := !tags.NonNamespaced

	sw.Do(getterClusterComment, m)

	sw.Do(getterCluster, m)

	//sw.Do(interfaceTemplate1, m)
	sw.Do(interfaceClusterTemplate1, m)
	if !tags.NoVerbs {
		if !genStatus(t) {
			tags.SkipVerbs = append(tags.SkipVerbs, "updateStatus")
			tags.SkipVerbs = append(tags.SkipVerbs, "applyStatus")
		}
		interfaceSuffix := ""
		if len(extendedMethods) > 0 {
			interfaceSuffix = "\n"
		}
		sw.Do("\n"+generateInterface(defaultVerbTemplates, tags)+interfaceSuffix, m)
	}

	sw.Do(interfaceTemplate4, m)
	if g.singleClusterTypedClientsPackagePath != "" {
		sw.Do(clusteredExternalInterface, m)
	} else {
		sw.Do(clusteredInterface, m)
	}

	if namespaced {
		sw.Do(clusteredMethodNamespacedCluster, m)
		sw.Do(clusteredMethodNamespacedList, m)
		sw.Do(clusteredMethodNamespacedWatch, m)

		if g.singleClusterTypedClientsPackagePath != "" {
			sw.Do(namespacerExternalInterface, m)
			sw.Do(namespacerExternalStruct, m)
			sw.Do(namespaceExternalMethod, m)
		} else {
			sw.Do(namespacerInterface, m)
			sw.Do(namespacerStruct, m)
			sw.Do(namespaceMethod, m)
		}
	} else {
		sw.Do(clusteredMethodCluster, m)
		sw.Do(clusteredMethodList, m)
		sw.Do(clusteredMethodWatch, m)
	}

	return sw.Error()
}

func generateInterface(defaultVerbTemplates map[string]string, tags util.Tags) string {
	// need an ordered list here to guarantee order of generated methods.
	out := []string{}
	for _, m := range util.SupportedVerbs {
		if tags.HasVerb(m) && len(defaultVerbTemplates[m]) > 0 {
			out = append(out, defaultVerbTemplates[m])
		}
	}
	return strings.Join(out, "\n")
}

func buildSubresourceDefaultVerbTemplates(generateApply bool) map[string]string {
	m := map[string]string{
		"list": `List(ctx context.Context, $.type|private$Name string, opts $.ListOptions|raw$) (*$.resultType|raw$List, error)`,
		"get":  `Get(ctx context.Context, $.type|private$Name string, options $.GetOptions|raw$) (*$.resultType|raw$, error)`,
	}

	return m
}

func buildDefaultVerbTemplates(generateApply bool) map[string]string {
	m := map[string]string{
		"list":    `List(ctx context.Context, opts $.ListOptions|raw$) (*$.resultType|raw$List, error)`,
		"watch":   `Watch(ctx context.Context, opts $.ListOptions|raw$) ($.watchInterface|raw$, error)`,
		"cluster": `Cluster(logicalcluster.Path) $if .namespaced$$.type|public$Namespacer$else$$.upstreamClientInterface$$end$`,
	}
	return m
}

var getterClusterComment = `
// $.type|publicPlural$ClusterGetter has a method to return a $.type|public$ClusterInterface.
// A group's client should implement this interface.`

var getterCluster = `
type $.type|publicPlural$ClusterGetter interface {
	$.type|publicPlural$() $.type|public$ClusterInterface
}
`

// this type's interface, typed client will implement this interface.
var interfaceClusterTemplate1 = `
// $.type|public$ClusterInterface has methods to work with $.type|public$ resources.
type $.type|public$ClusterInterface interface {`

var interfaceTemplate4 = `
	$.type|public$Expansion
}
`

var clusteredInterface = `
type $.type|privatePlural$ClusterInterface struct {
	clientCache kcpclient.Cache[*$.kcpClientCacheType$]
}
`

var clusteredExternalInterface = `
type $.type|privatePlural$ClusterInterface struct {
	clientCache kcpclient.Cache[*$.kcpExternalClientCacheType$]
}
`

var clusteredMethodNamespacedCluster = `
// Cluster scopes the client down to a particular cluster.
func (c *$.type|privatePlural$ClusterInterface) Cluster(clusterPath logicalcluster.Path) $.type|public$Namespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &$.type|privatePlural$Namespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}
`

var clusteredMethodCluster = `
// Cluster scopes the client down to a particular cluster.
func (c *$.type|privatePlural$ClusterInterface) Cluster(clusterPath logicalcluster.Path) $.upstreamClientInterface$ {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).$.type|publicPlural$()
}
`

var clusteredMethodList = `
// List returns the entire collection of all $.type|publicPlural$ that are available in all clusters.
func (c *$.type|privatePlural$ClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*$.resultType|raw$List, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).$.type|publicPlural$($if .namespaced$metav1.NamespaceAll$end$).List(ctx, opts)
}
`

var clusteredMethodNamespacedList = `
// List returns the entire collection of all $.type|publicPlural$ that are available in all clusters.
func (c *$.type|privatePlural$ClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*$.resultType|raw$List, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).$.type|publicPlural$(metav1.NamespaceAll).List(ctx, opts)
}
`

var clusteredMethodWatch = `
// Watch begins to watch all $.type|publicPlural$ across all clusters.
func (c *$.type|privatePlural$ClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).$.type|publicPlural$().Watch(ctx, opts)
}
`

var clusteredMethodNamespacedWatch = `
// Watch begins to watch all $.type|publicPlural$ across all clusters.
func (c *$.type|privatePlural$ClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).$.type|publicPlural$(metav1.NamespaceAll).Watch(ctx, opts)
}
`

var namespacerExternalInterface = `
// $.type|public$Namespacer can scope to objects within a namespace, returning a $.type|public$Interface.
type $.type|public$Namespacer interface {
	Namespace(name string) $.upstreamClientInterface$
}
`

var namespacerInterface = `
// $.type|public$Namespacer can scope to objects within a namespace, returning a $.type|public$Interface.
type $.type|public$Namespacer interface {
	Namespace(name string) $.type|public$Interface
}
`

var namespacerStruct = `
type $.type|privatePlural$Namespacer struct {
	clientCache kcpclient.Cache[*$.kcpClientCacheType$]
	clusterPath logicalcluster.Path
}
`

var namespacerExternalStruct = `
type $.type|privatePlural$Namespacer struct {
	clientCache kcpclient.Cache[*$.kcpExternalClientCacheType$]
	clusterPath logicalcluster.Path
}
`

var namespaceExternalMethod = `
func (n *$.type|privatePlural$Namespacer) Namespace(namespace string) $.upstreamClientInterface$ {
	return n.clientCache.ClusterOrDie(n.clusterPath).$.type|publicPlural$(namespace)
}
`

var namespaceMethod = `
func (n *$.type|privatePlural$Namespacer) Namespace(namespace string) $.type|public$Interface {
	return n.clientCache.ClusterOrDie(n.clusterPath).$.type|publicPlural$(namespace)
}
`
