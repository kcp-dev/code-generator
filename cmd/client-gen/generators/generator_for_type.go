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
	}

	if generateApply {
		// Generated apply configuration type references required for generated Apply function
		_, gvString := util.ParsePathGroupVersion(g.inputPackage)
		m["inputApplyConfig"] = types.Ref(path.Join(g.applyConfigurationPackage, gvString), t.Name.Name+"ApplyConfiguration")
	}

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
		"cluster": `Cluster(logicalcluster.Path) $if .namespaced$$.type|public$Namespacer$else$upstreamNodeMagic$end$`,
	}
	return m
}

// group client will implement this interface.
var getterComment = `
// $.type|publicPlural$Getter has a method to return a $.type|public$Interface.
// A group's client should implement this interface.`

var getterClusterComment = `
// $.type|publicPlural$ClusterGetter has a method to return a $.type|public$ClusterInterface.
// A group's client should implement this interface.`

var getterNamespaced = `
type $.type|publicPlural$Getter interface {
	$.type|publicPlural$(namespace string) $.type|public$Interface
}
`

var getterNonNamespaced = `
type $.type|publicPlural$Getter interface {
	$.type|publicPlural$() $.type|public$Interface
}
`

var getterCluster = `
type $.type|publicPlural$ClusterGetter interface {
	$.type|publicPlural$() $.type|public$ClusterInterface
}
`

// this type's interface, typed client will implement this interface.
var interfaceTemplate1 = `
// $.type|public$Interface has methods to work with $.type|public$ resources.
type $.type|public$Interface interface {`

// this type's interface, typed client will implement this interface.
var interfaceClusterTemplate1 = `
// $.type|public$ClusterInterface has methods to work with $.type|public$ resources.
type $.type|public$ClusterInterface interface {`

var interfaceTemplate4 = `
	$.type|public$Expansion
}
`

// The following string slices are similar to maps, but with combinable keys used as indices.
// Each entry defines whether it supports lists and/or apply, and if namespacedness matters,
// namespaces; each bit is then toggled:
// * noList, noApply: index 0;
// * withList, noApply: index 2;
// * noList, withApply: index 4;
// * withList, withApply: index 6.
// When namespacedness matters, the namespaced variants are the same as the above, and
// the non-namespaced variants are offset by 1.
// Go enforces index unicity in these kinds of declarations.

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

var listTemplate = `
// List takes label and field selectors, and returns the list of $.resultType|publicPlural$ that match those selectors.
func (c *$.type|privatePlural$) List(ctx context.Context, opts $.ListOptions|raw$) (*$.resultType|raw$List, error) {
    if watchListOptions, hasWatchListOptionsPrepared, watchListOptionsErr  := $.PrepareWatchListOptionsFromListOptions|raw$(opts); watchListOptionsErr  != nil {
        klog.Warningf("Failed preparing watchlist options for $.type|resource$, falling back to the standard LIST semantics, err = %v", watchListOptionsErr )
    } else if hasWatchListOptionsPrepared {
        result, err := c.watchList(ctx, watchListOptions)
        if err == nil {
            $.CheckWatchListFromCacheDataConsistencyIfRequested|raw$(ctx, "watchlist request for $.type|resource$", c.list, opts, result)
            return result, nil
        }
        klog.Warningf("The watchlist request for $.type|resource$ ended with an error, falling back to the standard LIST semantics, err = %v", err)
    }
    result, err := c.list(ctx, opts)
    if err == nil {
        $.CheckListFromCacheDataConsistencyIfRequested|raw$(ctx, "list request for $.type|resource$", c.list, opts, result)
    }
    return result, err
}
`

var privateListTemplate = `
// list takes label and field selectors, and returns the list of $.resultType|publicPlural$ that match those selectors.
func (c *$.type|privatePlural$) list(ctx context.Context, opts $.ListOptions|raw$) (result *$.resultType|raw$List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil{
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &$.resultType|raw$List{}
	err = c.GetClient().Get().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		VersionedParams(&opts, $.schemeParameterCodec|raw$).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}
`

