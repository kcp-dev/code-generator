package main

import (
	"flag"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	"github.com/kcp-dev/client-gen/listerpoc/example/informers/externalversions"
)

// List all ConfigMaps (kcp-wide)
// List all ConfigMaps in a workspace
// List all ConfigMaps in a namespace in a workspace
// Get a ConfigMap with a specific name in a namespace in a workspace
func main() {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	f := externalversions.NewSharedInformerFactory(clientset, 0)
	clusterInf := f.Core().V1().ConfigMaps()
	inf := clusterInf.Informer()
	stopChn := make(chan struct{})
	go inf.Run(stopChn)
	cache.WaitForCacheSync(stopChn, inf.HasSynced)
	lst := clusterInf.Lister()
	allCms, err := lst.List(labels.Everything())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("across all of KCP")
	printobjs(allCms)

	for _, cm := range allCms {
		cluster := cm.GetZZZ_DeprecatedClusterName()
		wsLst := lst.Cluster(logicalcluster.New(cluster))
		allWsCms, err := wsLst.List(labels.Everything())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Workspace: %s\n", cluster)
		printobjs(allWsCms)

		nsLst := wsLst.ConfigMaps(cm.GetNamespace())
		allNsCms, err := nsLst.List(nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Workspace: %s\nNamespace: %s\n", cluster, cm.GetNamespace())
		printobjs(allNsCms)

		oneCm, err := nsLst.Get(cm.GetName())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Workspace: %s\nNamespace: %s\nName: %s\n", cluster, cm.GetNamespace(), cm.GetName())
		printobjs([]*v1.ConfigMap{oneCm})
	}
}

func printobjs(cms []*v1.ConfigMap) {
	for _, cm := range cms {
		fmt.Printf("\t\t %s/%s/%s\n", cm.GetZZZ_DeprecatedClusterName(), cm.GetNamespace(), cm.GetName())
	}
	fmt.Println()
}
