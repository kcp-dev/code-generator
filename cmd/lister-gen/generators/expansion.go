/*
Copyright 2016 The Kubernetes Authors.

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
	"strings"

	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
	"k8s.io/klog/v2"

	"k8s.io/code-generator/cmd/client-gen/generators/util"
	"k8s.io/code-generator/pkg/static"
)

// expansionGenerator produces a file for a expansion interfaces.
type expansionGenerator struct {
	generator.GoGenerator
	outputPath     string
	typeToGenerate *types.Type
	// TODO: Upstream this.
	imports       namer.ImportTracker
	outputPackage string
	version       string
	group         string
	// SingleClusterListersPackagePath is the package path for the single cluster listers when using upstream listers.
	singleClusterListersPackagePath string
	// staticExpansionsListers is a map of type names to the lister interface they expand.
	staticExpansionsListers []string
}

func (g *expansionGenerator) Imports(c *generator.Context) (imports []string) {
	imports = append(imports, g.imports.ImportLines()...)
	// KCP specific
	imports = append(imports, "github.com/kcp-dev/logicalcluster/v3")
	imports = append(imports, "kcpcache \"github.com/kcp-dev/apimachinery/v2/pkg/cache\"")
	imports = append(imports, "k8s.io/api/core/v1")

	if g.singleClusterListersPackagePath != "" {
		// Sorry :(
		imp := strings.ToLower(g.group+g.version+"listers \"") + g.singleClusterListersPackagePath + "/" + strings.ToLower(g.group) + "/" + strings.ToLower(g.version) + "\""
		imports = append(imports, imp)
	}
	return
}

func (g *expansionGenerator) Filter(c *generator.Context, t *types.Type) bool {
	return t == g.typeToGenerate
}

func (g *expansionGenerator) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *expansionGenerator) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	klog.V(5).Infof("Generating expansion for %s", t)
	m := map[string]interface{}{
		"type":           t,
		"externalLister": strings.ToLower(g.group+g.version+"listers.") + t.Name.Name + "Lister",
	}

	sw := generator.NewSnippetWriter(w, c, "$", "$")
	tags := util.MustParseClientGenTags(append(t.SecondClosestCommentLines, t.CommentLines...))

	// Overrides is all or nothing when it comes to cluster or non-cluster listers.
	override, exists := g.overrideExists(t.String())
	if exists {
		klog.Infof("Overriding expansion for %s", t)
		sw.Do(override, m)
	} else {
		// default expansion behavior
		sw.Do(expansionInterfaceTemplate, m)
		if g.singleClusterListersPackagePath == "" {
			sw.Do(expansionClusterInterfaceTemplate, m)
		} else {
			sw.Do(expansionClusterExternalInterfaceTemplate, m)
		}

		if !tags.NonNamespaced {
			sw.Do(namespacedExpansionInterfaceTemplate, m)
		}
	}

	return sw.Error()
}

func (g *expansionGenerator) overrideExists(typeName string) (string, bool) {
	for _, listers := range g.staticExpansionsListers {
		source, target := strings.Split(listers, ":")[0], strings.Split(listers, ":")[1]
		if source == typeName {
			klog.Infof("Found expansion override for %s", typeName)
			override := static.GetListersExpansions(target)
			return override, true
		}
	}
	return "", false
}

var expansionInterfaceTemplate = `
// $.type|public$ListerExpansion allows custom methods to be added to
// $.type|public$Lister.
type $.type|public$ListerExpansion interface {}
`

var namespacedExpansionInterfaceTemplate = `
// $.type|public$NamespaceListerExpansion allows custom methods to be added to
// $.type|public$NamespaceLister.
type $.type|public$NamespaceListerExpansion interface {}
`

var expansionClusterInterfaceTemplate = `
// $.type|public$ClusterListerExpansion allows custom methods to be added to
// $.type|public$Lister.
type $.type|public$ClusterListerExpansion interface {
	// Cluster returns a lister that can list and get $.type|public$ in one workspace.
	Cluster(clusterName logicalcluster.Name) $.type|public$Lister
}
`

var expansionClusterExternalInterfaceTemplate = `
// $.type|public$ClusterListerExpansion allows custom methods to be added to
// $.type|public$Lister.
type $.type|public$ClusterListerExpansion interface {
	// Cluster returns a lister that can list and get $.type|public$ in one workspace.
	Cluster(clusterName logicalcluster.Name) $.externalLister$
}
`