var listSubresourceTemplate = `
// List takes $.type|raw$ name, label and field selectors, and returns the list of $.resultType|publicPlural$ that match those selectors.
func (c *$.type|privatePlural$) List(ctx context.Context, $.type|private$Name string, opts $.ListOptions|raw$) (result *$.resultType|raw$List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil{
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &$.resultType|raw$List{}
	err = c.GetClient().Get().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name($.type|private$Name).
		SubResource("$.subresourcePath$").
		VersionedParams(&opts, $.schemeParameterCodec|raw$).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}
`

var getTemplate = `
// Get takes name of the $.type|private$, and returns the corresponding $.resultType|private$ object, and an error if there is any.
func (c *$.type|privatePlural$) Get(ctx context.Context, name string, options $.GetOptions|raw$) (result *$.resultType|raw$, err error) {
	result = &$.resultType|raw${}
	err = c.GetClient().Get().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name(name).
		VersionedParams(&options, $.schemeParameterCodec|raw$).
		Do(ctx).
		Into(result)
	return
}
`

var getSubresourceTemplate = `
// Get takes name of the $.type|private$, and returns the corresponding $.resultType|raw$ object, and an error if there is any.
func (c *$.type|privatePlural$) Get(ctx context.Context, $.type|private$Name string, options $.GetOptions|raw$) (result *$.resultType|raw$, err error) {
	result = &$.resultType|raw${}
	err = c.GetClient().Get().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name($.type|private$Name).
		SubResource("$.subresourcePath$").
		VersionedParams(&options, $.schemeParameterCodec|raw$).
		Do(ctx).
		Into(result)
	return
}
`

var deleteTemplate = `
// Delete takes name of the $.type|private$ and deletes it. Returns an error if one occurs.
func (c *$.type|privatePlural$) Delete(ctx context.Context, name string, opts $.DeleteOptions|raw$) error {
	return c.GetClient().Delete().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}
`

var createSubresourceTemplate = `
// Create takes the representation of a $.inputType|private$ and creates it.  Returns the server's representation of the $.resultType|private$, and an error, if there is any.
func (c *$.type|privatePlural$) Create(ctx context.Context, $.type|private$Name string, $.inputType|private$ *$.inputType|raw$, opts $.CreateOptions|raw$) (result *$.resultType|raw$, err error) {
	result = &$.resultType|raw${}
	err = c.GetClient().Post().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name($.type|private$Name).
		SubResource("$.subresourcePath$").
		VersionedParams(&opts, $.schemeParameterCodec|raw$).
		Body($.inputType|private$).
		Do(ctx).
		Into(result)
	return
}
`

var createTemplate = `
// Create takes the representation of a $.inputType|private$ and creates it.  Returns the server's representation of the $.resultType|private$, and an error, if there is any.
func (c *$.type|privatePlural$) Create(ctx context.Context, $.inputType|private$ *$.inputType|raw$, opts $.CreateOptions|raw$) (result *$.resultType|raw$, err error) {
	result = &$.resultType|raw${}
	err = c.GetClient().Post().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		VersionedParams(&opts, $.schemeParameterCodec|raw$).
		Body($.inputType|private$).
		Do(ctx).
		Into(result)
	return
}
`

var updateSubresourceTemplate = `
// Update takes the top resource name and the representation of a $.inputType|private$ and updates it. Returns the server's representation of the $.resultType|private$, and an error, if there is any.
func (c *$.type|privatePlural$) Update(ctx context.Context, $.type|private$Name string, $.inputType|private$ *$.inputType|raw$, opts $.UpdateOptions|raw$) (result *$.resultType|raw$, err error) {
	result = &$.resultType|raw${}
	err = c.GetClient().Put().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name($.type|private$Name).
		SubResource("$.subresourcePath$").
		VersionedParams(&opts, $.schemeParameterCodec|raw$).
		Body($.inputType|private$).
		Do(ctx).
		Into(result)
	return
}
`

