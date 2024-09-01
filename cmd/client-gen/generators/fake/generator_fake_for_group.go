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
	"fmt"
	"io"
	"path"
	"strings"

	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"

	"k8s.io/code-generator/cmd/client-gen/generators/util"
)

// genFakeForGroup produces a file for a group client, e.g. ExtensionsClient for the extension group.
type genFakeForGroup struct {
	generator.GoGenerator
	outputPackage     string // must be a Go import-path
	realClientPackage string // must be a Go import-path
	group             string
	version           string
	groupGoName       string
	// types in this group
	types   []*types.Type
	imports namer.ImportTracker
	// If the genGroup has been called. This generator should only execute once.
	called                               bool
	singleClusterTypedClientsPackagePath string
}

var _ generator.Generator = &genFakeForGroup{}

// We only want to call GenerateType() once per group.
func (g *genFakeForGroup) Filter(c *generator.Context, t *types.Type) bool {
	if !g.called {
		g.called = true
		return true
	}
	return false
}

func (g *genFakeForGroup) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *genFakeForGroup) Imports(c *generator.Context) (imports []string) {
	imports = g.imports.ImportLines()
	if len(g.types) != 0 {
		imports = append(imports, fmt.Sprintf("%s \"%s\"", strings.ToLower(path.Base(g.realClientPackage)), g.realClientPackage))
	}
	imports = append(imports, "upstream"+strings.ToLower(g.groupGoName+g.version+"client \""+g.singleClusterTypedClientsPackagePath+"/"+g.groupGoName+"/"+g.version+"\""))
	imports = append(imports, "kcptesting \"github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing\"")
	imports = append(imports, "metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"")
	return imports
}

func (g *genFakeForGroup) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")

	externalClient := len(g.singleClusterTypedClientsPackagePath) > 0

	m := map[string]interface{}{
		"type":                    t,
		"GroupGoName":             g.groupGoName,
		"Version":                 namer.IC(g.version),
		"Fake":                    c.Universe.Type(types.Name{Package: "k8s.io/client-go/testing", Name: "Fake"}),
		"RESTClientInterface":     c.Universe.Type(types.Name{Package: "k8s.io/client-go/rest", Name: "Interface"}),
		"RESTClient":              c.Universe.Type(types.Name{Package: "k8s.io/client-go/rest", Name: "RESTClient"}),
		"upstreamClientInterface": "upstream" + strings.ToLower(g.groupGoName+g.version+"client.") + g.groupGoName + strings.Title(g.version) + "Interface",
		"useUpstreamClient":       externalClient,
	}

	// Cluster clients.
	sw.Do(groupClusterClientTemplate, m)
	sw.Do(getterClusterMethod, m)
	for _, t := range g.types {
		tags, err := util.ParseClientGenTags(append(t.SecondClosestCommentLines, t.CommentLines...))
		if err != nil {
			return err
		}
		wrapper := map[string]interface{}{
			"type":                    t,
			"GroupGoName":             g.groupGoName,
			"Version":                 namer.IC(g.version),
			"realClientPackage":       strings.ToLower(path.Base(g.realClientPackage)),
			"upstreamClientInterface": "upstream" + strings.ToLower(g.groupGoName+g.version+"client.") + t.Name.Name + "Interface",
			"useUpstreamClient":       externalClient,
		}
		if tags.NonNamespaced {
			sw.Do(getterClusterImplNonNamespaced, wrapper)
			continue
		}
		sw.Do(getterClusterImplNamespaced, wrapper)
	}

	// non-cluster clients
	sw.Do(groupClientTemplate, m)

	sw.Do(getRESTClient, m)

	for _, t := range g.types {
		tags, err := util.ParseClientGenTags(append(t.SecondClosestCommentLines, t.CommentLines...))
		if err != nil {
			return err
		}
		wrapper := map[string]interface{}{
			"type":                    t,
			"GroupGoName":             g.groupGoName,
			"Version":                 namer.IC(g.version),
			"realClientPackage":       strings.ToLower(path.Base(g.realClientPackage)),
			"upstreamClientInterface": "upstream" + strings.ToLower(g.groupGoName+g.version+"client.") + t.Name.Name + "Interface",
			"useUpstreamClient":       externalClient,
		}
		if tags.NonNamespaced {
			sw.Do(getterImplNonNamespaced, wrapper)
			continue
		}
		sw.Do(getterImplNamespaced, wrapper)
	}

	return sw.Error()
}

var groupClusterClientTemplate = `
type $.GroupGoName$$.Version$ClusterClient struct {
	*kcptesting.Fake
}
`

var groupClientTemplate = `
type $.GroupGoName$$.Version$Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}
`

var getterClusterImplNamespaced = `
func (c *$.GroupGoName$$.Version$ClusterClient) $.type|publicPlural$(namespace string) $.realClientPackage$.$.type|public$ClusterInterface {
	return &$.type|privatePlural$ClusterClient{Fake: c.Fake}
}
`

var getterClusterImplNonNamespaced = `
func (c *$.GroupGoName$$.Version$ClusterClient) $.type|publicPlural$() $.realClientPackage$.$.type|public$ClusterInterface {
	return &$.type|privatePlural$ClusterClient{Fake: c.Fake}
}
`

var getterClusterMethod = `
$if .useUpstreamClient$
func (c *$.GroupGoName$$.Version$ClusterClient) Cluster(clusterPath logicalcluster.Path) $.upstreamClientInterface$ {
$else$
func (c *$.GroupGoName$$.Version$ClusterClient) Cluster(clusterPath logicalcluster.Path) $.GroupGoName$$.Version$Client {
$end$
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &$.GroupGoName$$.Version$Client{Fake: c.Fake, ClusterPath: clusterPath}
}
`

var getRESTClient = `
// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *$.GroupGoName$$.Version$Client) RESTClient() $.RESTClientInterface|raw$ {
	var ret *$.RESTClient|raw$
	return ret
}
`

var getterImplNonNamespaced = `
$if .useUpstreamClient$
func (c *$.GroupGoName$$.Version$Client) $.type|publicPlural$() $.upstreamClientInterface$ {
$else$
func (c *$.GroupGoName$$.Version$Client) $.type|publicPlural$() $.GroupGoName$$.Version$Client {
$end$
	return &$.type|privatePlural$Client{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
`

var getterImplNamespaced = `
$if .useUpstreamClient$
func (c *$.GroupGoName$$.Version$Client) $.type|publicPlural$(namespace string) $.upstreamClientInterface$ {
$else$
func (c *$.GroupGoName$$.Version$Client) $.type|publicPlural$(namespace string) $.GroupGoName$$.Version$Client {
$end$
	return &$.type|privatePlural$Client{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
`