var updateTemplate = `
// Update takes the representation of a $.inputType|private$ and updates it. Returns the server's representation of the $.resultType|private$, and an error, if there is any.
func (c *$.type|privatePlural$) Update(ctx context.Context, $.inputType|private$ *$.inputType|raw$, opts $.UpdateOptions|raw$) (result *$.resultType|raw$, err error) {
	result = &$.resultType|raw${}
	err = c.GetClient().Put().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name($.inputType|private$.Name).
		VersionedParams(&opts, $.schemeParameterCodec|raw$).
		Body($.inputType|private$).
		Do(ctx).
		Into(result)
	return
}
`

var watchTemplate = `
// Watch returns a $.watchInterface|raw$ that watches the requested $.type|privatePlural$.
func (c *$.type|privatePlural$) Watch(ctx context.Context, opts $.ListOptions|raw$) ($.watchInterface|raw$, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil{
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.GetClient().Get().
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		VersionedParams(&opts, $.schemeParameterCodec|raw$).
		Timeout(timeout).
		Watch(ctx)
}
`

var watchListTemplate = `
// watchList establishes a watch stream with the server and returns the list of $.resultType|publicPlural$
func (c *$.type|privatePlural$) watchList(ctx context.Context, opts $.ListOptions|raw$) (result *$.resultType|raw$List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil{
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
    result = &$.resultType|raw$List{}
	err = c.client.Get().
		$if .namespaced$Namespace(c.ns).$end$
		Resource("$.type|resource$").
		VersionedParams(&opts, $.schemeParameterCodec|raw$).
		Timeout(timeout).
		WatchList(ctx).
        Into(result)
    return
}
`

var patchTemplate = `
// Patch applies the patch and returns the patched $.resultType|private$.
func (c *$.type|privatePlural$) Patch(ctx context.Context, name string, pt $.PatchType|raw$, data []byte, opts $.PatchOptions|raw$, subresources ...string) (result *$.resultType|raw$, err error) {
	result = &$.resultType|raw${}
	err = c.GetClient().Patch(pt).
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, $.schemeParameterCodec|raw$).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
`

var applyTemplate = `
// Apply takes the given apply declarative configuration, applies it and returns the applied $.resultType|private$.
func (c *$.type|privatePlural$) Apply(ctx context.Context, $.inputType|private$ *$.inputApplyConfig|raw$, opts $.ApplyOptions|raw$) (result *$.resultType|raw$, err error) {
	if $.inputType|private$ == nil {
		return nil, fmt.Errorf("$.inputType|private$ provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := $.jsonMarshal|raw$($.inputType|private$)
	if err != nil {
		return nil, err
	}
    name := $.inputType|private$.Name
	if name == nil {
		return nil, fmt.Errorf("$.inputType|private$.Name must be provided to Apply")
	}
	result = &$.resultType|raw${}
	err = c.GetClient().Patch($.ApplyPatchType|raw$).
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name(*name).
		VersionedParams(&patchOpts, $.schemeParameterCodec|raw$).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
`

var applySubresourceTemplate = `
// Apply takes top resource name and the apply declarative configuration for $.subresourcePath$,
// applies it and returns the applied $.resultType|private$, and an error, if there is any.
func (c *$.type|privatePlural$) Apply(ctx context.Context, $.type|private$Name string, $.inputType|private$ *$.inputApplyConfig|raw$, opts $.ApplyOptions|raw$) (result *$.resultType|raw$, err error) {
	if $.inputType|private$ == nil {
		return nil, fmt.Errorf("$.inputType|private$ provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := $.jsonMarshal|raw$($.inputType|private$)
	if err != nil {
		return nil, err
	}

	result = &$.resultType|raw${}
	err = c.GetClient().Patch($.ApplyPatchType|raw$).
		$if .namespaced$Namespace(c.GetNamespace()).$end$
		Resource("$.type|resource$").
		Name($.type|private$Name).
		SubResource("$.subresourcePath$").
		VersionedParams(&patchOpts, $.schemeParameterCodec|raw$).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
`
